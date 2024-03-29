# ------- BASE IMAGE ------- #
FROM golang:1.20 AS base

WORKDIR /usr/src/app/

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -v -o build/go-server .

# ------- EXECUTION IMAGE ------- #
FROM alpine:3.17.2

RUN apk --no-cache add ca-certificates

WORKDIR /build/

COPY --from=base /usr/src/app/build/go-server .
COPY --from=base /usr/src/app/config/config.yaml .

# Non root user provisioning for least privilege
RUN addgroup -S  executor
RUN adduser -S executor -G executor
RUN chown -R executor:executor /build/
USER executor

EXPOSE 8080
CMD ["./go-server"]