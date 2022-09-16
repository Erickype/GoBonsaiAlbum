package main

import (
	"context"
	"fmt"
	pbUsers "github.com/Erickype/GoBonsaiAlbum/gRPC/users"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		panic("Connection error:  " + err.Error())
	}

	clientService := pbUsers.NewUsersClient(conn)

	CreateUser(clientService)
}

func CreateUser(clientService pbUsers.UsersClient) {
	res, err := clientService.CreateUser(context.Background(), &pbUsers.CreateUserReq{
		UserName:     "Erick",
		UserLastname: "Carreras",
		UserNickname: "Erickype",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}
