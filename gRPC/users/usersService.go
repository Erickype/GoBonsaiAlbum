package users

import (
	"context"
	"github.com/Erickype/GoBonsaiAlbum/gRPC/models"
	"github.com/Erickype/GoBonsaiAlbum/mysql/users"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/runtime/protoimpl"
)

type server struct {
	UnimplementedUsersServer
}

func (s *server) CreateUser(ctx context.Context, req *CreateUserReq) (*CreateUserRes, error) {
	user := models.User{
		UserName:     req.GetUserName(),
		UserLastname: req.GetUserLastname(),
		UserNickname: req.GetUserNickname(),
	}

	users.CreateUser(&user)

	return &CreateUserRes{
		state:         protoimpl.MessageState{},
		sizeCache:     0,
		unknownFields: nil,
		Id:            0,
	}, nil
}

func ImplementServer() *grpc.Server {

	serv := grpc.NewServer()

	//Server register
	RegisterUsersServer(serv, &server{})

	return serv
}
