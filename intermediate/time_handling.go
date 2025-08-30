package intermediate

import (
	"fmt"
	"time"
)

func time_handling() {
	// Current local time
	fmt.Println(time.Now())

	//Specific time
	SpecificTime := time.Date(2003, time.October, 15, 1, 15, 3, 0, time.Now().Local().Location())
	fmt.Println(SpecificTime)

	//Parse time
	parsedTime, _ := time.Parse("2006-01-02", "2003-10-15") // Mon Jan 2 15:04:05 MST 2006
	parsedTime1, _ := time.Parse("06-01-02", "03-10-15") // Mon Jan 2 15:04:05 MST 2006
	fmt.Println(parsedTime)
	fmt.Println(parsedTime1)

	// Format Time
	t := time.Now()
	fmt.Println("Formatted time:", t.Format("06-01-02 15:04"))

	// Add Time
	oneDayLater := t.Add(time.Hour * 24)
	fmt.Println(oneDayLater)
	fmt.Println(oneDayLater.Weekday())

	// Round Time
	fmt.Println("Rounded Time:", t.Round(time.Hour))

	loc, _ := time.LoadLocation("Asia/Kolkata")
	t = time.Date(2024, time.July, 8, 14, 16, 40, 00, time.UTC)

	// Convert this to the specific time zone
	tLocal := t.In(loc)
	roundedTime := t.Round(time.Hour)
	roundedTimeLocal := roundedTime.In(loc)
	fmt.Println("Original Time UTC:", t)
	fmt.Println("Original Time Local:", tLocal)
	fmt.Println("Rounded Time UTC:", roundedTime)
	fmt.Println("Rounded Time Local:", roundedTimeLocal)

}
