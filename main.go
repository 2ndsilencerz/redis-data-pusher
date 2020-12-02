package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/2ndSilencerz/redis-data-pusher/config"
	"github.com/2ndSilencerz/redis-data-pusher/database"
	"github.com/2ndSilencerz/redis-data-pusher/model"
	"github.com/go-co-op/gocron"
	"time"
)

var ctx = context.Background()
var db = database.InitDB()
var redisInst = database.InitRedis()

func main() {
	interval := config.GetCronInterval()
	durations, err := time.ParseDuration(string(rune(interval)))
	s1 := gocron.NewScheduler(time.UTC)
	_, err = s1.Every(uint64(durations)).Second().Do(loadData)
	if err != nil {
		config.LogToFile(fmt.Sprintf("Error: %v\n", err))
	}

	fmt.Println("starting scheduler")
	s1.StartBlocking()
}

func loadData() {
	var userSet []model.User
	if db == nil {
		return
	}
	db.Where("redis_push = true").Find(&userSet)

	if len(userSet) > 0 {
		pushToRedis(userSet)
	}
}

func pushToRedis(userSet []model.User) {
	fmt.Println("pushing new data %T", userSet)
	for _, v := range userSet {
		config.LogToFile(fmt.Sprintf("pushing new data %T : %v\n", v, v))
		jsonUser, err := json.Marshal(v)
		if err != nil {
			config.LogToFile(fmt.Sprintf("Error: %v\n", err))
		}

		err = redisInst.Set(ctx, v.Nama, jsonUser, 0).Err()
		if err != nil {
			config.LogToFile(fmt.Sprintf("Error: %v\n", err))
		}

		err = db.Model(&v).Update("redis_push", false).Error
		if err != nil {
			config.LogToFile(fmt.Sprintf("Error: %v\n", err))
		}
	}

	err := redisInst.Close()
	if err != nil {
		config.LogToFile(fmt.Sprintf("Error: %v", err))
	}
}
