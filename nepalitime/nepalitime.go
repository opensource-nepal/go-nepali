// Package nepalitime provides nepali time functionality equivalent to time package.
package nepalitime

import (
	"fmt"
	"time"

	"github.com/sugat009/nepali/dateConverter"
)

// NOTE:
// Only Date() and FromTime() are supposed to create NepaliTime object

type NepaliTime struct {
	// note that these members represent nepali values
	year        int
	month       int
	day         int
	englishTime time.Time
}

// Converts single digit number into two digits.
// Adds zero on the number if the number is less than 10.
//
// eg. 19 => 19 and 8 => 08
//
// Note: if number is 144 it will return 144
func twoDigitNumber(number int) string {
	if number < 10 && number >= 0 {
		return fmt.Sprintf("0%d", number)
	}
	return fmt.Sprint(number)
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

func (nt NepaliTime) ToTime() time.Time {
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

// Date returns the Time corresponding to
//
// yyyy-mm-dd hh:mm:ss + nsec nanoseconds
func Date(year, month, day, hour, min, sec, nsec int) NepaliTime {
	enYear, enMonth, enDay := dateConverter.NepaliToEnglish(year, month, day)
	location, _ := time.LoadLocation("Asia/Kathmandu")
	englishTime := time.Date(enYear, time.Month(enMonth), enDay, hour, min, sec, nsec, location)
	return NepaliTime{year, month, day, englishTime}
}

// Converts Time object to NepaliTime
func FromTime(t time.Time) NepaliTime {
	enYear, enMonth, enDay := t.Date()
	year, month, day := dateConverter.EnglishToNepali(enYear, int(enMonth), enDay)
	return NepaliTime{year, month, day, t}
}

// Now returns the current nepali time.
func Now() NepaliTime {
	return FromTime(time.Now())
}
