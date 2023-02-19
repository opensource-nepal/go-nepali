// Unit tests for date converter
//
// To run only this test
// go test -v github.com/opensource-nepal/go-nepali/dateConverter
package dateConverter_test

import (
	"testing"

	"github.com/opensource-nepal/go-nepali/dateConverter"
	"github.com/stretchr/testify/assert"
)

// EnglishToNepali

func TestEnglishToNepaliReturnErrorOnMaxYearRange(t *testing.T) {
	date, err := dateConverter.EnglishToNepali(2060, 1, 4)
	assert.Nil(t, date)
	assert.NotNil(t, err)
}

func TestEnglishToNepaliReturnErrorOnMinYearRange(t *testing.T) {
	date, err := dateConverter.EnglishToNepali(1920, 1, 4)
	assert.Nil(t, date)
	assert.NotNil(t, err)
}

func TestEnglishToNepaliReturnErrorOnMinMonthRange(t *testing.T) {
	date, err := dateConverter.EnglishToNepali(2023, 0, 4)
	assert.Nil(t, date)
	assert.NotNil(t, err)
}

func TestEnglishToNepaliReturnErrorOnMaxMonthRange(t *testing.T) {
	date, err := dateConverter.EnglishToNepali(2023, 13, 4)
	assert.Nil(t, date)
	assert.NotNil(t, err)
}

func TestEnglishToNepaliReturnErrorOnMinDayRange(t *testing.T) {
	date, err := dateConverter.EnglishToNepali(2023, 1, 0)
	assert.Nil(t, date)
	assert.NotNil(t, err)
}

func TestEnglishToNepaliReturnErrorOnMaxDayRange(t *testing.T) {
	date, err := dateConverter.EnglishToNepali(2023, 1, 40)
	assert.Nil(t, date)
	assert.NotNil(t, err)
}

func TestEnglishToNepaliReturnsValidPastNepaliDate(t *testing.T) {
	date, err := dateConverter.EnglishToNepali(1994, 8, 13)
	assert.Nil(t, err)
	assert.EqualValues(t, *date, [3]int{2051, 4, 29})
}

func TestEnglishToNepaliReturnsValidRecentNepaliDate(t *testing.T) {
	date, err := dateConverter.EnglishToNepali(2023, 1, 28)
	assert.Nil(t, err)
	assert.EqualValues(t, *date, [3]int{2079, 10, 14})
}

func TestEnglishToNepaliReturnsValidFutureNepaliDate(t *testing.T) {
	date, err := dateConverter.EnglishToNepali(2030, 11, 26)
	assert.Nil(t, err)
	assert.EqualValues(t, *date, [3]int{2087, 8, 10})
}

func TestEnglishToNepaliForMinEdgeDate(t *testing.T) {
	date, err := dateConverter.EnglishToNepali(1944, 1, 1)
	assert.Nil(t, err)
	assert.EqualValues(t, *date, [3]int{2000, 9, 17})
}

func TestEnglishToNepaliForMaxEdgeDate(t *testing.T) {
	date, err := dateConverter.EnglishToNepali(2042, 12, 31)
	assert.Nil(t, err)
	assert.EqualValues(t, *date, [3]int{2099, 9, 16})
}

// NepaliToEnglish

func TestNepaliToEnglishReturnErrorOnMaxYearRange(t *testing.T) {
	date, err := dateConverter.NepaliToEnglish(3000, 1, 4)
	assert.Nil(t, date)
	assert.NotNil(t, err)
}

func TestNepaliToEnglishReturnErrorOnMinYearRange(t *testing.T) {
	date, err := dateConverter.NepaliToEnglish(1920, 1, 4)
	assert.Nil(t, date)
	assert.NotNil(t, err)
}

func TestNepaliToEnglishReturnErrorOnMinMonthRange(t *testing.T) {
	date, err := dateConverter.NepaliToEnglish(2079, 0, 4)
	assert.Nil(t, date)
	assert.NotNil(t, err)
}

func TestNepaliToEnglishReturnErrorOnMaxMonthRange(t *testing.T) {
	date, err := dateConverter.NepaliToEnglish(2079, 13, 4)
	assert.Nil(t, date)
	assert.NotNil(t, err)
}

func TestNepaliToEnglishReturnErrorOnMinDayRange(t *testing.T) {
	date, err := dateConverter.NepaliToEnglish(2079, 1, 0)
	assert.Nil(t, date)
	assert.NotNil(t, err)
}

func TestNepaliToEnglishReturnErrorOnMaxDayRange(t *testing.T) {
	date, err := dateConverter.NepaliToEnglish(2079, 1, 40)
	assert.Nil(t, date)
	assert.NotNil(t, err)
}

func TestNepaliToEnglishReturnsValidPastEnglishDate(t *testing.T) {
	date, err := dateConverter.NepaliToEnglish(2051, 4, 29)
	assert.Nil(t, err)
	assert.EqualValues(t, *date, [3]int{1994, 8, 13})
}

func TestNepaliToEnglishReturnsValidRecentEnglishDate(t *testing.T) {
	date, err := dateConverter.NepaliToEnglish(2079, 10, 14)
	assert.Nil(t, err)
	assert.EqualValues(t, *date, [3]int{2023, 1, 28})
}

func TestNepaliToEnglishReturnsValidFutureEnglishDate(t *testing.T) {
	date, err := dateConverter.NepaliToEnglish(2087, 8, 10)
	assert.Nil(t, err)
	assert.EqualValues(t, *date, [3]int{2030, 11, 26})
}

func TestNepaliToEnglishReturnsValidEnglishLeapYearDate(t *testing.T) {
	date, err := dateConverter.NepaliToEnglish(2080, 12, 15)
	assert.Nil(t, err)
	assert.EqualValues(t, *date, [3]int{2024, 3, 28})
}

func TestNepaliToEnglishForMinEdgeDate(t *testing.T) {
	date, err := dateConverter.NepaliToEnglish(2000, 1, 1)
	assert.Nil(t, err)
	assert.EqualValues(t, *date, [3]int{1943, 4, 14})
}

func TestNepaliToEnglishForMaxEdgeDate(t *testing.T) {
	date, err := dateConverter.NepaliToEnglish(2099, 12, 30)
	assert.Nil(t, err)
	assert.EqualValues(t, *date, [3]int{2043, 4, 13})
}
