package main

import (
	"bufio"
	"github.com/2ndSilencerz/redis-data-pusher/config"
	"os"
	"testing"
)

func TestLog(t *testing.T) {
	contents := "testing 1"
	config.LogToFile(contents)
	// exempt +07 (or Z07) from time string since logging dont use it
	//bytesString := []byte(config.InstantTimeString())
	//contents =  string(bytesString[:19]) + " " + contents
	contents = config.InstantTimeString() + " " + contents

	file, err := os.OpenFile("logs/log", os.O_RDONLY, 0777)
	if err != nil {
		t.Errorf("Error: %v", err)
		return
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	var contentGotten string
	for fileScanner.Scan() {
		contentGotten = fileScanner.Text()
	}

	if contentGotten != contents {
		t.Errorf("Expected %v got %v", contents, contentGotten)
	}

	err = file.Close()
	if err != nil {
		t.Errorf("Error: %v", err)
	}
}
