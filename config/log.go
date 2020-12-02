package config

import (
	"fmt"
	"log"
	"os"
	"time"
)

func loadFile() *os.File {
	file, err := os.OpenFile("logs/log", os.O_RDWR, 0777)
	if err != nil {
		file, err = os.Create("logs/log")
		if err != nil {
			panic(fmt.Errorf("Error: %v\n", err))
		}
	}
	return file
}

func LogToFile(content string) {
	file := loadFile()
	currentTime := time.Now().Format("2006/01/02 15:04:05 Z07")
	content = fmt.Sprintln(currentTime, content)
	log.SetOutput(file)
	log.Println(content)

	err := file.Close()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
