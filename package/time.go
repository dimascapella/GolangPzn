package main

import (
	"fmt"
	"time"
)

func main() {
	timeNow := time.Now()
	fmt.Println(timeNow.Local())

	timeDate := time.Date(2022, 01, 01, 01, 01, 01, 01, time.UTC)
	fmt.Println(timeDate)

	timeParse, _ := time.Parse(time.RFC3339, "2022-01-07T20:20:20Z")
	fmt.Println(timeParse)
}
