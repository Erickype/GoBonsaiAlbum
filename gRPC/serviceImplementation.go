package gRPC

import (
	"context"
	"github.com/Erickype/GoBonsaiAlbum/models"
	mysql "github.com/Erickype/GoBonsaiAlbum/mysql/users"
	"google.golang.org/genproto/googleapis/type/date"
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

func (s *server) UpdateUser(_ context.Context, req *UpdateUserReq) (*UpdateUserRes, error) {

	updateUserRes := &UpdateUserRes{
		state:         protoimpl.MessageState{},
		sizeCache:     0,
		unknownFields: nil,
		Updated:       false,
		Error:         "",
	}
	user := models.User{
		Id:           req.User.GetId(),
		UserName:     req.User.GetUserName(),
		UserLastname: req.User.GetUserLastname(),
		UserNickname: req.User.GetUserNickname(),
		CreatedAt:    date.Date{},
	}

	result, err := mysql.UpdateUser(&user)
	if err != nil {
		updateUserRes.Error = err.Error()
		return updateUserRes, err
	}

	updateUserRes.Updated = result != 0

	return updateUserRes, nil
}

func (s *server) DeleteUser(_ context.Context, req *DeleteUserReq) (*DeleteUserRes, error) {

	deleteUserRes := &DeleteUserRes{
		state:         protoimpl.MessageState{},
		sizeCache:     0,
		unknownFields: nil,
		Deleted:       false,
		Error:         "",
	}

	result, err := mysql.DeleteUser(req.Id)
	if err != nil {
		deleteUserRes.Error = err.Error()
		return deleteUserRes, err
	}

	deleteUserRes.Deleted = result != 0

	return deleteUserRes, nil
}

func (s *server) LoadSavedUsers(users []*models.User) {
	s.savedUsers = nil
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
