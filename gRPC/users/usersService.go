package users

import (
	"context"
	"github.com/Erickype/GoBonsaiAlbum/models"
	"github.com/Erickype/GoBonsaiAlbum/mysql/users"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/runtime/protoimpl"
	"log"
)

type server struct {
	UnimplementedUsersServer
}

func (s *server) CreateUser(_ context.Context, req *CreateUserReq) (*CreateUserRes, error) {
	user := models.User{
		UserName:     req.GetUserName(),
		UserLastname: req.GetUserLastname(),
		UserNickname: req.GetUserNickname(),
	}

	result, err := users.CreateUser(&user)

	if err != nil {
		log.Fatalln(err)
	}

	return &CreateUserRes{
		state:         protoimpl.MessageState{},
		sizeCache:     0,
		unknownFields: nil,
		Id:            int32(result),
	}, err
}

func ImplementServer() *grpc.Server {

	serv := grpc.NewServer()

	//Server register
	RegisterUsersServer(serv, &server{})

	return serv
}
