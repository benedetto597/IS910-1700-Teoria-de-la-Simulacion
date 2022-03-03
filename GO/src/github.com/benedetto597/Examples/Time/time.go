package main

import (
	"fmt"
	"time"
)

func CurrentTime() {
	currentTime := time.Now()
	fmt.Println("The time is", currentTime)
}

func PortionTime() {
	currentTime := time.Now()
	fmt.Println("The time is", currentTime)

	fmt.Println("The year is", currentTime.Year())
	fmt.Println("The month is", currentTime.Month())
	fmt.Println("The day is", currentTime.Day())
	fmt.Println("The hour is", currentTime.Hour())
	fmt.Println("The minute is", currentTime.Hour())
	fmt.Println("The second is", currentTime.Second())
}

func PortionTimeFormat() {
	currentTime := time.Now()
	fmt.Println("The time is", currentTime)

	fmt.Printf("%d-%d-%d %d:%d:%d\n",
		currentTime.Year(),
		currentTime.Month(),
		currentTime.Day(),
		currentTime.Hour(),
		currentTime.Hour(),
		currentTime.Second())
}

func CurrentDate() {
	theTime := time.Date(2021, 8, 15, 14, 30, 45, 100, time.Local)
	fmt.Println("The time is", theTime)
}

func FormatDate() {
	theTime := time.Date(2021, 8, 15, 14, 30, 45, 100, time.Local)
	fmt.Println("The time is", theTime)

	fmt.Println(theTime.Format("2006-1-2 15:4:5")) // Output: 2021-08-15 14:30:45
}

func FormatDateTwelveHours() {
	theTime := time.Date(2021, 8, 15, 14, 30, 45, 100, time.Local)
	fmt.Println("The time is", theTime)

	fmt.Println(theTime.Format("2006-1-2 15:4:5"))
	fmt.Println(theTime.Format("2006-01-02 03:04:05 pm"))
}

func FormatDateRFC() {
	theTime := time.Date(2021, 8, 15, 14, 30, 45, 100, time.Local)
	fmt.Println("The time is", theTime)

	fmt.Println(theTime.Format(time.RFC3339Nano))
}

func ParseTime() {
	timeString := "2021-08-15 02:30:45"
	theTime, err := time.Parse("2006-01-02 03:04:05", timeString)
	if err != nil {
		fmt.Println("Could not parse time:", err)
	}
	fmt.Println("The time is", theTime)

	fmt.Println(theTime.Format(time.RFC3339Nano))
}

func ParseTimeRFC() {
	timeString := "2021-08-15T14:30:45.0000001-05:00"
	theTime, err := time.Parse(time.RFC3339Nano, timeString)
	if err != nil {
		fmt.Println("Could not parse time:", err)
	}
	fmt.Println("The time is", theTime)

	fmt.Println(theTime.Format(time.RFC3339Nano))
}

func TimeZoneUTC() {
	theTime := time.Date(2021, 8, 15, 14, 30, 45, 100, time.Local)
	fmt.Println("The time is", theTime)
	fmt.Println(theTime.Format(time.RFC3339Nano))

	utcTime := theTime.UTC()
	fmt.Println("The UTC time is", utcTime)
	fmt.Println(utcTime.Format(time.RFC3339Nano))
}

func TimeZoneLocalToUTC() {
	theTime := time.Date(2021, 8, 15, 14, 30, 45, 100, time.Local)
	fmt.Println("The time is", theTime)
	fmt.Println(theTime.Format(time.RFC3339Nano))

	utcTime := theTime.UTC()
	fmt.Println("The UTC time is", utcTime)
	fmt.Println(utcTime.Format(time.RFC3339Nano))

	localTime := utcTime.Local()
	fmt.Println("The Local time is", localTime)
	fmt.Println(localTime.Format(time.RFC3339Nano))
}

func CompareTimes() {
	firstTime := time.Date(2021, 8, 15, 14, 30, 45, 100, time.UTC)
	fmt.Println("The first time is", firstTime)

	secondTime := time.Date(2021, 12, 25, 16, 40, 55, 200, time.UTC)
	fmt.Println("The second time is", secondTime)

	fmt.Println("First time before second?", firstTime.Before(secondTime))
	fmt.Println("First time after second?", firstTime.After(secondTime))

	fmt.Println("Second time before first?", secondTime.Before(firstTime))
	fmt.Println("Second time after first?", secondTime.After(firstTime))
}

func CompareTimes() {
	firstTime := time.Date(2021, 8, 15, 14, 30, 45, 100, time.UTC)
	fmt.Println("The first time is", firstTime)

	secondTime := time.Date(2021, 12, 25, 16, 40, 55, 200, time.UTC)
	fmt.Println("The second time is", secondTime)

	fmt.Println("First time before second?", firstTime.Before(secondTime))
	fmt.Println("First time after second?", firstTime.After(secondTime))

	fmt.Println("Second time before first?", secondTime.Before(firstTime))
	fmt.Println("Second time after first?", secondTime.After(firstTime))
}

func CompareTimesSubstraction() {
	firstTime := time.Date(2021, 8, 15, 14, 30, 45, 100, time.UTC)
	fmt.Println("The first time is", firstTime)

	secondTime := time.Date(2021, 12, 25, 16, 40, 55, 200, time.UTC)
	fmt.Println("The second time is", secondTime)

	fmt.Println("Duration between first and second time is", firstTime.Sub(secondTime))
	fmt.Println("Duration between second and first time is", secondTime.Sub(firstTime))
}

func AddTime() {
	toAdd := 1 * time.Hour
	fmt.Println("1:", toAdd)

	toAdd += 1 * time.Minute
	fmt.Println("2:", toAdd)

	toAdd += 1 * time.Second
	fmt.Println("3:", toAdd)
}

func AddSub() {
	oneHourOneMinute := 1*time.Hour + 1*time.Minute
	tenMinutes := 1*time.Hour - 50*time.Minute

	toAdd += 1 * time.Second
	fmt.Println("3:", toAdd)

	toAdd -= 1*time.Minute + 1*time.Second
	fmt.Println("4:", toAdd)
}

func AddDate() {
	theTime := time.Date(2021, 8, 15, 14, 30, 45, 100, time.UTC)
	fmt.Println("The time is", theTime)

	toAdd := 24 * time.Hour
	fmt.Println("Adding", toAdd)

	newTime := theTime.Add(toAdd)
	fmt.Println("The new time is", newTime)
}
