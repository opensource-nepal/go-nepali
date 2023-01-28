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

func TestEnglishToNepaliReturnErrorOnInvalidMonth(t *testing.T) {
	// month 0
	date, err := dateConverter.EnglishToNepali(2023, 0, 4)
	assert.Nil(t, date)
	assert.NotNil(t, err)

	// month greater than 12
	date, err = dateConverter.EnglishToNepali(2023, 13, 4)
	assert.Nil(t, date)
	assert.NotNil(t, err)
}

func TestEnglishToNepaliReturnErrorOnInvalidDay(t *testing.T) {
	// day 0
	date, err := dateConverter.EnglishToNepali(2023, 1, 0)
	assert.Nil(t, date)
	assert.NotNil(t, err)

	// day greater than month's max day (eg. 40)
	date, err = dateConverter.EnglishToNepali(2023, 1, 40)
	assert.Nil(t, date)
	assert.NotNil(t, err)
}

func TestEnglishToNepaliReturnsValidPastNepaliDate(t *testing.T) {
	date, err := dateConverter.EnglishToNepali(1994, 8, 13)
	assert.NotNil(t, date)
	assert.Nil(t, err)
	assert.Equal(t, date[0], 2051)
	assert.Equal(t, date[1], 4)
	assert.Equal(t, date[2], 29)
}

func TestEnglishToNepaliReturnsValidRecentNepaliDate(t *testing.T) {
	date, err := dateConverter.EnglishToNepali(2023, 1, 28)
	assert.NotNil(t, date)
	assert.Nil(t, err)
	assert.Equal(t, date[0], 2079)
	assert.Equal(t, date[1], 10)
	assert.Equal(t, date[2], 14)
}

func TestEnglishToNepaliReturnsValidFutureNepaliDate(t *testing.T) {
	date, err := dateConverter.EnglishToNepali(2030, 11, 26)
	assert.NotNil(t, date)
	assert.Nil(t, err)
	assert.Equal(t, date[0], 2087)
	assert.Equal(t, date[1], 8)
	assert.Equal(t, date[2], 10)
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

func TestNepaliToEnglishReturnErrorOnInvalidMonth(t *testing.T) {
	// month 0
	date, err := dateConverter.NepaliToEnglish(2079, 0, 4)
	assert.Nil(t, date)
	assert.NotNil(t, err)

	// month greater than 12
	date, err = dateConverter.NepaliToEnglish(2079, 13, 4)
	assert.Nil(t, date)
	assert.NotNil(t, err)
}

func TestNepaliToEnglishReturnErrorOnInvalidDay(t *testing.T) {
	// day 0
	date, err := dateConverter.NepaliToEnglish(2079, 1, 0)
	assert.Nil(t, date)
	assert.NotNil(t, err)

	// day greater than month's max day (eg. 40)
	date, err = dateConverter.NepaliToEnglish(2079, 1, 40)
	assert.Nil(t, date)
	assert.NotNil(t, err)
}

func TestNepaliToEnglishReturnsValidPastEnglishDate(t *testing.T) {
	date, err := dateConverter.NepaliToEnglish(2051, 4, 29)
	assert.NotNil(t, date)
	assert.Nil(t, err)
	assert.Equal(t, date[0], 1994)
	assert.Equal(t, date[1], 8)
	assert.Equal(t, date[2], 13)
}

func TestNepaliToEnglishReturnsValidRecentEnglishDate(t *testing.T) {
	date, err := dateConverter.NepaliToEnglish(2079, 10, 14)
	assert.NotNil(t, date)
	assert.Nil(t, err)
	assert.Equal(t, date[0], 2023)
	assert.Equal(t, date[1], 1)
	assert.Equal(t, date[2], 28)
}

func TestNepaliToEnglishReturnsValidFutureEnglishDate(t *testing.T) {
	date, err := dateConverter.NepaliToEnglish(2087, 8, 10)
	assert.NotNil(t, date)
	assert.Nil(t, err)
	assert.Equal(t, date[0], 2030)
	assert.Equal(t, date[1], 11)
	assert.Equal(t, date[2], 26)
}
