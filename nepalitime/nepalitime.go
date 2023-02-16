// Package nepalitime provides nepali time functionality equivalent to time package.
package nepalitime

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
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
func (obj *NepaliTime) GetEnglishTime() *time.Time {
	return obj.englishTime
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

// formats the nepalitime object into the passed format
func (obj *NepaliTime) Format(format string) (string, error) {
	index, num, timeStr := 0, len(format), ""

	for index < num {
		char := string(format[index])
		index++

		if char == "%" && index < num {
			char = string(format[index])

			if char == "%" {
				timeStr += char
			} else if char == "-" {
				specialChar := char

				if (index + 1) < num {
					index++
					char = string(format[index])
					res, err := obj.getFormatString(specialChar + char)
					if err != nil {
						return "", errors.New("error while formatting NepaliTime with given format")
					}
					timeStr += res
				}
			} else {
				res, err := obj.getFormatString(char)
				if err != nil {
					return "", errors.New("error while formatting NepaliTime with given format")
				}
				timeStr += res
			}
			index++
		} else {
			timeStr += char
		}
	}

	return timeStr, nil
}

// utility method that operates based on the type of directive
func (obj *NepaliTime) getFormatString(directive string) (string, error) {
	switch directive {
	case "d":
		return obj.day_(), nil
	case "-d":
		return obj.dayNonzero(), nil
	case "m":
		return obj.monthNumber(), nil
	case "-m":
		return obj.monthNumberNonzero(), nil
	case "y":
		return obj.yearHalf(), nil
	case "Y":
		return obj.yearFull(), nil
	case "H":
		return obj.hour24(), nil
	case "-H":
		return obj.hour24Nonzero(), nil
	case "I":
		return obj.hour12(), nil
	case "-I":
		return obj.hour12Nonzero(), nil
	case "p":
		return obj.ampm(), nil
	case "M":
		return obj.minute(), nil
	case "-M":
		return obj.minuteNonzero(), nil
	case "S":
		return obj.second(), nil
	case "-S":
		return obj.secondNonzero(), nil
	case "f":
		return obj.nanosecond_(), nil
	case "-f":
		return obj.nanosecondNonZero(), nil
	default:
		return "", errors.New("error while getting format string for passed directive")
	}
}

// %d
func (obj *NepaliTime) day_() string {
	day := strconv.Itoa(obj.day)

	if len(day) < 2 {
		day = "0" + day
	}

	return day
}

// -d
func (obj *NepaliTime) dayNonzero() string {
	day := strconv.Itoa(obj.day)

	return day
}

// %m
func (obj *NepaliTime) monthNumber() string {
	month := strconv.Itoa(obj.month)

	if len(month) < 2 {
		month = "0" + month
	}

	return month
}

// %-m
func (obj *NepaliTime) monthNumberNonzero() string {
	month := strconv.Itoa(obj.month)

	return month
}

// %y
func (obj *NepaliTime) yearHalf() string {
	year := strconv.Itoa(obj.year)

	return year[2:]
}

// %Y
func (obj *NepaliTime) yearFull() string {
	return strconv.Itoa(obj.year)
}

// %H
func (obj *NepaliTime) hour24() string {
	hour := strconv.Itoa(obj.Hour())

	if len(hour) < 2 {
		hour = "0" + hour
	}

	return hour
}

// %-H
func (obj *NepaliTime) hour24Nonzero() string {
	return strconv.Itoa(obj.Hour())
}

// %I
func (obj *NepaliTime) hour12() string {
	hour := obj.Hour()

	if hour > 12 {
		hour -= 12
	}
	if hour == 0 {
		hour = 12
	}

	hourStr := strconv.Itoa(hour)
	if len(hourStr) < 2 {
		hourStr = "0" + hourStr
	}

	return hourStr
}

// %-I
func (obj *NepaliTime) hour12Nonzero() string {
	hour := obj.Hour()

	if hour > 12 {
		hour -= 12
	}
	if hour == 0 {
		hour = 12
	}

	return strconv.Itoa(hour)
}

// %p
func (obj *NepaliTime) ampm() string {
	ampm := "AM"

	if obj.Hour() > 12 {
		ampm = "PM"
	}

	return ampm
}

// %M
func (obj *NepaliTime) minute() string {
	minute := strconv.Itoa(obj.Minute())

	if len(minute) < 2 {
		minute = "0" + minute
	}

	return minute
}

// %-M
func (obj *NepaliTime) minuteNonzero() string {
	return strconv.Itoa(obj.Minute())
}

// %s
func (obj *NepaliTime) second() string {
	second := strconv.Itoa(obj.Second())

	if len(second) < 2 {
		second = "0" + second
	}

	return second
}

// %-s
func (obj *NepaliTime) secondNonzero() string {
	return strconv.Itoa(obj.Second())
}

// %f
func (obj *NepaliTime) nanosecond_() string {
	nsec := strconv.Itoa(obj.Nanosecond())

	if len(nsec) < 6 {
		nsec = strings.Repeat("0", 6 - len(nsec)) + nsec
	}

	return nsec
}

// %-f
func (obj *NepaliTime) nanosecondNonZero() string {
	return strconv.Itoa(obj.Nanosecond())
}
