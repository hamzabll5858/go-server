package redis

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

func GetRedisCache()  *redis.Client{
	cache := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%v:%v", viper.GetString("redis.addr"),viper.GetString("redis.port")),
		Password: viper.GetString("redis.password"),
		DB: 0,
	})

	return cache
}