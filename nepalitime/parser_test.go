package nepalitime_test

import (
	"fmt"
	"testing"

	"github.com/opensource-nepal/go-nepali/nepalitime"
	"github.com/stretchr/testify/assert"
)

// TODO: currently we only provide a read-only
// value of NepaliTime struct so, to test only
// the String() method is to be tested
// these tests might require changes once manipulation
// of the NepaliTime struct is provided

func TestParseWithYmdSlashFormat(t *testing.T) {
	// 2079 magh 14
	datetimeStr := "2079/10/14"
	format := "%Y/%m/%d"

	expected := "2079-10-14 00:00:00"

	got, err := nepalitime.Parse(datetimeStr, format)

	assert.Equal(t, expected, got.String(), fmt.Sprintf("expected: %s - got: %s", expected, got))
	assert.Nil(t, err, "error should be nil")
}

func TestParserWithymdSlashFormat(t *testing.T) {
	// 2079 magh 14
	datetimeStr := "79/10/14"
	format := "%y/%m/%d"

	expected := "2079-10-14 00:00:00"

	got, err := nepalitime.Parse(datetimeStr, format)

	assert.Equal(t, expected, got.String(), fmt.Sprintf("expected: %s - got: %s", expected, got))
	assert.Nil(t, err, "error should be nil")
}

func TestParserWithdmYSlashForm(t *testing.T) {
	// 2079 magh 14
	datetimeStr := "14/10/2079"
	format := "%d/%m/%Y"

	expected := "2079-10-14 00:00:00"

	got, err := nepalitime.Parse(datetimeStr, format)

	assert.Equal(t, expected, got.String(), fmt.Sprintf("expected: %s - got: %s", expected, got))
	assert.Nil(t, err, "error should be nil")
}

func TestParserWithdmySlashForm(t *testing.T) {
	// 2079 magh 14
	datetimeStr := "14/10/79"
	format := "%d/%m/%y"

	expected := "2079-10-14 00:00:00"

	got, err := nepalitime.Parse(datetimeStr, format)

	assert.Equal(t, expected, got.String(), fmt.Sprintf("expected: %s - got: %s", expected, got))
	assert.Nil(t, err, "error should be nil")
}

func TestParserWithmdYSlashForm(t *testing.T) {
	// 2079 magh 14
	datetimeStr := "10/14/2079"
	format := "%m/%d/%Y"

	expected := "2079-10-14 00:00:00"

	got, err := nepalitime.Parse(datetimeStr, format)

	assert.Equal(t, expected, got.String(), fmt.Sprintf("expected: %s - got: %s", expected, got))
	assert.Nil(t, err, "error should be nil")
}

func TestParserWithmdySlashForm(t *testing.T) {
	// 2079 magh 14
	datetimeStr := "10/14/79"
	format := "%m/%d/%y"

	expected := "2079-10-14 00:00:00"

	got, err := nepalitime.Parse(datetimeStr, format)

	assert.Equal(t, expected, got.String(), fmt.Sprintf("expected: %s - got: %s", expected, got))
	assert.Nil(t, err, "error should be nil")
}

func TestParseWithYmdDashFormat(t *testing.T) {
	// 2079 magh 14
	datetimeStr := "2079-10-14"
	format := "%Y-%m-%d"

	expected := "2079-10-14 00:00:00"

	got, err := nepalitime.Parse(datetimeStr, format)

	assert.Equal(t, expected, got.String(), fmt.Sprintf("expected: %s - got: %s", expected, got))
	assert.Nil(t, err, "error should be nil")
}

func TestParserWithymdDashFormat(t *testing.T) {
	// 2079 magh 14
	datetimeStr := "79-10-14"
	format := "%y-%m-%d"

	expected := "2079-10-14 00:00:00"

	got, err := nepalitime.Parse(datetimeStr, format)

	assert.Equal(t, expected, got.String(), fmt.Sprintf("expected: %s - got: %s", expected, got))
	assert.Nil(t, err, "error should be nil")
}

func TestParserWithdmYDashForm(t *testing.T) {
	// 2079 magh 14
	datetimeStr := "14-10-2079"
	format := "%d-%m-%Y"

	expected := "2079-10-14 00:00:00"

	got, err := nepalitime.Parse(datetimeStr, format)

	assert.Equal(t, expected, got.String(), fmt.Sprintf("expected: %s - got: %s", expected, got))
	assert.Nil(t, err, "error should be nil")
}

func TestParserWithdmyDashForm(t *testing.T) {
	// 2079 magh 14
	datetimeStr := "14-10-79"
	format := "%d-%m-%y"

	expected := "2079-10-14 00:00:00"

	got, err := nepalitime.Parse(datetimeStr, format)

	assert.Equal(t, expected, got.String(), fmt.Sprintf("expected: %s - got: %s", expected, got))
	assert.Nil(t, err, "error should be nil")
}

func TestParserWithmdYDashForm(t *testing.T) {
	// 2079 magh 14
	datetimeStr := "10-14-2079"
	format := "%m-%d-%Y"

	expected := "2079-10-14 00:00:00"

	got, err := nepalitime.Parse(datetimeStr, format)

	assert.Equal(t, expected, got.String(), fmt.Sprintf("expected: %s - got: %s", expected, got))
	assert.Nil(t, err, "error should be nil")
}

func TestParserWithmdyDashForm(t *testing.T) {
	// 2079 magh 14
	datetimeStr := "10-14-79"
	format := "%m-%d-%y"

	expected := "2079-10-14 00:00:00"

	got, err := nepalitime.Parse(datetimeStr, format)

	assert.Equal(t, expected, got.String(), fmt.Sprintf("expected: %s - got: %s", expected, got))
	assert.Nil(t, err, "error should be nil")
}

func TestParserWithYmdHMSForm(t *testing.T) {
	// 2079 magh 14 21:27:30
	datetimeStr := "2079/10/14 21:27:30"
	format := "%Y/%m/%d %H:%M:%S"

	expected := "2079-10-14 21:27:30"

	got, err := nepalitime.Parse(datetimeStr, format)

	assert.Equal(t, expected, got.String(), fmt.Sprintf("expected: %s - got: %s", expected, got))
	assert.Nil(t, err, "error should be nil")
}

func TestParserWithYmdIMSForm(t *testing.T) {
	// 2079 magh 14 21:27:30
	datetimeStr := "2079/10/14 09:27:30 PM"
	format := "%Y/%m/%d %I:%M:%S %p"

	expected := "2079-10-14 21:27:30"

	got, err := nepalitime.Parse(datetimeStr, format)

	assert.Equal(t, expected, got.String(), fmt.Sprintf("expected: %s - got: %s", expected, got))
	assert.Nil(t, err, "error should be nil")
}

func TestParserWithYmdIMSWithoutAmPm(t *testing.T) {
	datetimeStr := "2079/10/14 09:27:30"
	format := "%Y/%m/%d %I:%M:%S"

	expected := "2079-10-14 09:27:30"

	got, err := nepalitime.Parse(datetimeStr, format)

	assert.Equal(t, expected, got.String(), fmt.Sprintf("expected: %s - got: %s", expected, got))
	assert.Nil(t, err, "error should be nil")
}

func TestParserWithYmdHMSf(t *testing.T) {
	datetimeStr := "2079/10/14 09:27:30:1"
	format := "%Y/%m/%d %H:%M:%S:%f"

	expected := "2079-10-14 09:27:30"

	got, err := nepalitime.Parse(datetimeStr, format)

	assert.Equal(t, expected, got.String(), fmt.Sprintf("expected: %s - got: %s", expected, got))
	assert.Equal(t, got.Nanosecond(), 100000)
	assert.Nil(t, err, "error should be nil")
}

func TestParseWithSpecialCharacters(t *testing.T) {
	// 2079 magh 14 21:27:30
	datetimeStr := "2079/10/14 21:27:30 $*#()%"
	format := "%Y/%m/%d %H:%M:%S $*#()%%"

	expected := "2079-10-14 21:27:30"

	got, err := nepalitime.Parse(datetimeStr, format)

	assert.Equal(t, expected, got.String(), fmt.Sprintf("expected: %s - got: %s", expected, got))
	assert.Nil(t, err, "error should be nil")
}

func TestParseWithSpecialCharactersBetweenYearAndMonth(t *testing.T) {
	// 2079 magh 14 21:27:30
	datetimeStr := "2079/$*#()%10/14"
	format := "%Y/$*#()%%%m/%d"
	expected := "2079-10-14 00:00:00"

	got, err := nepalitime.Parse(datetimeStr, format)

	assert.Equal(t, expected, got.String(), fmt.Sprintf("expected: %s - got: %s", expected, got))
	assert.Nil(t, err, "error should be nil")
}

func TestParseWithErrorNousFormats(t *testing.T) {
	datetimeStr := "2079/10/14"
	format := "%y-%m-%d"

	got, err := nepalitime.Parse(datetimeStr, format)

	assert.Nil(t, got, "NepaliTime object should be nil")
	assert.EqualError(t, err, "datetime string did not match with given format", "error message did not match")
}

func TestParseWithRandomFormats(t *testing.T) {
	datetimeStr := "2079/10/14"
	format := "%k"

	got, err := nepalitime.Parse(datetimeStr, format)

	assert.Nil(t, got, "NepaliTime object should be nil")
	assert.EqualError(t, err, "the format '%k' isn't supported", "error message did not match")
}

func TestParseForInvalidYear(t *testing.T) {
	datetimeStr := "0000"
	format := "%Y"

	got, err := nepalitime.Parse(datetimeStr, format)

	assert.Nil(t, got, "NepaliTime object should be nil")
	assert.EqualError(t, err, "date is out of range", "error message did not match")
}

func TestParseForValidYear(t *testing.T) {
	datetimeStr := "2079"
	format := "%Y"
	got, err := nepalitime.Parse(datetimeStr, format)
	assert.Nil(t, err, "error should be nil")
	assert.Equal(t, got.String(), "2079-01-01 00:00:00")
}

func TestParseWithoutYear(t *testing.T) {
	datetimeStr := "01-02"
	format := "%m-%d"

	got, err := nepalitime.Parse(datetimeStr, format)

	assert.Nil(t, got, "NepaliTime object should be nil")
	assert.NotNil(t, err)
}

func TestParseWithHyphenFormat(t *testing.T) {
	datetimeStr := "2079-1-2"
	format := "%Y-%-m-%-d"
	got, err := nepalitime.Parse(datetimeStr, format)
	assert.Nil(t, err, "error should be nil")
	assert.Equal(t, got.String(), "2079-01-02 00:00:00")
}

func TestParseWithValidRegexFormat(t *testing.T) {
	datetimeStr := "2079"
	format := `(?P<Y>\d\d\d\d)`
	got, err := nepalitime.Parse(datetimeStr, format)
	assert.Nil(t, got, "NepaliTime object should be nil")
	assert.EqualError(t, err, "datetime string did not match with given format", "error message did not match")
}

func TestParseFor12AM(t *testing.T) {
	datetimeStr := "2079/10/14 12:00:00 AM"
	format := "%Y/%m/%d %I:%M:%S %p"

	got, err := nepalitime.Parse(datetimeStr, format)

	assert.Nil(t, err, "error should be nil")
	assert.Equal(t, got.String(), "2079-10-14 00:00:00")
}
