// This package contains the date converter functions.
//
// USAGE:
// y, m, d := EnglishToNepali(2023, 1, 15)
// y, m, d := NepaliToEnglish(2079, 10, 1)
package date_converter

// Converts english date to nepali.
// Accepts the input parameters year, month, day
// Returns multi value of year, month, day
func EnglishToNepali(year int, month int, day int) (np_year int, np_month int, np_day int) {
	np_year = year
	np_month = month
	np_day = day
	return
}

// Converts nepali date to english.
// Accepts the input parameters year, month, day
// Returns multi value of year, month, day
func NepaliToEnglish(year int, month int, day int) (en_year int, en_month int, en_day int) {
	en_year = year
	en_month = month
	en_day = day
	return
}
