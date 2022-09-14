package users

import "context"

type server struct {
	UnimplementedUsersServer
}

func (s *server) CreateUser(ctx context.Context, req *CreateUserReq, res *CreateUserRes) {

}
