package grpc

import (
	context "context"
	"encoding/json"
	"go-server/pkg/logger"
	"go-server/pkg/models"
)

type Server struct {
	UnimplementedUserServer
}

func (s *Server) GetUser(ctx context.Context, in *UserRequest) (*UserReply, error) {
	db := models.DB

	var user models.User

	db.First(&user, in.GetId())

	data ,err := json.Marshal(user)
	if err != nil{
		logger.Logger.Error("Error: GRPC User failed to fetch")
	}

	var reply  UserReply
	err = json.Unmarshal(data,&reply)
	if err != nil{
		logger.Logger.Error("Error: GRPC User failed to fetch")
	}

	return &reply, nil
}