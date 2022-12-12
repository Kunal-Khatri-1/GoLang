package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Getting hours, days, minutes, seconds and, difference between two dates [future and past] ")

	// returns time.Location for a place for which the time is to be fetched
	loc, _ := time.LoadLocation("Asia/Kolkata")
	fmt.Println(loc)
	// getting the current time in the passed location
	now := time.Now().In(loc)
	fmt.Println(now)

	// getting past date
	pastDate := time.Date(2003, time.November, 14, 4, 30, 0, 0, loc)
	fmt.Println(pastDate)

	// difference in the two dates
	pastDiff := now.Sub(pastDate)
	fmt.Println(pastDiff)

	// getting hours, mins, seconds
	hrs := pastDiff.Hours()
	mins := pastDiff.Minutes()
	secs := pastDiff.Seconds()

	fmt.Println(hrs, mins, secs)

	// getting future date
	futureDate := time.Date(2023, time.January, 1, 1, 1, 1, 1, loc)
	fmt.Println(futureDate)

	futureDiff := futureDate.Sub(now)
	fmt.Println(futureDiff)

	hrs = futureDiff.Hours()
	mins = futureDiff.Minutes()
	secs = futureDiff.Seconds()

	fmt.Println(hrs, mins, secs)
}
