package main

import (
	"fmt"
	"time"
)

func main() {
	string_to_date("2021-08-15 02:30:45")
}

func string_to_date(date string) time.Time {
	theTime, err := time.Parse("2006-01-02 03:04:05", date)
	if err != nil {
		fmt.Println("Could not parse time:", err)
	}
	fmt.Println("The time is", theTime)
	return theTime
}
