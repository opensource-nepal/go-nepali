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
func (obj *NepaliTime) String() string {
	h, m, s := obj.Clock()
	return fmt.Sprintf(
		"%d-%s-%s %s:%s:%s",
		obj.year,
		twoDigitNumber(obj.month),
		twoDigitNumber(obj.day),
		twoDigitNumber(h),
		twoDigitNumber(m),
		twoDigitNumber(s),
	)
}

// Get's the corresponding english time
func (obj *NepaliTime) GetEnglishTime() time.Time {
	return *obj.englishTime
}

// Date returns the year, month, and day
func (obj *NepaliTime) Date() (year, month, day int) {
	return obj.year, obj.month, obj.day
}

// Year returns the year in which nepalitime occurs.
func (obj *NepaliTime) Year() int {
	return obj.year
}

// Month returns the month of the year.
func (obj *NepaliTime) Month() int {
	return obj.month
}

// Day returns the day of the month.
func (obj *NepaliTime) Day() int {
	return obj.day
}

// Weekday returns the day of the week.
// Sunday = 0,
// Monday = 1,
// Saturday = 6
func (obj *NepaliTime) Weekday() time.Weekday {
	return obj.englishTime.Weekday()
}

// Clock returns the hour, minute, and second of the day.
func (obj *NepaliTime) Clock() (hour, min, sec int) {
	return obj.englishTime.Clock()
}

// Hour returns the hour within the day, in the range [0, 23].
func (obj *NepaliTime) Hour() int {
	return obj.englishTime.Hour()
}

// Minute returns the minute offset within the hour, in the range [0, 59].
func (obj *NepaliTime) Minute() int {
	return obj.englishTime.Minute()
}

// Second returns the second offset within the minute, in the range [0, 59].
func (obj *NepaliTime) Second() int {
	return obj.englishTime.Second()
}

// Nanosecond returns the nanosecond offset within the second,
// in the range [0, 999999999].
func (obj *NepaliTime) Nanosecond() int {
	return obj.englishTime.Nanosecond()
}

// Format formats the nepalitime object into the passed format
func (obj *NepaliTime) Format(format string) (string, error) {
	formatter := NewFormatter(obj)

	formattedNepaliTime, err := formatter.Format(format)
	if err != nil {
		return "", err
	}

	return formattedNepaliTime, nil
}
