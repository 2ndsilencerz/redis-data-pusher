package database

import (
	"fmt"
	"github.com/2ndSilencerz/redis-data-pusher/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// InitDB used to get db instance to use
func InitDB() *gorm.DB {
	dbHost := config.GetDatabaseConfig()
	db, err := gorm.Open(postgres.Open(dbHost), &gorm.Config{})
	if err != nil {
		config.LogToFile(fmt.Sprintf("Error: %v\n", err))
	}
	return db
}
