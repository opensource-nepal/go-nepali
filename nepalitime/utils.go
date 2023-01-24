package nepalitime

import (
	"fmt"
	"time"

	"github.com/sugat009/nepali/constants"
	"github.com/sugat009/nepali/dateConverter"
)

// NOTE:
// Only Date() and FromTime() are supposed to create NepaliTime object

// Date returns the Time corresponding to
//
// yyyy-mm-dd hh:mm:ss + nsec nanoseconds
func Date(year, month, day, hour, min, sec, nsec int) (*NepaliTime, error) {
	englishDate, err := dateConverter.NepaliToEnglish(year, month, day)
	if err != nil {
		return nil, err
	}

	location, _ := time.LoadLocation(constants.Timezone)
	englishTime := time.Date(englishDate[0], time.Month(englishDate[1]), englishDate[2], hour, min, sec, nsec, location)
	return &NepaliTime{year, month, day, &englishTime}, nil
}

// Converts Time object to NepaliTime
func FromTime(t time.Time) (*NepaliTime, error) {
	enYear, enMonth, enDay := t.Date()
	englishDate, err := dateConverter.EnglishToNepali(enYear, int(enMonth), enDay)
	if err != nil {
		return nil, err
	}

	return &NepaliTime{englishDate[0], englishDate[1], englishDate[2], &t}, nil
}

// Now returns the current nepali time.
// this function should always work
// and should not return nil in normal circumstances
func Now() *NepaliTime {
	now, _ := FromTime(GetCurrentEnglishTime(""))

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

// Gets current English date(tarik) along with time level precision
// Also, takes in a format parameter which can define how you want
// your output to be. If you want the default format then pass in
// an empty string
func GetCurrentEnglishTime(format string) time.Time {
	location, _ := time.LoadLocation(constants.Timezone)
	currentTime := time.Now().In(location)

	if format != "" {
		currentTime.Format(format)
	}

	return currentTime
}
