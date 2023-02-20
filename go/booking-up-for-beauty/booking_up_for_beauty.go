package booking

import (
	"log"
	"strconv"
	"time"
)

func assertNoErr(err error, message string) {
	if err != nil {
		log.Fatal("Error: " + message)
	}
}

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	layout := "1/02/2006 15:04:05"
	t, err := time.Parse(layout, date)
	assertNoErr(err, "Error parsing time "+date)

	return t
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	layout := "January 2, 2006 15:04:05"
	t, err := time.Parse(layout, date)
	assertNoErr(err, "Error parsing time "+date)

	if t.Before(time.Now()) {
		return true
	} else {
		return false
	}
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	layout := "Monday, January 2, 2006 15:04:05"
	t, err := time.Parse(layout, date)
	assertNoErr(err, "Error parsing time "+date)

	noon := time.Date(t.Year(), t.Month(), t.Day(), 12, 0, 0, 0, time.UTC)
	sixpm := time.Date(t.Year(), t.Month(), t.Day(), 18, 0, 0, 0, time.UTC)

	if t.After(noon) && t.Before(sixpm) {
		return true
	} else {
		return false
	}

}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	layout := "1/2/2006 15:04:05"
	t, err := time.Parse(layout, date)
	assertNoErr(err, "Error parsing time "+date)

	weekday := t.Weekday().String()
	year, month, day := t.Date()
	hour, minute, _ := t.Clock()

	text := "You have an appointment on " + weekday + ", " + month.String() + " " + strconv.Itoa(day) +
		", " + strconv.Itoa(year) + ", at " + strconv.Itoa(hour) + ":" + strconv.Itoa(minute) + "."

	return text
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	layout := "2006-01-02 15:04:05 +0000 UTC"
	anniversary := "2020-09-15 00:00:00 +0000 UTC"
	t, err := time.Parse(layout, anniversary)
	assertNoErr(err, "Error parsing time "+anniversary)

	yearDelta := time.Now().Year() - t.Year()

	newAnniversary := t.AddDate(yearDelta, 0, 0)

	return newAnniversary
}
