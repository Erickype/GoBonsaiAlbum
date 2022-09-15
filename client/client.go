package main

import (
	"context"
	"fmt"
	pbUsers "github.com/Erickype/GoBonsaiAlbum/gRPC/users"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		panic("Connection error:  " + err.Error())
	}

	clientService := pbUsers.NewUsersClient(conn)

	res, err := clientService.CreateUser(context.Background(), &pbUsers.CreateUserReq{
		UserName:     "Erick",
		UserLastname: "Carrasco",
		UserNickname: "Erickype",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}
