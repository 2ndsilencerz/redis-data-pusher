package database

import (
	"fmt"
	"github.com/2ndSilencerz/redis-data-pusher/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

// InitDB used to get db instance to use
func InitDB() *gorm.DB {
	dbHost := config.GetDatabaseConfig()
	db, err := gorm.Open(postgres.Open(dbHost), &gorm.Config{})
	if err != nil {
		msg := fmt.Sprintf("Error: %v\n", err)
		config.LogToFile(msg)
		log.Panicln(msg)
	}
	return db
}

// CloseDB need to be called each time a db instance are used/called
// so the connection will close and will not use more memory
func CloseDB(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		config.LogToFile(fmt.Sprintf("Error: %v", err))
	}
	err = sqlDB.Close()
	if err != nil {
		config.LogToFile(fmt.Sprintf("Error: %v", err))
	}
}
