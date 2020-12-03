package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/2ndSilencerz/redis-data-pusher/config"
	"github.com/2ndSilencerz/redis-data-pusher/database"
	"github.com/2ndSilencerz/redis-data-pusher/model"
	"github.com/go-co-op/gocron"
	"log"
	"strconv"
	"time"
)

var ctx = context.Background()

func main() {
	interval := config.GetCronInterval()
	durations, err := time.ParseDuration(strconv.Itoa(interval) + "s")
	if err != nil {
		msg := fmt.Sprintf("Error: %v", err)
		config.LogToFile(msg)
		log.Fatalf(msg)
	}
	s1 := gocron.NewScheduler(time.UTC)
	_, err = s1.Every(uint64(durations.Seconds())).Second().Do(loadData)
	if err != nil {
		msg := fmt.Sprintf("Error: %v", err)
		config.LogToFile(msg)
		log.Fatalf(msg)
	}

	config.PrintToConsole(fmt.Sprintf("starting scheduler with interval %v second(s)",
		durations.Seconds()))
	s1.StartBlocking()
}

func loadData() {
	config.PrintToConsole("reading database")
	db := database.InitDB()
	var userSet []model.User
	db.Where("redis_push = true").Find(&userSet)
	database.CloseDB(db)

	if len(userSet) > 0 {
		pushToRedis(userSet)
	} else {
		config.PrintToConsole(fmt.Sprintf("no new data found"))
	}
}

func pushToRedis(userSet []model.User) {
	db := database.InitDB()
	defer database.CloseDB(db)
	redisInst := database.InitRedis()
	config.PrintToConsole(fmt.Sprintf("pushing new dataset %T\n", userSet))
	for _, v := range userSet {
		config.LogToFile(fmt.Sprintf("pushing new data %T : %v", v, v.Nama))
		jsonUser, err := json.Marshal(v)
		if err != nil {
			config.LogToFile(fmt.Sprintf("Error: %v", err))
			continue
		}

		err = redisInst.Set(ctx, v.Nama, jsonUser, 0).Err()
		if err != nil {
			config.LogToFile(fmt.Sprintf("Error: %v", err))
			continue
		}

		err = db.Model(&v).Where("nama = ?", v.Nama).Update("redis_push", false).Error
		if err != nil {
			config.LogToFile(fmt.Sprintf("Error: %v", err))
			continue
		}
	}

	err := redisInst.Close()
	if err != nil {
		config.LogToFile(fmt.Sprintf("Error: %v", err))
	}
}
