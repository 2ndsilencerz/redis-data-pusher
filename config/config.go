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

// GetDatabaseConfig for database connection usage in database.go
func GetDatabaseConfig() string {
	v := loadConfig()
	return v.GetString("postgres")
}

// GetCronInterval for interval setting in cron scheduler
func GetCronInterval() int {
	v := loadConfig()
	return v.GetInt("cronInterval")
}

// RedisAuth is a data for redis configuration value
type RedisAuth struct {
	Addr     string
	Password string
	DB       int
}

// GetRedisAuth used to return value for redis config
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
