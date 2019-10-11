package util

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/go-redis/redis"
)

func NewRedisClient() *redis.Client{
	redisClient := redis.NewClient(&redis.Options{
		Addr:     beego.AppConfig.String("redis::addr"),
		Password: beego.AppConfig.String("redis::password"), // no password set //foobared
		DB:       1,  // use default DB
	})
	
	pong, err := redisClient.Ping().Result()
	if err != nil {
		fmt.Printf("ping error[%s]\n", err.Error())
	}
	fmt.Printf("ping result: %s\n", pong)
	return redisClient
}