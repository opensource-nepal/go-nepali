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

func TestNepaliTimeFormatYmdSlash(t *testing.T) {
	res, _ := globalNepaliTime.Format("%Y/%m/%d")

	assert.Equal(t, "2079/10/14", res, "%Y/%m/%d did not match")
}

func TestNepaliTimeFormatymdSlash(t *testing.T) {
	res, _ := globalNepaliTime.Format("%y/%m/%d")

	assert.Equal(t, "79/10/14", res, "%y/%m/%d did not match")
}

func TestNepaliTimeFormatdmYSlash(t *testing.T) {
	res, _ := globalNepaliTime.Format("%d/%m/%Y")

	assert.Equal(t, "14/10/2079", res, "%d/%m/%Y did not match")
}

func TestNepaliTimeFormatdmySlash(t *testing.T) {
	res, _ := globalNepaliTime.Format("%d/%m/%y")

	assert.Equal(t, "14/10/79", res, "%d/%m/%y did not match")
}

func TestNepaliTimeFormatmdYSlash(t *testing.T) {
	res, _ := globalNepaliTime.Format("%m/%d/%Y")

	assert.Equal(t, "10/14/2079", res, "%d/%m/%Y did not match")
}

func TestNepaliTimeFormatmdySlash(t *testing.T) {
	res, _ := globalNepaliTime.Format("%m/%d/%y")

	assert.Equal(t, "10/14/79", res, "%d/%m/%y did not match")
}

func TestNepaliTimeFormatYmdDash(t *testing.T) {
	res, _ := globalNepaliTime.Format("%Y-%m-%d")

	assert.Equal(t, "2079-10-14", res, "%Y-%m-%d did not match")
}

func TestNepaliTimeFormatymdDash(t *testing.T) {
	res, _ := globalNepaliTime.Format("%y-%m-%d")

	assert.Equal(t, "79-10-14", res, "%y-%m-%d did not match")
}

func TestNepaliTimeFormatdmYDash(t *testing.T) {
	res, _ := globalNepaliTime.Format("%d-%m-%Y")

	assert.Equal(t, "14-10-2079", res, "%d-%m-%Y did not match")
}

func TestNepaliTimeFormatdmyDash(t *testing.T) {
	res, _ := globalNepaliTime.Format("%d-%m-%y")

	assert.Equal(t, "14-10-79", res, "%d-%m-%y did not match")
}

func TestNepaliTimeFormatmdYDash(t *testing.T) {
	res, _ := globalNepaliTime.Format("%m-%d-%Y")

	assert.Equal(t, "10-14-2079", res, "%m-%d-%Y did not match")
}

func TestNepaliTimeFormatmdyDash(t *testing.T) {
	res, _ := globalNepaliTime.Format("%m-%d-%y")

	assert.Equal(t, "10-14-79", res, "%m-%d-%y did not match")
}

func TestNepaliTimeFormatYmdHMSSlash(t *testing.T) {
	res, _ := globalNepaliTime.Format("%Y/%m/%d %H:%M:%S")

	assert.Equal(t, "2079/10/14 16:23:17", res, "%Y/%m/%d %H:%M:%S did not match")
}

func TestNepaliTimeFormatYmdIMSPSlash(t *testing.T) {
	res, _ := globalNepaliTime.Format("%Y/%m/%d %I:%M:%S %p")

	assert.Equal(t, "2079/10/14 04:23:17 PM", res, "%Y/%m/%d %I:%M:%S %p did not match")
}

func TestNepaliTimeFormatYmdIMSWithoutAMPMSlash(t *testing.T) {
	res, _ := globalNepaliTime.Format("%Y/%m/%d %I:%M:%S")

	assert.Equal(t, "2079/10/14 04:23:17", res, "%Y/%m/%d %I:%M:%S did not match")
}

func TestNepaliTimeFormatYmdIMSfSlash(t *testing.T) {
	res, _ := globalNepaliTimeLeadingZeros.Format("%Y/%m/%d %I:%M:%S:%f")

	assert.Equal(t, "2079/01/02 03:04:05:000111", res, "%Y/%m/%d %I:%M:%S:%f did not match")
}

func TestNepaliTimeFormatYmdSlashNoLeadingZero(t *testing.T) {
	res, _ := globalNepaliTimeLeadingZeros.Format("%Y/%-m/%-d")

	assert.Equal(t, "2079/1/2", res, "%Y/%-m/%-d did not match")
}

func TestNepaliTimeFormatymdSlashNoLeadingZero(t *testing.T) {
	res, _ := globalNepaliTimeLeadingZeros.Format("%y/%-m/%-d")

	assert.Equal(t, "79/1/2", res, "%y/%-m/%-d did not match")
}

func TestNepaliTimeFormatdmYSlashNoLeadingZero(t *testing.T) {
	res, _ := globalNepaliTimeLeadingZeros.Format("%-d/%-m/%Y")

	assert.Equal(t, "2/1/2079", res, "%-d/%-m/%Y did not match")
}

func TestNepaliTimeFormatdmySlashNoLeadingZero(t *testing.T) {
	res, _ := globalNepaliTimeLeadingZeros.Format("%-d/%-m/%y")

	assert.Equal(t, "2/1/79", res, "%-d/%-m/%y did not match")
}

func TestNepaliTimeFormatmdYSlashNoLeadingZero(t *testing.T) {
	res, _ := globalNepaliTimeLeadingZeros.Format("%-m/%-d/%Y")

	assert.Equal(t, "1/2/2079", res, "%-d/%-m/%Y did not match")
}

func TestNepaliTimeFormatmdySlashNoLeadingZero(t *testing.T) {
	res, _ := globalNepaliTimeLeadingZeros.Format("%-m/%-d/%y")

	assert.Equal(t, "1/2/79", res, "%-d/%-m/%y did not match")
}

func TestNepaliTimeFormatYmdIMSfSlashNoLeadingZero(t *testing.T) {
	res, _ := globalNepaliTimeLeadingZeros.Format("%Y/%-m/%-d %-I:%-M:%-S:%-f")

	assert.Equal(t, "2079/1/2 3:4:5:111", res, "%Y/%-m/%-d %-I:%-M:%-S:%-f did not match")
}

func TestNepaliTimeFormatYmdDashNoLeadingZero(t *testing.T) {
	res, _ := globalNepaliTimeLeadingZeros.Format("%Y-%-m-%-d")

	assert.Equal(t, "2079-1-2", res, "%Y-%-m-%-d did not match")
}

func TestNepaliTimeFormatymdDashNoLeadingZero(t *testing.T) {
	res, _ := globalNepaliTimeLeadingZeros.Format("%y-%-m-%-d")

	assert.Equal(t, "79-1-2", res, "%y-%-m-%-d did not match")
}

func TestNepaliTimeFormatdmYDashNoLeadingZero(t *testing.T) {
	res, _ := globalNepaliTimeLeadingZeros.Format("%-d-%-m-%Y")

	assert.Equal(t, "2-1-2079", res, "%-d-%-m-%Y did not match")
}

func TestNepaliTimeFormatdmyDashLeadingZero(t *testing.T) {
	res, _ := globalNepaliTimeLeadingZeros.Format("%-d-%-m-%y")

	assert.Equal(t, "2-1-79", res, "%-d-%-m-%y did not match")
}

func TestNepaliTimeFormatmdYDashNoLeadingZero(t *testing.T) {
	res, _ := globalNepaliTimeLeadingZeros.Format("%-m-%-d-%Y")

	assert.Equal(t, "1-2-2079", res, "%-d-%-m-%Y did not match")
}

func TestNepaliTimeFormatmdyDashNoLeadingZero(t *testing.T) {
	res, _ := globalNepaliTimeLeadingZeros.Format("%-m-%-d-%y")

	assert.Equal(t, "1-2-79", res, "%-d-%-m-%y did not match")
}

func TestNepaliTimeFormatYmdHMSSlashNoLeadingZero(t *testing.T) {
	res, _ := globalNepaliTimeLeadingZeros.Format("%Y/%-m/%-d %-H:%-M:%-S")

	assert.Equal(t, "2079/1/2 3:4:5", res, "%Y/%-m/%-d %-H:%-M:%-S did not match")
}

func TestNepaliTimeFormatYmdIMSfDashNoLeadingZero(t *testing.T) {
	res, _ := globalNepaliTimeLeadingZeros.Format("%Y-%-m-%-d %-I-%-M-%-S-%-f")

	assert.Equal(t, "2079-1-2 3-4-5-111", res, "%Y-%-m-%-d %-I-%-M-%-S-%-f did not match")
}

func TestNepaliTimeFormatSpecialCharacters(t *testing.T) {
	res, _ := globalNepaliTime.Format("%Y/%m/%d %I:%M:%S $*#()%%")

	assert.Equal(t, "2079/10/14 04:23:17 $*#()%", res, "%Y/%m/%d %I:%M:%S $*#()%% did not match")
}

func TestNepaliTimeFormatSpecialCharactersBetweenYearAndMonth(t *testing.T) {
	res, _ := globalNepaliTime.Format("%Y/$*#()%%%m/%d")

	assert.Equal(t, "2079/$*#()%10/14", res, "%Y/%m/%d %I:%M:%S:%f did not match")
}

func TestNepaliTimeFormatErrornouseFormats(t *testing.T) {
	res, err := globalNepaliTime.Format("%k")

	assert.Equal(t, res, "", "response should be empty string")
	assert.EqualError(t, err, "error while formatting NepaliTime with given format", "error message not matched")
}

func TestNepaliTimeFormatErrornouseFormatsNoLeadingZeros(t *testing.T) {
	res, err := globalNepaliTime.Format("%-k")

	assert.Equal(t, res, "", "response should be empty string")
	assert.EqualError(t, err, "error while formatting NepaliTime with given format", "error message not matched")
}

func TestNepaliTimeFormatWithoutYear(t *testing.T) {
	res, _ := globalNepaliTime.Format("%m/%d")

	assert.Equal(t, res, "10/14", "%m/%d did not match")
}

func TestNepaliTimeFormatHourLessThan10(t *testing.T) {
	date, _ := nepalitime.Date(2079, 11, 4, 9, 10, 11, 123)

	res, _ := date.Format("%Y/%m/%d %H::%M::%S::%f")

	assert.Equal(t, "2079/11/04 09::10::11::000123", res, "%Y/%m/%d %H::%M::%S::%f did not match")
}

func TestNepaliTimeFormatHourEqualTo0(t *testing.T) {
	date, _ := nepalitime.Date(2079, 11, 4, 0, 10, 11, 123)

	res, _ := date.Format("%Y/%m/%d %I::%M::%S::%f")

	assert.Equal(t, "2079/11/04 12::10::11::000123", res, "%Y/%m/%d %I::%M::%S::%f did not match")
}

func TestNepaliTimeFormatHourGreaterThan12(t *testing.T) {
	date, _ := nepalitime.Date(2079, 11, 4, 13, 10, 11, 123)

	res, _ := date.Format("%Y/%m/%d %-I::%M::%S::%f")

	assert.Equal(t, "2079/11/04 1::10::11::000123", res, "%Y/%m/%d %-I::%M::%S::%f did not match")
}
