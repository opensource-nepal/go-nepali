// This package contains the date converter functions.
//
// USAGE:
// y, m, d := EnglishToNepali(2023, 1, 15)
// y, m, d := NepaliToEnglish(2079, 10, 1)
package dateConverter

import (
	"errors"
	"math"
)

// Reference date for conversion is 2000/01/01 BS and 1943/4/14 AD
var npInitialYear int16 = 2000
var referenceEnDate = [3]int16{1943, 4, 14}

type NepaliMonthData struct {
	monthData [12]int8
	yearDays  int16
}

var enMonths = [12]int8{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
var enLeapMonths = [12]int8{31, 29, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

var npMonthData = [100]NepaliMonthData{
	{[12]int8{30, 32, 31, 32, 31, 30, 30, 30, 29, 30, 29, 31}, 365}, // 2000 BS - 1943/1944 AD
	{[12]int8{31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30}, 365},
	{[12]int8{31, 31, 32, 32, 31, 30, 30, 29, 30, 29, 30, 30}, 365},
	{[12]int8{31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 31}, 366},
	{[12]int8{30, 32, 31, 32, 31, 30, 30, 30, 29, 30, 29, 31}, 365},
	{[12]int8{31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30}, 365},
	{[12]int8{31, 31, 32, 32, 31, 30, 30, 29, 30, 29, 30, 30}, 365},
	{[12]int8{31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 31}, 366},
	{[12]int8{31, 31, 31, 32, 31, 31, 29, 30, 30, 29, 29, 31}, 365},
	{[12]int8{31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30}, 365},
	{[12]int8{31, 31, 32, 32, 31, 30, 30, 29, 30, 29, 30, 30}, 365},
	{[12]int8{31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 31}, 366},
	{[12]int8{31, 31, 31, 32, 31, 31, 29, 30, 30, 29, 30, 30}, 365},
	{[12]int8{31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30}, 365},
	{[12]int8{31, 31, 32, 32, 31, 30, 30, 29, 30, 29, 30, 30}, 365},
	{[12]int8{31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 31}, 366},
	{[12]int8{31, 31, 31, 32, 31, 31, 29, 30, 30, 29, 30, 30}, 365},
	{[12]int8{31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30}, 365},
	{[12]int8{31, 32, 31, 32, 31, 30, 30, 29, 30, 29, 30, 30}, 365},
	{[12]int8{31, 32, 31, 32, 31, 30, 30, 30, 29, 30, 29, 31}, 366},
	{[12]int8{31, 31, 31, 32, 31, 31, 30, 29, 30, 29, 30, 30}, 365},
	{[12]int8{31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30}, 365},
	{[12]int8{31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 30}, 365},
	{[12]int8{31, 32, 31, 32, 31, 30, 30, 30, 29, 30, 29, 31}, 366},
	{[12]int8{31, 31, 31, 32, 31, 31, 30, 29, 30, 29, 30, 30}, 365},
	{[12]int8{31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30}, 365},
	{[12]int8{31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 31}, 366},
	{[12]int8{30, 32, 31, 32, 31, 30, 30, 30, 29, 30, 29, 31}, 365},
	{[12]int8{31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30}, 365},
	{[12]int8{31, 31, 32, 31, 32, 30, 30, 29, 30, 29, 30, 30}, 365},
	{[12]int8{31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 31}, 366},
	{[12]int8{30, 32, 31, 32, 31, 30, 30, 30, 29, 30, 29, 31}, 365},
	{[12]int8{31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30}, 365},
	{[12]int8{31, 31, 32, 32, 31, 30, 30, 29, 30, 29, 30, 30}, 365},
	{[12]int8{31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 31}, 366},
	{[12]int8{30, 32, 31, 32, 31, 31, 29, 30, 30, 29, 29, 31}, 365},
	{[12]int8{31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30}, 365},
	{[12]int8{31, 31, 32, 32, 31, 30, 30, 29, 30, 29, 30, 30}, 365},
	{[12]int8{31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 31}, 366},
	{[12]int8{31, 31, 31, 32, 31, 31, 29, 30, 30, 29, 30, 30}, 365},
	{[12]int8{31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30}, 365},
	{[12]int8{31, 31, 32, 32, 31, 30, 30, 29, 30, 29, 30, 30}, 365},
	{[12]int8{31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 31}, 366},
	{[12]int8{31, 31, 31, 32, 31, 31, 29, 30, 30, 29, 30, 30}, 365},
	{[12]int8{31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30}, 365},
	{[12]int8{31, 32, 31, 32, 31, 30, 30, 29, 30, 29, 30, 30}, 365},
	{[12]int8{31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 31}, 366},
	{[12]int8{31, 31, 31, 32, 31, 31, 30, 29, 30, 29, 30, 30}, 365},
	{[12]int8{31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30}, 365},
	{[12]int8{31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 30}, 365},
	{[12]int8{31, 32, 31, 32, 31, 30, 30, 30, 29, 30, 29, 31}, 366},
	{[12]int8{31, 31, 31, 32, 31, 31, 30, 29, 30, 29, 30, 30}, 365},
	{[12]int8{31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30}, 365},
	{[12]int8{31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 30}, 365},
	{[12]int8{31, 32, 31, 32, 31, 30, 30, 30, 29, 30, 29, 31}, 366},
	{[12]int8{31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30}, 365},
	{[12]int8{31, 31, 32, 31, 32, 30, 30, 29, 30, 29, 30, 30}, 365},
	{[12]int8{31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 31}, 366},
	{[12]int8{30, 32, 31, 32, 31, 30, 30, 30, 29, 30, 29, 31}, 365},
	{[12]int8{31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30}, 365},
	{[12]int8{31, 31, 32, 32, 31, 30, 30, 29, 30, 29, 30, 30}, 365},
	{[12]int8{31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 31}, 366},
	{[12]int8{30, 32, 31, 32, 31, 31, 29, 30, 29, 30, 29, 31}, 365},
	{[12]int8{31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30}, 365},
	{[12]int8{31, 31, 32, 32, 31, 30, 30, 29, 30, 29, 30, 30}, 365},
	{[12]int8{31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 31}, 366},
	{[12]int8{31, 31, 31, 32, 31, 31, 29, 30, 30, 29, 29, 31}, 365},
	{[12]int8{31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30}, 365},
	{[12]int8{31, 31, 32, 32, 31, 30, 30, 29, 30, 29, 30, 30}, 365},
	{[12]int8{31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 31}, 366},
	{[12]int8{31, 31, 31, 32, 31, 31, 29, 30, 30, 29, 30, 30}, 365},
	{[12]int8{31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30}, 365},
	{[12]int8{31, 32, 31, 32, 31, 30, 30, 29, 30, 29, 30, 30}, 365},
	{[12]int8{31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 31}, 366},
	{[12]int8{31, 31, 31, 32, 31, 31, 30, 29, 30, 29, 30, 30}, 365},
	{[12]int8{31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30}, 365},
	{[12]int8{31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 30}, 365},
	{[12]int8{31, 32, 31, 32, 31, 30, 30, 30, 29, 30, 29, 31}, 366},
	{[12]int8{31, 31, 31, 32, 31, 31, 30, 29, 30, 29, 30, 30}, 365},
	{[12]int8{31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30}, 365},
	{[12]int8{31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 30}, 365},
	{[12]int8{31, 32, 31, 32, 31, 30, 30, 30, 29, 30, 29, 31}, 366}, // 2081 BS
	{[12]int8{31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30}, 365},
	{[12]int8{31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30}, 365},
	{[12]int8{31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 31}, 366},
	{[12]int8{30, 32, 31, 32, 31, 30, 30, 30, 29, 30, 29, 31}, 365},
	{[12]int8{31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30}, 365},
	{[12]int8{31, 31, 32, 32, 31, 30, 30, 29, 30, 29, 30, 30}, 365},
	{[12]int8{31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 31}, 366},
	{[12]int8{30, 32, 31, 32, 31, 30, 30, 30, 29, 30, 29, 31}, 365},
	{[12]int8{31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30}, 365},
	{[12]int8{31, 31, 32, 32, 31, 30, 30, 29, 30, 29, 30, 30}, 365},
	{[12]int8{31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 31}, 366},
	{[12]int8{31, 31, 31, 32, 31, 31, 29, 30, 30, 29, 29, 31}, 365},
	{[12]int8{31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30}, 365},
	{[12]int8{31, 31, 32, 32, 31, 30, 30, 29, 30, 29, 30, 30}, 365},
	{[12]int8{31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 31}, 366},
	{[12]int8{31, 31, 31, 32, 31, 31, 29, 30, 30, 29, 30, 30}, 365},
	{[12]int8{31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30}, 365},
	{[12]int8{31, 31, 32, 32, 31, 30, 30, 29, 30, 29, 30, 30}, 365}, // 2099 BS - 2042/2043 AD
}

func enMinYear() int {
	return int(referenceEnDate[0]) + 1
}

func enMaxYear() int {
	return int(referenceEnDate[0]) + len(npMonthData) - 1
}

func npMinYear() int {
	return int(npInitialYear)
}

func npMaxYear() int {
	return int(npInitialYear) + len(npMonthData) - 1
}

/* Checks if the english year is leap year or not */
func isLeapYear(year int) bool {
	if year%4 == 0 {
		if year%100 == 0 {
			return (year%400 == 0)
		}
		return true
	}
	return false
}

func getEnMonths(year int) *[12]int8 {
	if isLeapYear(year) {
		return &enLeapMonths
	}
	return &enMonths
}

/*
Returns diff from the reference with its absolute reference.
Used in function `NepaliToEnglish`.

Eg. ref: 1943/4/14 - 1943/01/01
*/
func getDiffFromEnAbsoluteReference() int {
	var diff int = 0

	// adding sum of month of year till the reference month
	months := getEnMonths(int(referenceEnDate[0]))
	for i := 0; i < int(referenceEnDate[1])-1; i++ {
		diff += int(months[i])
	}

	return diff + int(referenceEnDate[2]) - 1 // added day too
}

// ENGLISH DATE CONVERSION

// checks if english date in within range 1944 - 2042
func checkEnglishDate(year int, month int, day int) bool {
	if year < enMinYear() || year > enMaxYear() {
		return false
	}
	if month < 1 || month > 12 {
		return false
	}
	if day < 1 || day > 31 {
		return false
	}
	return true
}

// counts and returns total days from the date 0000-01-01
func getTotalDaysFromEnglishDate(year int, month int, day int) int {
	var totalDays int = year*365 + day
	for i := 0; i < month-1; i++ {
		totalDays = totalDays + int(enMonths[i])
	}

	// adding leap days (ie. leap year count)
	if month <= 2 { // checking February month (where leap exists)
		year -= 1
	}
	totalDays += int(year/4) - int(year/100) + int(year/400)

	return totalDays
}

// NEPALI DATE CONVERSION

// checks if nepali date is in range
func checkNepaliDate(year int, month int, day int) bool {
	if year < npMinYear() || year > npMaxYear() {
		return false
	}
	if month < 1 || month > 12 {
		return false
	}

	if day < 1 || day > int(npMonthData[year-int(npInitialYear)].monthData[month-1]) {
		return false
	}
	return true
}

// counts and returns total days from the nepali initial date
func getTotalDaysFromNepaliDate(year int, month int, day int) int {
	var totalDays int = day - 1

	// adding days of months of initial year
	var yearIndex int = year - int(npInitialYear)
	for i := 0; i < month-1; i++ {
		totalDays = totalDays + int(npMonthData[yearIndex].monthData[i])
	}

	// adding days of year
	for i := 0; i < yearIndex; i++ {
		totalDays = totalDays + int(npMonthData[i].yearDays)
	}
	return totalDays
}

// Public methods

// Converts english date to nepali.
// Accepts the input parameters year, month, day.
// Returns dates in array and error.
func EnglishToNepali(year int, month int, day int) (*[3]int, error) {
	// VALIDATION
	// checking if date is in range
	if !checkEnglishDate(year, month, day) {
		return nil, errors.New("date is out of range")
	}

	// REFERENCE
	npYear, npMonth, npDay := int(npInitialYear), 1, 1

	// DIFFERENCE
	// calculating days count from the reference date
	var difference int = int(
		math.Abs(float64(
			getTotalDaysFromEnglishDate(
				year, month, day,
			) - getTotalDaysFromEnglishDate(
				int(referenceEnDate[0]), int(referenceEnDate[1]), int(referenceEnDate[2]),
			),
		)),
	)

	// YEAR
	// Incrementing year until the difference remains less than 365
	var yearDataIndex int = 0
	for difference >= int(npMonthData[yearDataIndex].yearDays) {
		difference -= int(npMonthData[yearDataIndex].yearDays)
		npYear += 1
		yearDataIndex += 1
	}

	// MONTH
	// Incrementing month until the difference remains less than next nepali month days (mostly 31)
	var i int = 0
	for difference >= int(npMonthData[yearDataIndex].monthData[i]) {
		difference -= int(npMonthData[yearDataIndex].monthData[i])
		npMonth += 1
		i += 1
	}

	// DAY
	// Remaining difference is the day
	npDay += difference

	return &[3]int{npYear, npMonth, npDay}, nil
}

// Converts nepali date to english.
// Accepts the input parameters year, month, day.
// Returns dates in array and error.
func NepaliToEnglish(year int, month int, day int) (*[3]int, error) {
	// VALIDATION
	// checking if date is in range
	if !checkNepaliDate(year, month, day) {
		return nil, errors.New("date is out of range")
	}

	// REFERENCE
	// For absolute reference, moving date to Jan 1
	// Eg. ref: 1943/4/14 => 1943/01/01
	enYear, enMonth, enDay := int(referenceEnDate[0]), 1, 1
	// calculating difference from the adjusted reference (eg. 1943/4/14 - 1943/01/01)
	referenceDiff := getDiffFromEnAbsoluteReference()

	// DIFFERENCE
	// calculating days count from the reference date
	var difference int = getTotalDaysFromNepaliDate(year, month, day) + referenceDiff

	// YEAR
	// Incrementing year until the difference remains less than 365 (or 365)
	for (difference >= 366 && isLeapYear(enYear)) || (difference >= 365 && !(isLeapYear(enYear))) {
		if isLeapYear(enYear) {
			difference -= 366
		} else {
			difference -= 365
		}
		enYear += 1
	}

	// MONTH
	// Incrementing month until the difference remains less than next english month (mostly 31)
	monthDays := getEnMonths(enYear)

	var i int = 0
	for difference >= int(monthDays[i]) {
		difference -= int(monthDays[i])
		enMonth += 1
		i += 1
	}

	// DAY
	// Remaining difference is the day
	enDay += difference

	return &[3]int{enYear, enMonth, enDay}, nil
}
