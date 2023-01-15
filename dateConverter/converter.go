// This package contains the date converter functions.
// Reference date for conversion is 2000/01/01 BS and 1943/4/14 AD
//
// USAGE:
// y, m, d := EnglishToNepali(2023, 1, 15)
// y, m, d := NepaliToEnglish(2079, 10, 1)
package dateConverter

import (
	"math"
)

type NepaliMonthData struct {
	monthData [12]int
	yearDays  int
}

func getEnMonths(isLeapYear bool) [12]int {
	if isLeapYear {
		return [12]int{31, 29, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	}
	return [12]int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
}

func getNpMonthsData() []NepaliMonthData {
	return []NepaliMonthData{
		{[12]int{30, 32, 31, 32, 31, 30, 30, 30, 29, 30, 29, 31}, 365}, // 2000 BS - 1944 AD
		{[12]int{31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30}, 365}, // 2001 BS
		{[12]int{31, 31, 32, 32, 31, 30, 30, 29, 30, 29, 30, 30}, 365},
		{[12]int{31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 31}, 366},
		{[12]int{30, 32, 31, 32, 31, 30, 30, 30, 29, 30, 29, 31}, 365},
		{[12]int{31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30}, 365},
		{[12]int{31, 31, 32, 32, 31, 30, 30, 29, 30, 29, 30, 30}, 365},
		{[12]int{31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 31}, 366},
		{[12]int{31, 31, 31, 32, 31, 31, 29, 30, 30, 29, 29, 31}, 365},
		{[12]int{31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30}, 365},
		{[12]int{31, 31, 32, 32, 31, 30, 30, 29, 30, 29, 30, 30}, 365},
		{[12]int{31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 31}, 366},
		{[12]int{31, 31, 31, 32, 31, 31, 29, 30, 30, 29, 30, 30}, 365},
		{[12]int{31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30}, 365},
		{[12]int{31, 31, 32, 32, 31, 30, 30, 29, 30, 29, 30, 30}, 365},
		{[12]int{31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 31}, 366},
		{[12]int{31, 31, 31, 32, 31, 31, 29, 30, 30, 29, 30, 30}, 365},
		{[12]int{31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30}, 365},
		{[12]int{31, 32, 31, 32, 31, 30, 30, 29, 30, 29, 30, 30}, 365},
		{[12]int{31, 32, 31, 32, 31, 30, 30, 30, 29, 30, 29, 31}, 366},
		{[12]int{31, 31, 31, 32, 31, 31, 30, 29, 30, 29, 30, 30}, 365},
		{[12]int{31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30}, 365},
		{[12]int{31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 30}, 365},
		{[12]int{31, 32, 31, 32, 31, 30, 30, 30, 29, 30, 29, 31}, 366},
		{[12]int{31, 31, 31, 32, 31, 31, 30, 29, 30, 29, 30, 30}, 365},
		{[12]int{31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30}, 365},
		{[12]int{31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 31}, 366},
		{[12]int{30, 32, 31, 32, 31, 30, 30, 30, 29, 30, 29, 31}, 365},
		{[12]int{31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30}, 365},
		{[12]int{31, 31, 32, 31, 32, 30, 30, 29, 30, 29, 30, 30}, 365},
		{[12]int{31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 31}, 366},
		{[12]int{30, 32, 31, 32, 31, 30, 30, 30, 29, 30, 29, 31}, 365},
		{[12]int{31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30}, 365},
		{[12]int{31, 31, 32, 32, 31, 30, 30, 29, 30, 29, 30, 30}, 365},
		{[12]int{31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 31}, 366},
		{[12]int{30, 32, 31, 32, 31, 31, 29, 30, 30, 29, 29, 31}, 365},
		{[12]int{31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30}, 365},
		{[12]int{31, 31, 32, 32, 31, 30, 30, 29, 30, 29, 30, 30}, 365},
		{[12]int{31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 31}, 366},
		{[12]int{31, 31, 31, 32, 31, 31, 29, 30, 30, 29, 30, 30}, 365},
		{[12]int{31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30}, 365},
		{[12]int{31, 31, 32, 32, 31, 30, 30, 29, 30, 29, 30, 30}, 365},
		{[12]int{31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 31}, 366},
		{[12]int{31, 31, 31, 32, 31, 31, 29, 30, 30, 29, 30, 30}, 365},
		{[12]int{31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30}, 365},
		{[12]int{31, 32, 31, 32, 31, 30, 30, 29, 30, 29, 30, 30}, 365},
		{[12]int{31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 31}, 366},
		{[12]int{31, 31, 31, 32, 31, 31, 30, 29, 30, 29, 30, 30}, 365},
		{[12]int{31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30}, 365},
		{[12]int{31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 30}, 365},
		{[12]int{31, 32, 31, 32, 31, 30, 30, 30, 29, 30, 29, 31}, 366},
		{[12]int{31, 31, 31, 32, 31, 31, 30, 29, 30, 29, 30, 30}, 365},
		{[12]int{31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30}, 365},
		{[12]int{31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 30}, 365},
		{[12]int{31, 32, 31, 32, 31, 30, 30, 30, 29, 30, 29, 31}, 366},
		{[12]int{31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30}, 365},
		{[12]int{31, 31, 32, 31, 32, 30, 30, 29, 30, 29, 30, 30}, 365},
		{[12]int{31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 31}, 366},
		{[12]int{30, 32, 31, 32, 31, 30, 30, 30, 29, 30, 29, 31}, 365},
		{[12]int{31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30}, 365},
		{[12]int{31, 31, 32, 32, 31, 30, 30, 29, 30, 29, 30, 30}, 365},
		{[12]int{31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 31}, 366},
		{[12]int{30, 32, 31, 32, 31, 31, 29, 30, 29, 30, 29, 31}, 365},
		{[12]int{31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30}, 365},
		{[12]int{31, 31, 32, 32, 31, 30, 30, 29, 30, 29, 30, 30}, 365},
		{[12]int{31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 31}, 366},
		{[12]int{31, 31, 31, 32, 31, 31, 29, 30, 30, 29, 29, 31}, 365},
		{[12]int{31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30}, 365},
		{[12]int{31, 31, 32, 32, 31, 30, 30, 29, 30, 29, 30, 30}, 365},
		{[12]int{31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 31}, 366},
		{[12]int{31, 31, 31, 32, 31, 31, 29, 30, 30, 29, 30, 30}, 365},
		{[12]int{31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30}, 365},
		{[12]int{31, 32, 31, 32, 31, 30, 30, 29, 30, 29, 30, 30}, 365},
		{[12]int{31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 31}, 366},
		{[12]int{31, 31, 31, 32, 31, 31, 30, 29, 30, 29, 30, 30}, 365},
		{[12]int{31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30}, 365},
		{[12]int{31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 30}, 365},
		{[12]int{31, 32, 31, 32, 31, 30, 30, 30, 29, 30, 29, 31}, 366},
		{[12]int{31, 31, 31, 32, 31, 31, 30, 29, 30, 29, 30, 30}, 365},
		{[12]int{31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30}, 365},
		{[12]int{31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 30}, 365},
		{[12]int{31, 31, 32, 32, 31, 30, 30, 30, 29, 30, 30, 30}, 366},
		{[12]int{30, 32, 31, 32, 31, 30, 30, 30, 29, 30, 30, 30}, 365},
		{[12]int{31, 31, 32, 31, 31, 30, 30, 30, 29, 30, 30, 30}, 365},
		{[12]int{31, 31, 32, 31, 31, 30, 30, 30, 29, 30, 30, 30}, 365},
		{[12]int{31, 32, 31, 32, 30, 31, 30, 30, 29, 30, 30, 30}, 366},
		{[12]int{30, 32, 31, 32, 31, 30, 30, 30, 29, 30, 30, 30}, 365},
		{[12]int{31, 31, 32, 31, 31, 31, 30, 30, 29, 30, 30, 30}, 366},
		{[12]int{30, 31, 32, 32, 30, 31, 30, 30, 29, 30, 30, 30}, 365},
		{[12]int{30, 32, 31, 32, 31, 30, 30, 30, 29, 30, 30, 30}, 365},
		{[12]int{30, 32, 31, 32, 31, 30, 30, 30, 29, 30, 30, 30}, 365},
		{[12]int{31, 31, 32, 31, 31, 31, 30, 30, 29, 30, 30, 30}, 366},
		{[12]int{30, 31, 32, 32, 31, 30, 30, 30, 29, 30, 30, 30}, 365},
		{[12]int{30, 32, 31, 32, 31, 30, 30, 30, 29, 30, 30, 30}, 365},
		{[12]int{31, 31, 32, 31, 31, 30, 30, 30, 29, 30, 30, 30}, 365},
		{[12]int{31, 31, 32, 31, 31, 31, 30, 29, 30, 30, 30, 30}, 366},
		{[12]int{30, 31, 32, 32, 31, 30, 30, 29, 30, 29, 30, 30}, 364},
		{[12]int{31, 32, 31, 32, 31, 30, 30, 30, 29, 30, 30, 30}, 366},
		{[12]int{31, 31, 32, 31, 31, 31, 29, 30, 29, 30, 29, 31}, 365},
		{[12]int{31, 31, 32, 31, 31, 31, 30, 29, 29, 30, 30, 30}, 365}, // 2099 BS - 2042 AD
	}
}

// ENGLISH DATE CONVERSION

// checks if english date in within range 1944 - 2042
func checkEnglishDate(year int, month int, day int) bool {
	if year < 1944 || year > 2042 {
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
	enMonths := getEnMonths(false)
	for i := 0; i < month-1; i++ {
		totalDays = totalDays + enMonths[i]
	}

	// adding leap days (ie. leap year count)
	if month <= 2 { // checking February month (where leap exists)
		year -= 1
	}
	totalDays += int(year/4) - int(year/100) + int(year/400)

	return totalDays
}

// NEPALI DATE CONVERSION

// checks if nepali date is in range 2000-2098
func checkNepaliDate(year int, month int, day int) bool {
	if year < 2000 || year > 2098 {
		return false
	}
	if month < 1 || month > 12 {
		return false
	}

	npMonthData := getNpMonthsData()
	if day < 1 || day > npMonthData[year-2000].monthData[month-1] {
		return false
	}
	return true
}

// counts and returns total days from the nepali date 2000-01-01
func getTotalDaysFromNepaliDate(year int, month int, day int) int {
	var totalDays int = day - 1
	npMonthData := getNpMonthsData()

	// adding days of months of 2000
	var yearIndex int = year - 2000
	for i := 0; i < month-1; i++ {
		totalDays = totalDays + npMonthData[yearIndex].monthData[i]
	}

	// adding days of year
	for i := 0; i < yearIndex; i++ {
		totalDays = totalDays + npMonthData[i].yearDays
	}
	return totalDays
}

// checks if the english year is leap year or not"""
func isLeapYear(year int) bool {
	if year%4 == 0 {
		if year%100 == 0 {
			return (year%400 == 0)
		}
		return true
	}
	return false
}

// Public methods

// Converts english date to nepali.
// Accepts the input parameters year, month, day
// Returns multi value of year, month, day
func EnglishToNepali(year int, month int, day int) (npYear int, npMonth int, npDay int) {
	// VALIDATION
	// checking if date is in range
	if !checkEnglishDate(year, month, day) {
		return  // TODO: raise an error here
	}

	// REFERENCE
	// Setting nepali reference to 2000/01/01 with english date 1943/04/14
	npYear = 2000
	npMonth = 1
	npDay = 1

	// DIFFERENCE
	// calculating difference days from 1943/04/14
	var difference int = int(math.Abs(float64(getTotalDaysFromEnglishDate(year, month, day) - getTotalDaysFromEnglishDate(1943, 4, 14))))

	// YEAR
	// Incrementing year until the difference remains less than 365
	npMonthData := getNpMonthsData()
	var yearDataIndex int = 0
	for difference >= npMonthData[yearDataIndex].yearDays {
		difference -= npMonthData[yearDataIndex].yearDays
		npYear += 1
		yearDataIndex += 1
	}

	// MONTH
	// Incrementing month until the difference remains less than next nepali month days (mostly 31)
	var i int = 0
	for difference >= npMonthData[yearDataIndex].monthData[i] {
		difference -= npMonthData[yearDataIndex].monthData[i]
		npMonth += 1
		i += 1
	}

	// DAY
	// Remaining difference is the day
	npDay += difference

	return
}

// Converts nepali date to english.
// Accepts the input parameters year, month, day
// Returns multi value of year, month, day
func NepaliToEnglish(year int, month int, day int) (enYear int, enMonth int, enDay int) {
	// VALIDATION
	// checking if date is in range
	if !checkNepaliDate(year, month, day) {
		return  // TODO: raise an error here
	}

	// REFERENCE
	// Setting english reference to 1944/01/01 with nepali date 2000/09/17
	enYear = 1944
	enMonth = 1
	enDay = 1

	// DIFFERENCE
	// calculating difference days from 2000/09/17
	var difference int = int(math.Abs(float64(getTotalDaysFromNepaliDate(year, month, day) - getTotalDaysFromNepaliDate(2000, 9, 17))))

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
	monthDays := getEnMonths(isLeapYear(enYear))
	var i int = 0
	for difference >= monthDays[i] {
		difference -= monthDays[i]
		enMonth += 1
		i += 1
	}

	// DAY
	// Remaining difference is the day
	enDay += difference

	return
}
