package nepalitime_test

import (
	"testing"

	"github.com/opensource-nepal/go-nepali/nepalitime"
	"github.com/stretchr/testify/assert"
)

func TestNepaliFormatterFormatYmdSlash(t *testing.T) {
	formatter := nepalitime.NewFormatter(globalNepaliTime)
	res := formatter.Format("%Y/%m/%d")

	assert.Equal(t, "2079/10/14", res, "%Y/%m/%d did not match")
}

func TestNepaliFormatterFormatymdSlash(t *testing.T) {
	formatter := nepalitime.NewFormatter(globalNepaliTime)
	res := formatter.Format("%y/%m/%d")

	assert.Equal(t, "79/10/14", res, "%y/%m/%d did not match")
}

func TestNepaliFormatterFormatdmYSlash(t *testing.T) {
	formatter := nepalitime.NewFormatter(globalNepaliTime)
	res := formatter.Format("%d/%m/%Y")

	assert.Equal(t, "14/10/2079", res, "%d/%m/%Y did not match")
}

func TestNepaliFormatterFormatdmySlash(t *testing.T) {
	formatter := nepalitime.NewFormatter(globalNepaliTime)
	res := formatter.Format("%d/%m/%y")

	assert.Equal(t, "14/10/79", res, "%d/%m/%y did not match")
}

func TestNepaliFormatterFormatdBYSlash(t *testing.T) {
	formatter := nepalitime.NewFormatter(globalNepaliTime)
	res := formatter.Format("%d/%B/%Y")

	assert.Equal(t, "14/Magh/2079", res, "%d/%B/%Y did not match")
}

func TestNepaliFormatterFormatdBYAComma(t *testing.T) {
	formatter := nepalitime.NewFormatter(globalNepaliTime)
	res := formatter.Format("%d %B, %Y, %A")

	assert.Equal(t, "14 Magh, 2079, Saturday", res, "%d %B, %Y, %A did not match")
}

func TestNepaliFormatterFormatdBYaComma(t *testing.T) {
	formatter := nepalitime.NewFormatter(globalNepaliTime)
	res := formatter.Format("%d %B, %Y, %a")

	assert.Equal(t, "14 Magh, 2079, Sat", res, "%d %B, %Y, %a did not match")
}

func TestNepaliFormatterFormatmdYSlash(t *testing.T) {
	formatter := nepalitime.NewFormatter(globalNepaliTime)
	res := formatter.Format("%m/%d/%Y")

	assert.Equal(t, "10/14/2079", res, "%d/%m/%Y did not match")
}

func TestNepaliFormatterFormatmdySlash(t *testing.T) {
	formatter := nepalitime.NewFormatter(globalNepaliTime)
	res := formatter.Format("%m/%d/%y")

	assert.Equal(t, "10/14/79", res, "%d/%m/%y did not match")
}

func TestNepaliFormatterFormatYmdDash(t *testing.T) {
	formatter := nepalitime.NewFormatter(globalNepaliTime)
	res := formatter.Format("%Y-%m-%d")

	assert.Equal(t, "2079-10-14", res, "%Y-%m-%d did not match")
}

func TestNepaliFormatterFormatymdDash(t *testing.T) {
	formatter := nepalitime.NewFormatter(globalNepaliTime)
	res := formatter.Format("%y-%m-%d")

	assert.Equal(t, "79-10-14", res, "%y-%m-%d did not match")
}

func TestNepaliFormatterFormatdmYDash(t *testing.T) {
	formatter := nepalitime.NewFormatter(globalNepaliTime)
	res := formatter.Format("%d-%m-%Y")

	assert.Equal(t, "14-10-2079", res, "%d-%m-%Y did not match")
}

func TestNepaliFormatterFormatdmyDash(t *testing.T) {
	formatter := nepalitime.NewFormatter(globalNepaliTime)
	res := formatter.Format("%d-%m-%y")

	assert.Equal(t, "14-10-79", res, "%d-%m-%y did not match")
}

func TestNepaliFormatterFormatdBYDash(t *testing.T) {
	formatter := nepalitime.NewFormatter(globalNepaliTime)
	res := formatter.Format("%d-%B-%Y")

	assert.Equal(t, "14-Magh-2079", res, "%d-%B-%Y-%A did not match")
}

func TestNepaliFormatterFormatdBYADash(t *testing.T) {
	formatter := nepalitime.NewFormatter(globalNepaliTime)
	res := formatter.Format("%d-%B-%Y-%A")

	assert.Equal(t, "14-Magh-2079-Saturday", res, "%d-%B-%Y-%A did not match")
}

func TestNepaliFormatterFormatdBYaDash(t *testing.T) {
	formatter := nepalitime.NewFormatter(globalNepaliTime)
	res := formatter.Format("%d-%B-%Y-%a")

	assert.Equal(t, "14-Magh-2079-Sat", res, "%d-%B-%Y-%a did not match")
}

func TestNepaliFormatterFormatmdYDash(t *testing.T) {
	formatter := nepalitime.NewFormatter(globalNepaliTime)
	res := formatter.Format("%m-%d-%Y")

	assert.Equal(t, "10-14-2079", res, "%m-%d-%Y did not match")
}

func TestNepaliFormatterFormatmdyDash(t *testing.T) {
	formatter := nepalitime.NewFormatter(globalNepaliTime)
	res := formatter.Format("%m-%d-%y")

	assert.Equal(t, "10-14-79", res, "%m-%d-%y did not match")
}

func TestNepaliFormatterFormatYmdHMSSlash(t *testing.T) {
	formatter := nepalitime.NewFormatter(globalNepaliTime)
	res := formatter.Format("%Y/%m/%d %H:%M:%S")

	assert.Equal(t, "2079/10/14 16:23:17", res, "%Y/%m/%d %H:%M:%S did not match")
}

func TestNepaliFormatterFormatYmdIMSPSlash(t *testing.T) {
	formatter := nepalitime.NewFormatter(globalNepaliTime)
	res := formatter.Format("%Y/%m/%d %I:%M:%S %p")

	assert.Equal(t, "2079/10/14 04:23:17 PM", res, "%Y/%m/%d %I:%M:%S %p did not match")
}

func TestNepaliFormatterFormatYmdIMSWithoutAMPMSlash(t *testing.T) {
	formatter := nepalitime.NewFormatter(globalNepaliTime)
	res := formatter.Format("%Y/%m/%d %I:%M:%S")

	assert.Equal(t, "2079/10/14 04:23:17", res, "%Y/%m/%d %I:%M:%S did not match")
}

func TestNepaliFormatterFormatYmdIMSfSlash(t *testing.T) {
	formatter := nepalitime.NewFormatter(globalNepaliTimeLeadingZeros)
	res := formatter.Format("%Y/%m/%d %I:%M:%S:%f")

	assert.Equal(t, "2079/01/02 03:04:05:000111", res, "%Y/%m/%d %I:%M:%S:%f did not match")
}

func TestNepaliFormatterFormatYmdSlashNoLeadingZero(t *testing.T) {
	formatter := nepalitime.NewFormatter(globalNepaliTimeLeadingZeros)
	res := formatter.Format("%Y/%-m/%-d")

	assert.Equal(t, "2079/1/2", res, "%Y/%-m/%-d did not match")
}

func TestNepaliFormatterFormatymdSlashNoLeadingZero(t *testing.T) {
	formatter := nepalitime.NewFormatter(globalNepaliTimeLeadingZeros)
	res := formatter.Format("%y/%-m/%-d")

	assert.Equal(t, "79/1/2", res, "%y/%-m/%-d did not match")
}

func TestNepaliFormatterFormatdmYSlashNoLeadingZero(t *testing.T) {
	formatter := nepalitime.NewFormatter(globalNepaliTimeLeadingZeros)
	res := formatter.Format("%-d/%-m/%Y")

	assert.Equal(t, "2/1/2079", res, "%-d/%-m/%Y did not match")
}

func TestNepaliFormatterFormatdmySlashNoLeadingZero(t *testing.T) {
	formatter := nepalitime.NewFormatter(globalNepaliTimeLeadingZeros)
	res := formatter.Format("%-d/%-m/%y")

	assert.Equal(t, "2/1/79", res, "%-d/%-m/%y did not match")
}

func TestNepaliFormatterFormatmdYSlashNoLeadingZero(t *testing.T) {
	formatter := nepalitime.NewFormatter(globalNepaliTimeLeadingZeros)
	res := formatter.Format("%-m/%-d/%Y")

	assert.Equal(t, "1/2/2079", res, "%-d/%-m/%Y did not match")
}

func TestNepaliFormatterFormatmdySlashNoLeadingZero(t *testing.T) {
	formatter := nepalitime.NewFormatter(globalNepaliTimeLeadingZeros)
	res := formatter.Format("%-m/%-d/%y")

	assert.Equal(t, "1/2/79", res, "%-d/%-m/%y did not match")
}

func TestNepaliFormatterFormatYmdIMSfSlashNoLeadingZero(t *testing.T) {
	formatter := nepalitime.NewFormatter(globalNepaliTimeLeadingZeros)
	res := formatter.Format("%Y/%-m/%-d %-I:%-M:%-S:%-f")

	assert.Equal(t, "2079/1/2 3:4:5:111", res, "%Y/%-m/%-d %-I:%-M:%-S:%-f did not match")
}

func TestNepaliFormatterFormatYmdDashNoLeadingZero(t *testing.T) {
	formatter := nepalitime.NewFormatter(globalNepaliTimeLeadingZeros)
	res := formatter.Format("%Y-%-m-%-d")

	assert.Equal(t, "2079-1-2", res, "%Y-%-m-%-d did not match")
}

func TestNepaliFormatterFormatymdDashNoLeadingZero(t *testing.T) {
	formatter := nepalitime.NewFormatter(globalNepaliTimeLeadingZeros)
	res := formatter.Format("%y-%-m-%-d")

	assert.Equal(t, "79-1-2", res, "%y-%-m-%-d did not match")
}

func TestNepaliFormatterFormatdmYDashNoLeadingZero(t *testing.T) {
	formatter := nepalitime.NewFormatter(globalNepaliTimeLeadingZeros)
	res := formatter.Format("%-d-%-m-%Y")

	assert.Equal(t, "2-1-2079", res, "%-d-%-m-%Y did not match")
}

func TestNepaliFormatterFormatdmyDashLeadingZero(t *testing.T) {
	formatter := nepalitime.NewFormatter(globalNepaliTimeLeadingZeros)
	res := formatter.Format("%-d-%-m-%y")

	assert.Equal(t, "2-1-79", res, "%-d-%-m-%y did not match")
}

func TestNepaliFormatterFormatmdYDashNoLeadingZero(t *testing.T) {
	formatter := nepalitime.NewFormatter(globalNepaliTimeLeadingZeros)
	res := formatter.Format("%-m-%-d-%Y")

	assert.Equal(t, "1-2-2079", res, "%-d-%-m-%Y did not match")
}

func TestNepaliFormatterFormatmdyDashNoLeadingZero(t *testing.T) {
	formatter := nepalitime.NewFormatter(globalNepaliTimeLeadingZeros)
	res := formatter.Format("%-m-%-d-%y")

	assert.Equal(t, "1-2-79", res, "%-d-%-m-%y did not match")
}

func TestNepaliFormatterFormatYmdHMSSlashNoLeadingZero(t *testing.T) {
	formatter := nepalitime.NewFormatter(globalNepaliTimeLeadingZeros)
	res := formatter.Format("%Y/%-m/%-d %-H:%-M:%-S")

	assert.Equal(t, "2079/1/2 3:4:5", res, "%Y/%-m/%-d %-H:%-M:%-S did not match")
}

func TestNepaliFormatterFormatYmdIMSfDashNoLeadingZero(t *testing.T) {
	formatter := nepalitime.NewFormatter(globalNepaliTimeLeadingZeros)
	res := formatter.Format("%Y-%-m-%-d %-I-%-M-%-S-%-f")

	assert.Equal(t, "2079-1-2 3-4-5-111", res, "%Y-%-m-%-d %-I-%-M-%-S-%-f did not match")
}

func TestNepaliFormatterFormatSpecialCharacters(t *testing.T) {
	formatter := nepalitime.NewFormatter(globalNepaliTime)
	res := formatter.Format("%Y/%m/%d %I:%M:%S $*#()%%")

	assert.Equal(t, "2079/10/14 04:23:17 $*#()%", res, "%Y/%m/%d %I:%M:%S $*#()%% did not match")
}

func TestNepaliFormatterFormatSpecialCharactersBetweenYearAndMonth(t *testing.T) {
	formatter := nepalitime.NewFormatter(globalNepaliTime)
	res := formatter.Format("%Y/$*#()%%%m/%d")

	assert.Equal(t, "2079/$*#()%10/14", res, "%Y/%m/%d %I:%M:%S:%f did not match")
}

func TestNepaliFormatterFormatReturnAsItIsOnUnknownFormats(t *testing.T) {
	formatter := nepalitime.NewFormatter(globalNepaliTime)
	res := formatter.Format("%k")

	assert.Equal(t, res, "%k", "Unknown format didn't returned as it is")
}

func TestNepaliFormatterFormatReturnAsItIsOnUnknownFormatsNoLeadingZeros(t *testing.T) {
	formatter := nepalitime.NewFormatter(globalNepaliTime)
	res := formatter.Format("%-k")

	assert.Equal(t, res, "%-k", "Unknown format didn't returned as it is")
}

func TestNepaliFormatterFormatWithoutYear(t *testing.T) {
	formatter := nepalitime.NewFormatter(globalNepaliTime)
	res := formatter.Format("%m/%d")

	assert.Equal(t, res, "10/14", "%m/%d did not match")
}

func TestNepaliFormatterFormatHourLessThan10(t *testing.T) {
	date, _ := nepalitime.Date(2079, 11, 4, 9, 10, 11, 123)
	formatter := nepalitime.NewFormatter(date)
	res := formatter.Format("%Y/%m/%d %H::%M::%S::%f")

	assert.Equal(t, "2079/11/04 09::10::11::000123", res, "%Y/%m/%d %H::%M::%S::%f did not match")
}

func TestNepaliFormatterFormatHourEqualTo0(t *testing.T) {
	date, _ := nepalitime.Date(2079, 11, 4, 0, 10, 11, 123)
	formatter := nepalitime.NewFormatter(date)
	res := formatter.Format("%Y/%m/%d %I::%M::%S::%f")

	assert.Equal(t, "2079/11/04 12::10::11::000123", res, "%Y/%m/%d %I::%M::%S::%f did not match")
}

func TestNepaliFormatterFormatHourGreaterThan12(t *testing.T) {
	date, _ := nepalitime.Date(2079, 11, 4, 13, 10, 11, 123)
	formatter := nepalitime.NewFormatter(date)
	res := formatter.Format("%Y/%m/%d %-I::%M::%S::%f")

	assert.Equal(t, "2079/11/04 1::10::11::000123", res, "%Y/%m/%d %-I::%M::%S::%f did not match")
}
