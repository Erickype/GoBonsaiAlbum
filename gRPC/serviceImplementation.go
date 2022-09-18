package gRPC

import (
	"context"
	"github.com/Erickype/GoBonsaiAlbum/models"
	mysql "github.com/Erickype/GoBonsaiAlbum/mysql/users"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/runtime/protoimpl"
	"log"
)

type server struct {
	UnimplementedServiceGRPCServer
	savedUsers []*User
}

func (s *server) CreateUser(_ context.Context, req *CreateUserReq) (*CreateUserRes, error) {
	user := models.User{
		UserName:     req.GetUserName(),
		UserLastname: req.GetUserLastname(),
		UserNickname: req.GetUserNickname(),
	}

	result, err := mysql.CreateUser(&user)

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

func (s *server) GetUsers(_ *GetUsersReq, stream ServiceGRPC_GetUsersServer) error {

	users := mysql.GetUsers()

	s.LoadSavedUsers(users)

	for _, user := range s.savedUsers {
		err := stream.Send(&GetUsersRes{
			state:         protoimpl.MessageState{},
			sizeCache:     0,
			unknownFields: nil,
			User:          user,
		})
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *server) LoadSavedUsers(users []*models.User) {
	for _, user := range users {
		var savedUser = User{
			state:         protoimpl.MessageState{},
			sizeCache:     0,
			unknownFields: nil,
			Id:            user.Id,
			UserName:      user.UserName,
			UserLastname:  user.UserLastname,
			UserNickname:  user.UserNickname,
		}
		s.savedUsers = append(s.savedUsers, &savedUser)
	}
}

func ImplementServer() *grpc.Server {
	serv := grpc.NewServer()
	//Server register
	RegisterServiceGRPCServer(serv, &server{})
	return serv
}
