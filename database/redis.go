package database

import (
	"github.com/2ndSilencerz/redis-data-pusher/config"
	"github.com/go-redis/redis/v8"
)

// InitRedis used to get redis instance
func InitRedis() *redis.Client {
	auth := config.GetRedisAuth()
	return redis.NewClient(&redis.Options{
		Addr:     auth.Addr,
		Password: auth.Password, // no password set
		DB:       auth.DB,       // use default DB
	})
}
