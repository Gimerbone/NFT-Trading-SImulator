package utils

import (
	"time"
)

var timezone = time.FixedZone("GMT+7", 7*60*60)

func CurrentDate() string {
	// returns a string of current date in this format
	// Sunday May 18, 2025
	currentTime := time.Now().In(timezone)
	return currentTime.Format("Monday January 2, 2006")
}
