package nepalitime

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoDigitNumberWhenNumberIsLessThan0AndGreaterThan0(t *testing.T) {
	number := 1
	want := "01"

	got := twoDigitNumber(number)

	assert.Equal(t, want, got, "want and got should be equal")
}

func TestTwoDigitNumberWhenNumberIsEqualTo10(t *testing.T) {
	number := 10
	want := "10"

	got := twoDigitNumber(number)

	assert.Equal(t, want, got, "the value returned should be 10")
}

func TestTwoDigitNumberGreaterThan10(t *testing.T) {
	number := 11
	want := "11"

	got := twoDigitNumber(number)

	assert.Equal(t, want, got, "the value returned should be 11")
}

func TestTwoDigitWhenNumberEqualTo0(t *testing.T) {
	number := 0
	want := "00"

	got := twoDigitNumber(number)

	assert.Equal(t, want, got, "the value returned should be 00")
}

func TestTwoDigitWhenNumberLessThan0(t *testing.T) {
	number := -1
	want := "-1"

	got := twoDigitNumber(number)

	assert.Equal(t, want, got, "the value returned should be -1")
}
