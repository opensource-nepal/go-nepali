// Package nepalitime provides nepali time functionality equivalent to time package.
package nepalitime

import (
	"fmt"
	"time"
)

// NOTE:
// Only Date() and FromEnglishTime() from utils are supposed to create NepaliTime object

type NepaliTime struct {
	// note that these members represent nepali values
	year        int
	month       int
	day         int
	englishTime *time.Time
}

// String returns a string representing the duration in the form "2079-10-06 01:00:05"
func (nt NepaliTime) String() string {
	h, m, s := nt.englishTime.Clock()
	return fmt.Sprintf(
		"%d-%s-%s %s:%s:%s",
		nt.year,
		twoDigitNumber(nt.month),
		twoDigitNumber(nt.day),
		twoDigitNumber(h),
		twoDigitNumber(m),
		twoDigitNumber(s),
	)
}

// Get's the corresponding english time
func (nt NepaliTime) GetEnglishTime() *time.Time {
	return nt.englishTime
}

// Date returns the year, month, and day
func (nt NepaliTime) Date() (year, month, day int) {
	return nt.year, nt.month, nt.day
}

// Year returns the year in which nepalitime occurs.
func (nt NepaliTime) Year() int {
	return nt.year
}

// Month returns the month of the year.
func (nt NepaliTime) Month() int {
	return nt.month
}

// Day returns the day of the month.
func (nt NepaliTime) Day() int {
	return nt.day
}

// Weekday returns the day of the week.
// Sunday = 0,
// Monday = 1,
// Saturday = 6
func (nt NepaliTime) Weekday() time.Weekday {
	return nt.englishTime.Weekday()
}

// Clock returns the hour, minute, and second of the day.
func (nt NepaliTime) Clock() (hour, min, sec int) {
	return nt.englishTime.Clock()
}

// Hour returns the hour within the day, in the range [0, 23].
func (nt NepaliTime) Hour() int {
	return nt.englishTime.Hour()
}

// Minute returns the minute offset within the hour, in the range [0, 59].
func (nt NepaliTime) Minute() int {
	return nt.englishTime.Minute()
}

// Second returns the second offset within the minute, in the range [0, 59].
func (nt NepaliTime) Second() int {
	return nt.englishTime.Second()
}

// Nanosecond returns the nanosecond offset within the second,
// in the range [0, 999999999].
func (nt NepaliTime) Nanosecond() int {
	return nt.englishTime.Nanosecond()
}
