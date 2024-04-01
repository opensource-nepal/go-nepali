// Unit tests for NepaliTime object
//
// To run only this test
// go test -v github.com/opensource-nepal/go-nepali/nepalitime
package nepalitime_test

import (
	"testing"
	"time"

	"github.com/opensource-nepal/go-nepali/constants"
	"github.com/opensource-nepal/go-nepali/nepalitime"
	"github.com/stretchr/testify/assert"
)

// Not supposed to be changed.
// Most of the tests depend on it.
// Used for caching purpose.
var globalNepaliTime, _ = nepalitime.Date(2079, 10, 14, 16, 23, 17, 0)
var globalNepaliTimeLeadingZeros, _ = nepalitime.Date(2079, 01, 02, 03, 04, 05, 111)

func TestNepaliTimeString(t *testing.T) {
	assert.Equal(t, globalNepaliTime.String(), "2079-10-14 16:23:17")
}

func TestNepaliTimeStringForNonZeros(t *testing.T) {
	npTime, _ := nepalitime.Date(2079, 1, 4, 6, 3, 0, 0)
	assert.Equal(t, npTime.String(), "2079-01-04 06:03:00")
}

func TestNepaliTimeGetEnglishTime(t *testing.T) {
	enTime := globalNepaliTime.GetEnglishTime()
	assert.Equal(t, enTime.Year(), 2023)
	assert.Equal(t, enTime.Month(), time.Month(1))
	assert.Equal(t, enTime.Day(), 28)
	assert.Equal(t, enTime.Hour(), 16)
	assert.Equal(t, enTime.Minute(), 23)
	assert.Equal(t, enTime.Second(), 17)

	location := enTime.Location()
	assert.Equal(t, location.String(), constants.Timezone)
}

func TestNepaliTimeDate(t *testing.T) {
	y, m, d := globalNepaliTime.Date()
	assert.Equal(t, y, 2079)
	assert.Equal(t, m, 10)
	assert.Equal(t, d, 14)
}

func TestNepaliTimeYear(t *testing.T) {
	assert.Equal(t, globalNepaliTime.Year(), 2079)
}

func TestNepaliTimeMonth(t *testing.T) {
	assert.Equal(t, globalNepaliTime.Month(), 10)
}

func TestNepaliTimeDay(t *testing.T) {
	assert.Equal(t, globalNepaliTime.Day(), 14)
}

func TestNepaliTimeWeekday(t *testing.T) {
	assert.Equal(t, globalNepaliTime.Weekday(), time.Weekday(6))
}

func TestNepaliTimeClock(t *testing.T) {
	h, m, s := globalNepaliTime.Clock()
	assert.Equal(t, h, 16)
	assert.Equal(t, m, 23)
	assert.Equal(t, s, 17)
}

func TestNepaliTimeHour(t *testing.T) {
	assert.Equal(t, globalNepaliTime.Hour(), 16)
}

func TestNepaliTimeMinute(t *testing.T) {
	assert.Equal(t, globalNepaliTime.Minute(), 23)
}

func TestNepaliTimeSecond(t *testing.T) {
	assert.Equal(t, globalNepaliTime.Second(), 17)
}

func TestNepaliTimeNanosecond(t *testing.T) {
	assert.Equal(t, globalNepaliTime.Nanosecond(), 0)
}

func TestNepaliTimeFormat(t *testing.T) {
	res := globalNepaliTime.Format("%Y/%m/%d")

	assert.Equal(t, "2079/10/14", res, "%Y/%m/%d formatting did not match")
}

func TestNepaliTimeFormatReturnAsItIsOnUnknownFormat(t *testing.T) {
	res := globalNepaliTime.Format("%k")

	assert.Equal(t, "%k", res, "Unknown format didn't returned as it is")
}
