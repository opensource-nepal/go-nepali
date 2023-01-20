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

func (nt NepaliTime) String() string {
	h, m, s := nt.englishTime.Clock()
	return fmt.Sprintf("%d-%d-%d %d:%d:%d", nt.year, nt.month, nt.day, h, m, s)
}

// Date returns the year, month, and day
func (nt NepaliTime) Date() (year, month, day int) {
	year = int(nt.year)
	month = int(nt.month)
	day = int(nt.day)
	return
}

// Clock returns the hour, minute, and second of the day.
func (nt NepaliTime) Clock() (hour, min, sec int) {
	return nt.englishTime.Clock()
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
