package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/2ndSilencerz/redis-data-pusher/config"
	"github.com/2ndSilencerz/redis-data-pusher/database"
	"github.com/2ndSilencerz/redis-data-pusher/model"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func TestRead(t *testing.T) {

	// insert some random data to database
	config.PrintToConsole("inserting")
	timex := time.Now().Nanosecond()
	rand.Seed(int64(timex))
	rng := rand.Intn(999)
	userInsert := model.User{
		Nama:      "testing" + strconv.Itoa(timex),
		Umur:      rng,
		RedisPush: true,
	}
	dbInst := database.InitDB()
	err := dbInst.Create(userInsert).Error
	if err != nil {
		t.Errorf("Error: %v", err)
		return
	}

	// let it sleep for interval time so main program can push to redis
	config.PrintToConsole(fmt.Sprintf("sleep for %v second(s)", config.GetCronInterval()))
	interval := config.GetCronInterval()
	durations, err := time.ParseDuration(strconv.Itoa(interval) + "s")
	if err != nil {
		t.Errorf("Error: %v", err)
		return
	}
	time.Sleep(durations)

	// get data with exact key
	config.PrintToConsole("fetching")
	redisInstance := database.InitRedis()
	testContext := context.Background()
	jsonGotten := redisInstance.Get(testContext, userInsert.Nama).Val()
	//if err != nil {
	//	t.Errorf("Error: %v", err)
	//	return
	//}
	fmt.Println(jsonGotten)
	var userGotten model.User
	err = json.Unmarshal([]byte(jsonGotten), &userGotten)
	if err != nil {
		t.Errorf("Error: %v", err)
		return
	}

	if userGotten.Umur != userInsert.Umur {
		t.Errorf("Expected umur %v, got %v", userInsert.Umur, userGotten.Umur)
	}

	err = redisInstance.Close()
	if err != nil {
		t.Errorf("Error: %v", err)
	}
}
