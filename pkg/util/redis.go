package util

import (
	"github.com/astaxie/beego"
	"github.com/go-redis/redis/v7"
	"sync"
)

var (
	once sync.Once
	redisClient *redis.Client
)

/**
 * 单例模式 实现redis连接
 */
func NewRedisClient() (*redis.Client,error) {
	once.Do(func() {
		index,_:=  beego.AppConfig.Int("redis::index")
		redisClient = redis.NewClient(&redis.Options{
			Addr:     	beego.AppConfig.String("redis::addr"),
			Password: 	beego.AppConfig.String("redis::password"),
			DB:       	index,
			PoolSize:	10,
		})
	})

	_,err := redisClient.Ping().Result()
	if err!=nil {
		return nil,err
	}else{
		return redisClient,nil
	}
}