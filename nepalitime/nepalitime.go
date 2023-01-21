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
	year        int16
	month       int8
	day         int8
	englishTime time.Time
}

// adds zero on the number if the number is less than 10
//
// eg. 19 => 19 but 8 => 08
func addZeroOnNumber(number int) string {
	if number < 10 {
		return "0" + fmt.Sprint(number)
	}
	return fmt.Sprint(number)
}

// String returns a string representing the duration in the form "2079-10-06 01:00:05"
func (nt NepaliTime) String() string {
	h, m, s := nt.englishTime.Clock()
	return fmt.Sprintf(
		"%d-%s-%s %s:%s:%s",
		nt.year,
		addZeroOnNumber(int(nt.month)),
		addZeroOnNumber(int(nt.day)),
		addZeroOnNumber(h),
		addZeroOnNumber(m),
		addZeroOnNumber(s),
	)
}

func (nt NepaliTime) ToTime() time.Time {
	return nt.englishTime
}

// Date returns the year, month, and day
func (nt NepaliTime) Date() (year, month, day int) {
	year = int(nt.year)
	month = int(nt.month)
	day = int(nt.day)
	return
}

// Year returns the year in which nepalitime occurs.
func (nt NepaliTime) Year() int {
	return int(nt.year)
}

// Month returns the month of the year.
func (nt NepaliTime) Month() int {
	return int(nt.month)
}

// Day returns the day of the month.
func (nt NepaliTime) Day() int {
	return int(nt.day)
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
	return NepaliTime{int16(year), int8(month), int8(day), englishTime}
}

// Converts Time object to NepaliTime
func FromTime(t time.Time) NepaliTime {
	enYear, enMonth, enDay := t.Date()
	year, month, day := dateConverter.EnglishToNepali(enYear, int(enMonth), enDay)
	return NepaliTime{int16(year), int8(month), int8(day), t}
}

// Now returns the current nepali time.
func Now() NepaliTime {
	return FromTime(time.Now())
}
