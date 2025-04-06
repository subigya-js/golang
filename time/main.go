package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()
	fmt.Println("Current Time:", currentTime.Format("2006 January 2, 3:04:05 PM, Monday"))

	// Parse a string to date/time
	dateString := "2023-10-01"
	formattedDate, _ := time.Parse("2006-01-02", dateString)
	fmt.Println("Parsed Date:", formattedDate.Format("2006-01-02"))

	fmt.Println("Add one more day in current time:", currentTime.AddDate(0, 0, 1).Format("2006-01-02"))
	fmt.Println("Add one more hour in current time:", currentTime.Add(time.Hour).Format("2006-01-02 03:04:05 PM"))
}
