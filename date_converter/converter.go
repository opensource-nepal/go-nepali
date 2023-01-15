// This package contains the date converter functions.
// Reference date for conversion is 2000/01/01 BS and 1943/4/14 AD
//
// USAGE:
// y, m, d := EnglishToNepali(2023, 1, 15)
// y, m, d := NepaliToEnglish(2079, 10, 1)
package date_converter

import (
	"math"
)

type NepaliMonthData struct {
	monthData [12]int
	yearDays  int
}

func _getEnMonths(isLeapYear bool) [12]int {
	if isLeapYear {
		return [12]int{31, 29, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	}
	return [12]int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
}

func _getNpMothsData() []NepaliMonthData {
	return []NepaliMonthData{
		NepaliMonthData{[12]int{30, 32, 31, 32, 31, 30, 30, 30, 29, 30, 29, 31}, 365}, // 2000 BS - 1944 AD
		NepaliMonthData{[12]int{31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30}, 365}, // 2001 BS
		NepaliMonthData{[12]int{31, 31, 32, 32, 31, 30, 30, 29, 30, 29, 30, 30}, 365},
		NepaliMonthData{[12]int{31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 31}, 366},
		NepaliMonthData{[12]int{30, 32, 31, 32, 31, 30, 30, 30, 29, 30, 29, 31}, 365},
		NepaliMonthData{[12]int{31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30}, 365},
		NepaliMonthData{[12]int{31, 31, 32, 32, 31, 30, 30, 29, 30, 29, 30, 30}, 365},
		NepaliMonthData{[12]int{31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 31}, 366},
		NepaliMonthData{[12]int{31, 31, 31, 32, 31, 31, 29, 30, 30, 29, 29, 31}, 365},
		NepaliMonthData{[12]int{31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30}, 365},
		NepaliMonthData{[12]int{31, 31, 32, 32, 31, 30, 30, 29, 30, 29, 30, 30}, 365},
		NepaliMonthData{[12]int{31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 31}, 366},
		NepaliMonthData{[12]int{31, 31, 31, 32, 31, 31, 29, 30, 30, 29, 30, 30}, 365},
		NepaliMonthData{[12]int{31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30}, 365},
		NepaliMonthData{[12]int{31, 31, 32, 32, 31, 30, 30, 29, 30, 29, 30, 30}, 365},
		NepaliMonthData{[12]int{31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 31}, 366},
		NepaliMonthData{[12]int{31, 31, 31, 32, 31, 31, 29, 30, 30, 29, 30, 30}, 365},
		NepaliMonthData{[12]int{31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30}, 365},
		NepaliMonthData{[12]int{31, 32, 31, 32, 31, 30, 30, 29, 30, 29, 30, 30}, 365},
		NepaliMonthData{[12]int{31, 32, 31, 32, 31, 30, 30, 30, 29, 30, 29, 31}, 366},
		NepaliMonthData{[12]int{31, 31, 31, 32, 31, 31, 30, 29, 30, 29, 30, 30}, 365},
		NepaliMonthData{[12]int{31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30}, 365},
		NepaliMonthData{[12]int{31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 30}, 365},
		NepaliMonthData{[12]int{31, 32, 31, 32, 31, 30, 30, 30, 29, 30, 29, 31}, 366},
		NepaliMonthData{[12]int{31, 31, 31, 32, 31, 31, 30, 29, 30, 29, 30, 30}, 365},
		NepaliMonthData{[12]int{31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30}, 365},
		NepaliMonthData{[12]int{31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 31}, 366},
		NepaliMonthData{[12]int{30, 32, 31, 32, 31, 30, 30, 30, 29, 30, 29, 31}, 365},
		NepaliMonthData{[12]int{31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30}, 365},
		NepaliMonthData{[12]int{31, 31, 32, 31, 32, 30, 30, 29, 30, 29, 30, 30}, 365},
		NepaliMonthData{[12]int{31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 31}, 366},
		NepaliMonthData{[12]int{30, 32, 31, 32, 31, 30, 30, 30, 29, 30, 29, 31}, 365},
		NepaliMonthData{[12]int{31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30}, 365},
		NepaliMonthData{[12]int{31, 31, 32, 32, 31, 30, 30, 29, 30, 29, 30, 30}, 365},
		NepaliMonthData{[12]int{31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 31}, 366},
		NepaliMonthData{[12]int{30, 32, 31, 32, 31, 31, 29, 30, 30, 29, 29, 31}, 365},
		NepaliMonthData{[12]int{31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30}, 365},
		NepaliMonthData{[12]int{31, 31, 32, 32, 31, 30, 30, 29, 30, 29, 30, 30}, 365},
		NepaliMonthData{[12]int{31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 31}, 366},
		NepaliMonthData{[12]int{31, 31, 31, 32, 31, 31, 29, 30, 30, 29, 30, 30}, 365},
		NepaliMonthData{[12]int{31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30}, 365},
		NepaliMonthData{[12]int{31, 31, 32, 32, 31, 30, 30, 29, 30, 29, 30, 30}, 365},
		NepaliMonthData{[12]int{31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 31}, 366},
		NepaliMonthData{[12]int{31, 31, 31, 32, 31, 31, 29, 30, 30, 29, 30, 30}, 365},
		NepaliMonthData{[12]int{31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30}, 365},
		NepaliMonthData{[12]int{31, 32, 31, 32, 31, 30, 30, 29, 30, 29, 30, 30}, 365},
		NepaliMonthData{[12]int{31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 31}, 366},
		NepaliMonthData{[12]int{31, 31, 31, 32, 31, 31, 30, 29, 30, 29, 30, 30}, 365},
		NepaliMonthData{[12]int{31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30}, 365},
		NepaliMonthData{[12]int{31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 30}, 365},
		NepaliMonthData{[12]int{31, 32, 31, 32, 31, 30, 30, 30, 29, 30, 29, 31}, 366},
		NepaliMonthData{[12]int{31, 31, 31, 32, 31, 31, 30, 29, 30, 29, 30, 30}, 365},
		NepaliMonthData{[12]int{31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30}, 365},
		NepaliMonthData{[12]int{31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 30}, 365},
		NepaliMonthData{[12]int{31, 32, 31, 32, 31, 30, 30, 30, 29, 30, 29, 31}, 366},
		NepaliMonthData{[12]int{31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30}, 365},
		NepaliMonthData{[12]int{31, 31, 32, 31, 32, 30, 30, 29, 30, 29, 30, 30}, 365},
		NepaliMonthData{[12]int{31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 31}, 366},
		NepaliMonthData{[12]int{30, 32, 31, 32, 31, 30, 30, 30, 29, 30, 29, 31}, 365},
		NepaliMonthData{[12]int{31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30}, 365},
		NepaliMonthData{[12]int{31, 31, 32, 32, 31, 30, 30, 29, 30, 29, 30, 30}, 365},
		NepaliMonthData{[12]int{31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 31}, 366},
		NepaliMonthData{[12]int{30, 32, 31, 32, 31, 31, 29, 30, 29, 30, 29, 31}, 365},
		NepaliMonthData{[12]int{31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30}, 365},
		NepaliMonthData{[12]int{31, 31, 32, 32, 31, 30, 30, 29, 30, 29, 30, 30}, 365},
		NepaliMonthData{[12]int{31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 31}, 366},
		NepaliMonthData{[12]int{31, 31, 31, 32, 31, 31, 29, 30, 30, 29, 29, 31}, 365},
		NepaliMonthData{[12]int{31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30}, 365},
		NepaliMonthData{[12]int{31, 31, 32, 32, 31, 30, 30, 29, 30, 29, 30, 30}, 365},
		NepaliMonthData{[12]int{31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 31}, 366},
		NepaliMonthData{[12]int{31, 31, 31, 32, 31, 31, 29, 30, 30, 29, 30, 30}, 365},
		NepaliMonthData{[12]int{31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30}, 365},
		NepaliMonthData{[12]int{31, 32, 31, 32, 31, 30, 30, 29, 30, 29, 30, 30}, 365},
		NepaliMonthData{[12]int{31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 31}, 366},
		NepaliMonthData{[12]int{31, 31, 31, 32, 31, 31, 30, 29, 30, 29, 30, 30}, 365},
		NepaliMonthData{[12]int{31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30}, 365},
		NepaliMonthData{[12]int{31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 30}, 365},
		NepaliMonthData{[12]int{31, 32, 31, 32, 31, 30, 30, 30, 29, 30, 29, 31}, 366},
		NepaliMonthData{[12]int{31, 31, 31, 32, 31, 31, 30, 29, 30, 29, 30, 30}, 365},
		NepaliMonthData{[12]int{31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30}, 365},
		NepaliMonthData{[12]int{31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 30}, 365},
		NepaliMonthData{[12]int{31, 31, 32, 32, 31, 30, 30, 30, 29, 30, 30, 30}, 366},
		NepaliMonthData{[12]int{30, 32, 31, 32, 31, 30, 30, 30, 29, 30, 30, 30}, 365},
		NepaliMonthData{[12]int{31, 31, 32, 31, 31, 30, 30, 30, 29, 30, 30, 30}, 365},
		NepaliMonthData{[12]int{31, 31, 32, 31, 31, 30, 30, 30, 29, 30, 30, 30}, 365},
		NepaliMonthData{[12]int{31, 32, 31, 32, 30, 31, 30, 30, 29, 30, 30, 30}, 366},
		NepaliMonthData{[12]int{30, 32, 31, 32, 31, 30, 30, 30, 29, 30, 30, 30}, 365},
		NepaliMonthData{[12]int{31, 31, 32, 31, 31, 31, 30, 30, 29, 30, 30, 30}, 366},
		NepaliMonthData{[12]int{30, 31, 32, 32, 30, 31, 30, 30, 29, 30, 30, 30}, 365},
		NepaliMonthData{[12]int{30, 32, 31, 32, 31, 30, 30, 30, 29, 30, 30, 30}, 365},
		NepaliMonthData{[12]int{30, 32, 31, 32, 31, 30, 30, 30, 29, 30, 30, 30}, 365},
		NepaliMonthData{[12]int{31, 31, 32, 31, 31, 31, 30, 30, 29, 30, 30, 30}, 366},
		NepaliMonthData{[12]int{30, 31, 32, 32, 31, 30, 30, 30, 29, 30, 30, 30}, 365},
		NepaliMonthData{[12]int{30, 32, 31, 32, 31, 30, 30, 30, 29, 30, 30, 30}, 365},
		NepaliMonthData{[12]int{31, 31, 32, 31, 31, 30, 30, 30, 29, 30, 30, 30}, 365},
		NepaliMonthData{[12]int{31, 31, 32, 31, 31, 31, 30, 29, 30, 30, 30, 30}, 366},
		NepaliMonthData{[12]int{30, 31, 32, 32, 31, 30, 30, 29, 30, 29, 30, 30}, 364},
		NepaliMonthData{[12]int{31, 32, 31, 32, 31, 30, 30, 30, 29, 30, 30, 30}, 366},
		NepaliMonthData{[12]int{31, 31, 32, 31, 31, 31, 29, 30, 29, 30, 29, 31}, 365},
		NepaliMonthData{[12]int{31, 31, 32, 31, 31, 31, 30, 29, 29, 30, 30, 30}, 365}, // 2099 BS - 2042 AD
	}
}

// ENGLISH DATE CONVERSION

// checks if english date in within range 1944 - 2042
func _checkEnglishDate(year int, month int, day int) bool {
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
func _getTotalDaysFromEnglish_date(year int, month int, day int) int {
	var total_days int = year*365 + day
	enMonths := _getEnMonths(false)
	for i := 0; i < month-1; i++ {
		total_days = total_days + enMonths[i]
	}

	// adding leap days (ie. leap year count)
	if month <= 2 { // checking February month (where leap exists)
		year -= 1
	}
	total_days += int(year/4) - int(year/100) + int(year/400)

	return total_days
}

// returns difference days between two english dates
func _getDaysBetweenEnglishDates(year1 int, month1 int, day1 int, year2 int, month2 int, day2 int) int {
	return int(math.Abs(float64(_getTotalDaysFromEnglish_date(year1, month1, day1) - _getTotalDaysFromEnglish_date(year2, month2, day2))))
}

// Public methods

// Converts english date to nepali.
// Accepts the input parameters year, month, day
// Returns multi value of year, month, day
func EnglishToNepali(year int, month int, day int) (npYear int, npMonth int, npDay int) {
	// VALIDATION
	// checking if date is in range
	if !_checkEnglishDate(year, month, day) {
		panic("Date is out of range")
	}

	// REFERENCE
	// Setting nepali reference to 2000/01/01 with english date 1943/04/14
	npYear = 2000
	npMonth = 1
	npDay = 1

	// DIFFERENCE
	// calculating difference days from 1943/04/14
	var difference int = _getDaysBetweenEnglishDates(year, month, day, 1943, 4, 14)

	// YEAR
	// Incrementing year until the difference remains less than 365
	npMonthData := _getNpMothsData()
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
		difference = difference - npMonthData[yearDataIndex].monthData[i]
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
	enYear = year
	enMonth = month
	enDay = day
	panic("Not implemented yet!")
}
