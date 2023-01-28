// Unit tests for date converter
//
// To run only this test
// go test -v github.com/sugat009/nepali/dateConverter
package dateConverter_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sugat009/nepali/dateConverter"
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
