package utils

import (
	"fmt"
	"time"
)

func GetCurrentMonthAndYear() map[string]string {
	// Get the current time
	currentTime := time.Now()

	// Extract month and year from the current time
	date := fmt.Sprintf("%02d", currentTime.Day())
	month := currentTime.Month().String()
	year := fmt.Sprintf("%d", currentTime.Year())

	// Create a map to store the results
	result := map[string]string{
		"date":  date,
		"month": month,
		"year":  year,
	}

	return result
}
