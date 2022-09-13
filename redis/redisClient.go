package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
)

func NewRedisClient() *redis.Client {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "redis-15948.c90.us-east-1-3.ec2.cloud.redislabs.com:15948",
		Password: "mTquyoBo4faipzY5nOjI3amOyHmUDZRK", // no password
		DB:       0,                                  // default DB
	})

	return redisClient
}

func PingRedisClient(client *redis.Client, ctx context.Context) (string, error) {
	result, err := client.Ping(ctx).Result()

	if err != nil {
		return "", err
	} else {
		return result, nil
	}
}
