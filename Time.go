package booking

import (
	"fmt"
	"time"
)

/*
Implement the Schedule function to parse a textual representation of an appointment date into the corresponding time.Time format:

Schedule("7/25/2019 13:45:00")
// => 2019-07-25 13:45:00 +0000 UTC
*/
func Schedule(date string) time.Time {
	t, err := time.Parse("1/2/2006 15:04:05", date)
	if err != nil {
		return time.Time{}
	}
	return t
}

/*
Implement the HasPassed function that takes an appointment date and checks if the appointment was somewhere in the past:

HasPassed("July 25, 2019 13:45:00")
// => true
*/

func HasPassed(date string) bool {
	t, err := time.Parse("January 2, 2006 15:04:05", date)
	if err != nil {
		return false
	}
	return t.Before(time.Now())
}

/*
Implement the IsAfternoonAppointment function that takes an appointment date and checks if the appointment is in the afternoon (>= 12:00 and < 18:00):

IsAfternoonAppointment("Thursday, July 25, 2019 13:45:00")
// => true
*/

func IsAfternoonAppointment(date string) bool {
	t, err := time.Parse("Monday, January 2, 2006 15:04:05", date)
	if err != nil {
		return false
	}
	hour := t.Hour()
	return hour >= 12 && hour < 18
}

/*
Implement the Description function that takes an appointment date and returns a description of that date and time:

Description("7/25/2019 13:45:00")
// => "You have an appointment on Thursday, July 25, 2019, at 13:45.
*/

func Description(date string) string {
	t, err := time.Parse("1/2/2006 15:04:05", date)
	if err != nil {
		return "Invalid appointment date format"
	}
	formattedDate := t.Format("Monday, January 2, 2006")
	formattedTime := t.Format("15:04")
	description := fmt.Sprintf("You have an appointment on %s, at %s.", formattedDate, formattedTime)
	
	return description
}

/*
Implement the AnniversaryDate function that returns the anniversary date of the salon's opening for the current year in UTC.

Assuming the current year is 2020:

AnniversaryDate()

// => 2020-09-15 00:00:00 +0000 UTC
*/

func AnniversaryDate() time.Time {
	currentTime := time.Now().UTC()
	currentYear := currentTime.Year()
	anniversaryDate := time.Date(
		currentYear, // Year: current year
		time.September, // Month: September
		15, // Day: 15th
		0, // Hour: 0
		0, // Minute: 0
		0, // Second: 0
		0, // Nanosecond: 0
		time.UTC, // Location: UTC
	)
	
	return anniversaryDate
}
