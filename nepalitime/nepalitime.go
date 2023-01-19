// contains models/structs related to the package
package nepalitime

import (
	"time"

	"github.com/sugat009/nepali/dateConverter"
)

type NepaliTime struct {
	// note that these members represent nepali values
	year        int16
	month       int8
	day         int8
	englishTime time.Time
}

// Now returns the current nepali time.
func Now() NepaliTime {
	englishTime := time.Now()
	enYear, enMonth, enDay := englishTime.Date()
	year, month, day := dateConverter.EnglishToNepali(enYear, int(enMonth), enDay)
	return NepaliTime{int16(year), int8(month), int8(day), englishTime}
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
