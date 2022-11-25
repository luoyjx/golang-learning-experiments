package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()
	startTime := time.Date(
		currentTime.Year(),
		currentTime.Month(),
		currentTime.Day(),
		0,
		0,
		0,
		0,
		currentTime.Location(),
	)
	fmt.Println(startTime.UTC().Format("2006-01-02T15:04:05Z"))
}
