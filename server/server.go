package main

import (
	"context"
	"fmt"
	pbService "github.com/Erickype/GoBonsaiAlbum/gRPC"
	"github.com/Erickype/GoBonsaiAlbum/redis"
	"net"
)

func main() {
	ctx := context.Background()
	redisClient := redis.NewRedisClient()

	result, err := redis.PingRedisClient(redisClient, ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)

	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic("Cannot create tcp connection: " + err.Error())
	}

	serv := pbService.ImplementServer()

	err = serv.Serve(listen)

	if err != nil {
		panic("Cannot initialize server: " + err.Error())
	}
}
