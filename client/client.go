package main

import (
	"context"
	"fmt"
	pbService "github.com/Erickype/GoBonsaiAlbum/gRPC"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"log"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		panic("Connection error:  " + err.Error())
	}

	clientService := pbService.NewServiceGRPCClient(conn)

	//CreateUser(clientService)
	GetUsers(clientService)
	UpdateUser(clientService)
}

func CreateUser(clientService pbService.ServiceGRPCClient) {
	res, err := clientService.CreateUser(context.Background(), &pbService.CreateUserReq{
		UserName:     "Erick",
		UserLastname: "Carreras",
		UserNickname: "Erickype",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}

func GetUsers(clientService pbService.ServiceGRPCClient) {
	stream, err := clientService.GetUsers(context.Background(), &pbService.GetUsersReq{Id: 0})
	if err != nil {
		log.Fatalf("clientService.GetUsers failed: %v", err)
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("clientService.GetUsers failed: %v", err)
		}
		user := res.User

		log.Printf("User: %v, Name: %v, Lastname:%v, Nickname:%v", user.Id,
			user.UserName, user.UserLastname, user.UserNickname)
	}
}

func UpdateUser(clientService pbService.ServiceGRPCClient) {

	user := pbService.User{
		Id:           11,
		UserName:     "Joss",
		UserLastname: "San",
		UserNickname: "Joa",
	}

	res, err := clientService.UpdateUser(context.Background(), &pbService.UpdateUserReq{
		User: &user,
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Updated: %v, Error: %v", res.Updated, res.Error)
}
