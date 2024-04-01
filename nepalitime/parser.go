package nepalitime

import (
	"errors"
	"strconv"
	"strings"
)

var nepaliTimeReCache *nepaliTimeRegex

// equivalent to time.Parse()
func Parse(datetimeStr string, format string) (*NepaliTime, error) {
	nepalitime, err := validate(datetimeStr, format)

	if err != nil {
		return nil, err
	}

	return nepalitime, nil
}

func getNepaliTimeReObject() *nepaliTimeRegex {
	if nepaliTimeReCache == nil {
		nepaliTimeReCache = newNepaliTimeRegex()
	}

	return nepaliTimeReCache
}

// validates datetimeStr with the format
func validate(datetimeStr string, format string) (*NepaliTime, error) {
	// validate if parse result is not empty
	parsedResult, err := extract(datetimeStr, format)
	if err != nil {
		return nil, err
	} else {
		_, ok1 := parsedResult["Y"]
		_, ok2 := parsedResult["y"]

		if !ok1 && !ok2 {
			return nil, errors.New("unable to parse year")
		}
	}

	// validate the transformation
	transformedData, err := transform(parsedResult)
	if err != nil {
		return nil, err
	}

	nepaliDate, err := Date(
		transformedData["year"],
		transformedData["month"],
		transformedData["day"],
		transformedData["hour"],
		transformedData["minute"],
		transformedData["second"],
		transformedData["nanosecond"],
	)
	if err != nil {
		return nil, err
	}

	return nepaliDate, nil
}

// extracts year, month, day, hour, minute, etc from the given format
// eg.
// USAGE: extract("2078-01-12", "%Y-%m-%d")
// INPUT:
// datetime_str="2078-01-12"
// format="%Y-%m-%d"
// OUTPUT:
//
//	{
//		"Y": 2078,
//		"m": 1,
//		"d": 12,
//	}
func extract(datetimeStr string, format string) (map[string]string, error) {
	reCompiledFormat, err := getNepaliTimeReObject().compile(format)

	if err != nil {
		return nil, err
	}

	match := reCompiledFormat.FindStringSubmatch(datetimeStr)

	if len(match) < 1 {
		return nil, errors.New("datetime string did not match with given format")
	}

	result := make(map[string]string)

	for index, name := range reCompiledFormat.SubexpNames() {
		if index != 0 && name != "" {
			result[name] = match[index]
		}
	}

	return result, nil
}

// transforms different format data to uniform data
// eg.
// INPUT:
//
//	data = {
//	    "Y": 2078,
//	    "b": "Mangsir",
//	    "d": 12,
//	    ...
//	}
//
// OUTPUT:
//
//	{
//	    "year": 2078,
//	    "month": 8,
//	    "day": 12,
//	    ...
//	}
func transform(data map[string]string) (map[string]int, error) {
	var (
		year                           int
		month, day                     int = 1, 1
		hour, minute, second, fraction int = 0, 0, 0, 0
	)

	for key, val := range data {
		if key == "y" {
			intVal, err := strconv.Atoi(val)
			if err != nil {
				return nil, errors.New("invalid value in %y")
			}

			year = intVal
			year += 2000
		} else if key == "Y" {
			intVal, err := strconv.Atoi(val)
			if err != nil {
				return nil, errors.New("invalid value in %Y")
			}

			year = intVal
		} else if key == "m" {
			intVal, err := strconv.Atoi(val)
			if err != nil {
				return nil, errors.New("invalid value in %Y")
			}

			month = intVal
		} else if key == "d" {
			intVal, err := strconv.Atoi(val)
			if err != nil {
				return nil, errors.New("invalid value in %d")
			}

			day = intVal
		} else if key == "H" {
			intVal, err := strconv.Atoi(val)
			if err != nil {
				return nil, errors.New("invalid value in %H")
			}

			hour = intVal
		} else if key == "I" {
			intVal, err := strconv.Atoi(val)
			if err != nil {
				return nil, errors.New("invalid value in %I")
			}

			hour = intVal

			ampm, ok := data["p"]
			if !ok {
				ampm = ""
			}

			ampm = strings.ToLower(ampm)
			// if there is no AM/PM indicator, we'll treat it as
			if ampm == "" || ampm == "am" {
				if hour == 12 {
					hour = 0
				}
			} else if ampm == "pm" {
				if hour != 12 {
					hour += 12
				}
			}
		} else if key == "M" {
			intVal, err := strconv.Atoi(val)
			if err != nil {
				return nil, errors.New("invalid value in %M")
			}

			minute = intVal
		} else if key == "S" {
			intVal, err := strconv.Atoi(val)
			if err != nil {
				return nil, errors.New("invalid value in %S")
			}

			second = intVal
		} else if key == "f" {
			var err error

			s := val
			s += strings.Repeat("0", 6-len(val))

			fraction, err = strconv.Atoi(s)

			if err != nil {
				return nil, errors.New("invalid value in %f")
			}
		}
	}

	return map[string]int{
		"year":       year,
		"month":      month,
		"day":        day,
		"hour":       hour,
		"minute":     minute,
		"second":     second,
		"nanosecond": fraction,
	}, nil
}
