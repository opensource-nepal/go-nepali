package nepalitime

import (
	"strconv"
	"strings"
)

func NewFormatter(nepaliTime *NepaliTime) *NepaliFormatter {
	return &NepaliFormatter{nepaliTime: nepaliTime}
}

type NepaliFormatter struct {
	nepaliTime *NepaliTime
}

func (obj *NepaliFormatter) Format(format string) string {
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
					res := obj.getFormatString(specialChar + char)
					timeStr += res
				}
			} else {
				res := obj.getFormatString(char)
				timeStr += res
			}
			index++
		} else {
			timeStr += char
		}
	}

	return timeStr
}

// utility method that operates based on the type of directive
func (obj *NepaliFormatter) getFormatString(directive string) string {
	switch directive {
	case "d":
		return obj.day_()
	case "-d":
		return obj.dayNonzero()
	case "m":
		return obj.monthNumber()
	case "-m":
		return obj.monthNumberNonzero()
	case "y":
		return obj.yearHalf()
	case "Y":
		return obj.yearFull()
	case "H":
		return obj.hour24()
	case "-H":
		return obj.hour24Nonzero()
	case "I":
		return obj.hour12()
	case "-I":
		return obj.hour12Nonzero()
	case "p":
		return obj.ampm()
	case "M":
		return obj.minute()
	case "-M":
		return obj.minuteNonzero()
	case "S":
		return obj.second()
	case "-S":
		return obj.secondNonzero()
	case "f":
		return obj.nanosecond_()
	case "-f":
		return obj.nanosecondNonZero()
	default:
		// if not match return the directive
		return "%" + directive
	}
}

// %d
func (obj *NepaliFormatter) day_() string {
	day := strconv.Itoa(obj.nepaliTime.day)

	if len(day) < 2 {
		day = "0" + day
	}

	return day
}

// -d
func (obj *NepaliFormatter) dayNonzero() string {
	day := strconv.Itoa(obj.nepaliTime.day)

	return day
}

// %m
func (obj *NepaliFormatter) monthNumber() string {
	month := strconv.Itoa(obj.nepaliTime.month)

	if len(month) < 2 {
		month = "0" + month
	}

	return month
}

// %-m
func (obj *NepaliFormatter) monthNumberNonzero() string {
	month := strconv.Itoa(obj.nepaliTime.month)

	return month
}

// %y
func (obj *NepaliFormatter) yearHalf() string {
	year := strconv.Itoa(obj.nepaliTime.year)

	return year[2:]
}

// %Y
func (obj *NepaliFormatter) yearFull() string {
	return strconv.Itoa(obj.nepaliTime.year)
}

// %H
func (obj *NepaliFormatter) hour24() string {
	hour := strconv.Itoa(obj.nepaliTime.Hour())

	if len(hour) < 2 {
		hour = "0" + hour
	}

	return hour
}

// %-H
func (obj *NepaliFormatter) hour24Nonzero() string {
	return strconv.Itoa(obj.nepaliTime.Hour())
}

// %I
func (obj *NepaliFormatter) hour12() string {
	hour := obj.nepaliTime.Hour()

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
func (obj *NepaliFormatter) hour12Nonzero() string {
	hour := obj.nepaliTime.Hour()

	if hour > 12 {
		hour -= 12
	}
	if hour == 0 {
		hour = 12
	}

	return strconv.Itoa(hour)
}

// %p
func (obj *NepaliFormatter) ampm() string {
	ampm := "AM"

	if obj.nepaliTime.Hour() > 12 {
		ampm = "PM"
	}

	return ampm
}

// %M
func (obj *NepaliFormatter) minute() string {
	minute := strconv.Itoa(obj.nepaliTime.Minute())

	if len(minute) < 2 {
		minute = "0" + minute
	}

	return minute
}

// %-M
func (obj *NepaliFormatter) minuteNonzero() string {
	return strconv.Itoa(obj.nepaliTime.Minute())
}

// %s
func (obj *NepaliFormatter) second() string {
	second := strconv.Itoa(obj.nepaliTime.Second())

	if len(second) < 2 {
		second = "0" + second
	}

	return second
}

// %-s
func (obj *NepaliFormatter) secondNonzero() string {
	return strconv.Itoa(obj.nepaliTime.Second())
}

// %f
func (obj *NepaliFormatter) nanosecond_() string {
	nsec := strconv.Itoa(obj.nepaliTime.Nanosecond())

	if len(nsec) < 6 {
		nsec = strings.Repeat("0", 6-len(nsec)) + nsec
	}

	return nsec
}

// %-f
func (obj *NepaliFormatter) nanosecondNonZero() string {
	return strconv.Itoa(obj.nepaliTime.Nanosecond())
}
