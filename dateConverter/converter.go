// This package contains the date converter functions.
//
// USAGE:
// y, m, d := EnglishToNepali(2023, 1, 15)
// y, m, d := NepaliToEnglish(2079, 10, 1)
package dateConverter

import (
	"time"
	"errors"
	"math"
)

// Reference date for conversion is 2000/01/01 BS and 1943/4/14 AD
var npInitialYear int16 = 2000
var referenceEnDate = [3]int16{1943, 4, 14}
// A struct for Date
type Date struct {
	Year	int
	Month	int
	Day	int
	Calendar string
}

type NepaliMonthData struct {
	monthData [12]int8
	yearDays  int16
}

var enMonths = [12]int8{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
var enLeapMonths = [12]int8{31, 29, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

var npMonthData = [100]NepaliMonthData{
	{[12]int8{30, 32, 31, 32, 31, 30, 30, 30, 29, 30, 29, 31}, 365}, // 2000 BS - 1943/1944 AD
	{[12]int8{31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30}, 365}, // 2001 BS
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
	{[12]int8{31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 30}, 365}, // 2080 BS
	{[12]int8{31, 32, 31, 32, 31, 30, 30, 30, 29, 30, 29, 31}, 366},
	{[12]int8{30, 32, 31, 32, 31, 30, 30, 30, 29, 30, 30, 30}, 365},
	{[12]int8{31, 31, 32, 31, 31, 30, 30, 30, 29, 30, 30, 30}, 365},
	{[12]int8{31, 31, 32, 31, 31, 30, 30, 30, 29, 30, 30, 30}, 365},
	{[12]int8{31, 32, 31, 32, 30, 31, 30, 30, 29, 30, 30, 30}, 366},
	{[12]int8{30, 32, 31, 32, 31, 30, 30, 30, 29, 30, 30, 30}, 365},
	{[12]int8{31, 31, 32, 31, 31, 31, 30, 30, 29, 30, 30, 30}, 366},
	{[12]int8{30, 31, 32, 32, 30, 31, 30, 30, 29, 30, 30, 30}, 365},
	{[12]int8{30, 32, 31, 32, 31, 30, 30, 30, 29, 30, 30, 30}, 365},
	{[12]int8{30, 32, 31, 32, 31, 30, 30, 30, 29, 30, 30, 30}, 365},
	{[12]int8{31, 31, 32, 31, 31, 31, 30, 30, 29, 30, 30, 30}, 366},
	{[12]int8{30, 31, 32, 32, 31, 30, 30, 30, 29, 30, 30, 30}, 365},
	{[12]int8{30, 32, 31, 32, 31, 30, 30, 30, 29, 30, 30, 30}, 365},
	{[12]int8{31, 31, 32, 31, 31, 30, 30, 30, 29, 30, 30, 30}, 365},
	{[12]int8{31, 31, 32, 31, 31, 31, 30, 29, 30, 30, 30, 30}, 366},
	{[12]int8{30, 31, 32, 32, 31, 30, 30, 29, 30, 29, 30, 30}, 364},
	{[12]int8{31, 32, 31, 32, 31, 30, 30, 30, 29, 30, 30, 30}, 366},
	{[12]int8{31, 31, 32, 31, 31, 31, 29, 30, 29, 30, 29, 31}, 365},
	{[12]int8{31, 31, 32, 31, 31, 31, 30, 29, 29, 30, 30, 30}, 365}, // 2099 BS - 2042 AD
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
func checkEnglishDate(date Date) bool {
	if date.Year < enMinYear() || date.Year > enMaxYear() {
		return false
	}
	if date.Month < 1 || date.Month > 12 {
		return false
	}
	if date.Day < 1 || date.Day > 31 {
		return false
	}
	return true
}

// counts and returns total days from the date 0000-01-01
func getTotalDaysFromEnglishDate(date Date) int {
	var totalDays int = date.Year*365 + date.Day
	for i := 0; i < date.Month-1; i++ {
		totalDays = totalDays + int(enMonths[i])
	}

	// adding leap days (ie. leap year count)
	if date.Month <= 2 { // checking February month (where leap exists)
		date.Year -= 1
	}
	totalDays += int(date.Year/4) - int(date.Year/100) + int(date.Year/400)

	return totalDays
}

// NEPALI DATE CONVERSION

// checks if nepali date is in range
func checkNepaliDate(date Date) bool {
	if date.Year < npMinYear() || date.Year > npMaxYear() {
		return false
	}
	if date.Month < 1 || date.Month > 12 {
		return false
	}

	if date.Day < 1 || date.Day > int(npMonthData[date.Year-int(npInitialYear)].monthData[date.Month-1]) {
		return false
	}
	return true
}

// counts and returns total days from the nepali initial date
func getTotalDaysFromNepaliDate(date Date) int {
	var totalDays int = date.Day - 1

	// adding days of months of initial year
	var yearIndex int = date.Year - int(npInitialYear)
	for i := 0; i < date.Month-1; i++ {
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
func EnglishToNepali(year int, month int, day int) (*Date, error) {
	// constructing the input date
	enDate := Date{
		Year: year,
		Month: month,
		Day: day,
	}
	// VALIDATION
	// checking if date is in range
	if !checkEnglishDate(enDate) {
		return nil, errors.New("date is out of range")
	}

	// REFERENCE
	// constructing the output date
	npDate := Date{
		Year: int(npInitialYear),
		Month: 1,
		Day: 1,
		Calendar: "Nepali",
	}
	// DIFFERENCE
	// calculating days count from the reference date
	var difference int = int(
		math.Abs(float64(
			getTotalDaysFromEnglishDate(enDate) - getTotalDaysFromEnglishDate(Date{int(referenceEnDate[0]), int(referenceEnDate[1]), int(referenceEnDate[2]), ""}))))

	// YEAR
	// Incrementing year until the difference remains less than 365
	var yearDataIndex int = 0
	for difference >= int(npMonthData[yearDataIndex].yearDays) {
		difference -= int(npMonthData[yearDataIndex].yearDays)
		npDate.Year += 1
		yearDataIndex += 1
	}

	// MONTH
	// Incrementing month until the difference remains less than next nepali month days (mostly 31)
	var i int = 0
	for difference >= int(npMonthData[yearDataIndex].monthData[i]) {
		difference -= int(npMonthData[yearDataIndex].monthData[i])
		npDate.Month += 1
		i += 1
	}

	// DAY
	// Remaining difference is the day
	npDate.Day += difference

	return &npDate, nil
}

// Converts nepali date to english.
// Accepts the input parameters year, month, day.
// Returns dates in array and error.
func NepaliToEnglish(year int, month int, day int) (*Date, error) {
	// constructing the input date
	npDate := Date{
		Year: year,
		Month: month,
		Day: day,
	}
	// VALIDATION
	// checking if date is in range
	if !checkNepaliDate(npDate) {
		return nil, errors.New("date is out of range")
	}

	// REFERENCE
	// For absolute reference, moving date to Jan 1
	// Eg. ref: 1943/4/14 => 1943/01/01
	// constructing the output date
	enDate := Date{
		Year: int(referenceEnDate[0]),
		Month: 1,
		Day: 1,
		Calendar: "English",
	}
	// calculating difference from the adjusted reference (eg. 1943/4/14 - 1943/01/01)
	referenceDiff := getDiffFromEnAbsoluteReference()

	// DIFFERENCE
	// calculating days count from the reference date
	var difference int = getTotalDaysFromNepaliDate(npDate) + referenceDiff

	// YEAR
	// Incrementing year until the difference remains less than 365 (or 365)
	for (difference >= 366 && isLeapYear(enDate.Year)) || (difference >= 365 && !(isLeapYear(enDate.Year))) {
		if isLeapYear(enDate.Year) {
			difference -= 366
		} else {
			difference -= 365
		}
		enDate.Year += 1
	}

	// MONTH
	// Incrementing month until the difference remains less than next english month (mostly 31)
	monthDays := getEnMonths(enDate.Year)

	var i int = 0
	for difference >= int(monthDays[i]) {
		difference -= int(monthDays[i])
		enDate.Month += 1
		i += 1
	}

	// DAY
	// Remaining difference is the day
	enDate.Day += difference

	return &enDate, nil
}

// Weekday returns the weekday of the date.
func (d *Date) Weekday() string {
	var t time.Time
	if d.Calendar == "Nepali" {
		// Convert the Nepali date to English
		convertedEnDate, _ := NepaliToEnglish(d.Year, d.Month, d.Day)
		// Construct a time.Time object from the converted English date
		t = time.Date(convertedEnDate.Year, time.Month(convertedEnDate.Month), convertedEnDate.Day, 0, 0, 0, 0, time.UTC)
	} else {
		// Construct a time.Time object from the English date
		t = time.Date(d.Year, time.Month(d.Month), d.Day, 0, 0, 0, 0, time.UTC)
	}
	// Use the Weekday() function from the time package to get the weekday
	return t.Weekday().String()
}
