package nepalitime

import (
	"fmt"
	"time"

	"github.com/opensource-nepal/go-nepali/constants"
	"github.com/opensource-nepal/go-nepali/dateConverter"
)

// Date returns the Time corresponding to
//
// yyyy-mm-dd hh:mm:ss + nsec nanoseconds
func Date(year, month, day, hour, min, sec, nsec int) (*NepaliTime, error) {
	englishDate, err := dateConverter.NepaliToEnglish(year, month, day)
	if err != nil {
		return nil, err
	}

	englishTime := time.Date(englishDate[0], time.Month(englishDate[1]), englishDate[2],
		hour, min, sec, nsec, GetNepaliLocation())
	return &NepaliTime{year, month, day, &englishTime}, nil
}

// FromEnglishTime Converts Time object to NepaliTime
func FromEnglishTime(englishTime time.Time) (*NepaliTime, error) {
	englishTime = englishTime.In(GetNepaliLocation())
	enYear, enMonth, enDay := englishTime.Date()
	englishDate, err := dateConverter.EnglishToNepali(enYear, int(enMonth), enDay)
	if err != nil {
		return nil, err
	}

	return &NepaliTime{englishDate[0], englishDate[1], englishDate[2], &englishTime}, nil
}

// Now returns the current nepali time.
// this function should always work
// and should not return nil in normal circumstances
func Now() *NepaliTime {
	now, _ := FromEnglishTime(GetCurrentEnglishTime())

	return now
}

// adds zero on the number if the number is less than 10
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

// GetCurrentEnglishTime Gets current English date along with time level precision.
// Current Time of Asia/Kathmandu
func GetCurrentEnglishTime() time.Time {
	return time.Now().In(GetNepaliLocation())
}

// GetNepaliLocation Returns location for Asia/Kathmandu (constants.Timezone)
func GetNepaliLocation() *time.Location {
	location, _ := time.LoadLocation(constants.Timezone)
	// ignoring error since location with Asia/Kathmandu will not fail.
	return location
}
