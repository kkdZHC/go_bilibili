package cache

import (
	"go_bilibili/utils"

	"github.com/go-redis/redis"
	"github.com/spf13/viper"
)

var RedisClient *redis.Client

func InitRedis() {
	client := redis.NewClient(&redis.Options{
		Addr:     viper.GetString("redis.address"),
		Password: viper.GetString("redis.password"),
		DB:       2,
	})
	_, err := client.Ping().Result()
	if err != nil {
		utils.Logfile("[Error]", "redis error"+err.Error())
	}
	RedisClient = client
}
