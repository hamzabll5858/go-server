CGO_ENABLED=0
GOOS=linux
GOARCH=amd64
BIN="bin"
APP="service"
SRC=$(shell find . -name "*.go")
SERVICE_PREFIX=''
SERVICE_NAME='service'

ifeq (, $(shell which go))
$(warning "could not find go in $(PATH), Install golang")
endif

.PHONY: all build install_deps

default: all

all: build

build: install_deps
	go \
        build -a -installsuffix cgo -v \
        -o $(BIN)/$(APP) . \
        ;

install_deps:
	$(info ******************** downloading dependencies ********************)
	go get -v ./...
