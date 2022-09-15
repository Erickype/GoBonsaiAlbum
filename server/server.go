package main

import (
	"context"
	"fmt"
	pbUsers "github.com/Erickype/GoBonsaiAlbum/gRPC/users"
	"github.com/Erickype/GoBonsaiAlbum/redis"
	"net"
)

func main() {
	ctx := context.Background()
	redisClient := redis.NewRedisClient()

	result, err := redis.PingRedisClient(redisClient, ctx)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic("Cannot create tcp connection: " + err.Error())
	}

	serv := pbUsers.ImplementServer()

	err = serv.Serve(listen)

	if err != nil {
		panic("Cannot initialize server: " + err.Error())
	}
}
