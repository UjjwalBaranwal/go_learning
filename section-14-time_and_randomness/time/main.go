package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Printf("Current Time : %s\n", now)
	fmt.Printf("Year : %d , Month : %s , Day : %d\n", now.Year(), now.Month(), now.Day())
	fmt.Printf("Hour : %d , Minute : %d , second : %d \n", now.Hour(), now.Minute(), now.Second())
	fmt.Printf("Weekday: %s\n", now.Weekday().String())
	goLaunchDate := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Printf("\nGo's approximate launch date (UTC): %s\n", goLaunchDate)
	oneSecond := time.Second
	fiveMinutes := 5 * time.Minute
	oneHourThirtyMinutes := time.Hour + 30*time.Minute
	fmt.Printf("\nOne second: %v\n", oneSecond)
	fmt.Printf("Five minutes in nanoseconds: %d ns\n", fiveMinutes)
	fmt.Printf("Five minutes as string: %s\n", fiveMinutes.String())
	fmt.Printf("1h30m as string: %s\n", oneHourThirtyMinutes.String())

	futureTime := now.Add(2 * time.Hour)
	fmt.Printf("Current time + 2 hours: %s\n", futureTime)

	pastTime := now.Add(-30 * time.Minute)
	fmt.Printf("Current time + 30 minutes: %s\n", pastTime)

	durationSinceGoLaunch := now.Sub(goLaunchDate)

	fmt.Printf("Time since Go launch (approx): %s\n", durationSinceGoLaunch)
	fmt.Printf("Approx hours since Go launch: %.0f hours\n", durationSinceGoLaunch.Hours())
	if futureTime.After(now) {
		fmt.Println("FutureTime is indeed after now")
	}
	if pastTime.Before(now) {
		fmt.Println("PastTime is indeed before now")
	}
	fmt.Printf("\nSleeping for 30 second\n")
	shortDuration := 30 * time.Second
	time.Sleep(shortDuration)
	fmt.Println("Awake after sleep")
}
