package nepalitime_test

import (
	"testing"
	"time"

	"github.com/opensource-nepal/nepali/constants"
	"github.com/opensource-nepal/nepali/nepalitime"
	"github.com/stretchr/testify/assert"
)

func TestDateFunctionWithInvalidDateAndTimeArgument(t *testing.T) {
	nt, err := nepalitime.Date(1, 2, 3, 4, 5, 6, 7)
	assert.Nil(t, nt)
	assert.NotNil(t, err)
}

func TestDateFunctionWithValidDateAndTimeArgument(t *testing.T) {
	nt, err := nepalitime.Date(2079, 10, 15, 14, 29, 6, 7)
	assert.NotNil(t, nt)
	assert.Nil(t, err)
}

func TestFromEnglishTimeWithUnsupportedDate(t *testing.T) {
	enTime := time.Date(2083, 1, 29, 14, 29, 6, 7, nepalitime.GetNepaliLocation()) // 2083 is not supported by nepalitime
	nt, err := nepalitime.FromEnglishTime(enTime)
	assert.Nil(t, nt)
	assert.NotNil(t, err)
}

func TestFromEnglishTimeWithSupportedDate(t *testing.T) {
	enTime := time.Date(2023, 1, 29, 14, 29, 6, 7, nepalitime.GetNepaliLocation())
	nt, err := nepalitime.FromEnglishTime(enTime)
	assert.NotNil(t, nt)
	assert.Nil(t, err)
}

func TestFromEnglishTimeWithSingaporeTimezone(t *testing.T) {
	loc, _ := time.LoadLocation("Asia/Singapore")
	enTime := time.Date(2023, 1, 29, 1, 15, 0, 0, loc) // 2079/10/14 11:00 PM in Nepal
	nt, _ := nepalitime.FromEnglishTime(enTime)
	assert.Equal(t, nt.Year(), 2079)
	assert.Equal(t, nt.Month(), 10)
	assert.Equal(t, nt.Day(), 14)
	assert.Equal(t, nt.Hour(), 23)
	assert.Equal(t, nt.Minute(), 0)
}

func TestFromEnglishTimeWithNewYorkTimezone(t *testing.T) {
	loc, _ := time.LoadLocation("America/New_York")
	enTime := time.Date(2023, 1, 29, 14, 15, 0, 0, loc) // 2079/10/16 1:00 AM in Nepal
	nt, _ := nepalitime.FromEnglishTime(enTime)
	assert.Equal(t, nt.Year(), 2079)
	assert.Equal(t, nt.Month(), 10)
	assert.Equal(t, nt.Day(), 16)
	assert.Equal(t, nt.Hour(), 1)
	assert.Equal(t, nt.Minute(), 0)
}

func TestNow(t *testing.T) {
	npTime := nepalitime.Now()
	assert.NotNil(t, npTime)
	// TODO: Mocking needed for further check
}

func TestGetCurrentEnglishTime(t *testing.T) {
	enTime := nepalitime.GetCurrentEnglishTime()
	assert.NotNil(t, enTime)
	// TODO: Mocking needed for further check
}

func TestGetNepaliLocation(t *testing.T) {
	loc := nepalitime.GetNepaliLocation()
	assert.Equal(t, loc.String(), constants.Timezone)
}
