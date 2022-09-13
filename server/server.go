package main

import (
	"context"
	"fmt"
	"github.com/Erickype/GoBonsaiAlbum/mysql"
	"github.com/Erickype/GoBonsaiAlbum/redis"
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

	db, err := mysql.GetMysqlConnection()
	if err != nil {
		panic(err)
	}
	fmt.Println(db.Ping())
}
