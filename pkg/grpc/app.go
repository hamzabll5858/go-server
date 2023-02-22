package grpc

import (
	"github.com/spf13/viper"
	grpc_server "go-server/pkg/grpc/server"
	"go-server/pkg/models"

	"log"
	"net"

	"google.golang.org/grpc"
)

func InitGRPC() {
	println("gRPC User Services running on Addr: "+viper.GetString("grpc.url"))

	models.InitDB()
	listener, err := net.Listen("tcp", viper.GetString("grpc.url"))
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	grpc_server.RegisterUserServer(s, &grpc_server.Server{})
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}