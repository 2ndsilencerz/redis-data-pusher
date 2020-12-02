package config

import (
	"fmt"
	"github.com/spf13/viper"
)

func loadConfig() *viper.Viper {
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("json")
	v.AddConfigPath(".")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	return v
}

func GetDatabaseConfig() string {
	v := loadConfig()
	return v.GetString("postgres")
}

func GetCronInterval() int {
	v := loadConfig()
	return v.GetInt("cronInterval")
}

type RedisAuth struct {
	Addr     string
	Password string
	DB       int
}

func GetRedisAuth() RedisAuth {
	v := loadConfig()
	addr := v.GetString("redisHost")
	pass := v.GetString("redisPass")
	db := v.GetInt("redisDB")

	return RedisAuth{
		Addr:     addr,
		Password: pass,
		DB:       db,
	}
}
