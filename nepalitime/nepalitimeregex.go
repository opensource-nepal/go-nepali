// Package nepalitime contains functionality related to Nepali time pattern matching
package nepalitime

import (
	"fmt"
	"regexp"
	"strings"
)

type nepaliTimeRegex struct {
	PatternMap map[string]string
}

// nepaliTimeRe constructor
func newNepaliTimeRegex() *nepaliTimeRegex {
	obj := new(nepaliTimeRegex)
	obj.PatternMap = map[string]string{
		"d":  `(?P<d>3[0-2]|[1-2]\d|0[1-9]|[1-9]| [1-9])`,
		"-d": `(?P<d>3[0-2]|[1-2]\d|0[1-9]|[1-9]| [1-9])`, // same as "d"
		"f":  `(?P<f>[0-9]{1,6})`,
		"H":  `(?P<H>2[0-3]|[0-1]\d|\d)`,
		"-H": `(?P<H>2[0-3]|[0-1]\d|\d)`,
		"I":  `(?P<I>1[0-2]|0[1-9]|[1-9])`,
		"-I": `(?P<I>1[0-2]|0[1-9]|[1-9])`,
		"G":  `(?P<G>\d\d\d\d)`,
		"j":  `(?P<j>36[0-6]|3[0-5]\d|[1-2]\d\d|0[1-9]\d|00[1-9]|[1-9]\d|0[1-9]|[1-9])`,
		"m":  `(?P<m>1[0-2]|0[1-9]|[1-9])`,
		"-m": `(?P<m>1[0-2]|0[1-9]|[1-9])`, // same as "m"
		"M":  `(?P<M>[0-5]\d|\d)`,
		"-M": `(?P<M>[0-5]\d|\d)`, // same as "M"
		"S":  `(?P<S>6[0-1]|[0-5]\d|\d)`,
		"-S": `(?P<S>6[0-1]|[0-5]\d|\d)`, // same as "S"
		"w":  `(?P<w>[0-6])`,

		"y": `(?P<y>\d\d)`,
		"Y": `(?P<Y>\d\d\d\d)`,
		"z": `(?P<z>[+-]\d\d:?[0-5]\d(:?[0-5]\d(\.\d{1,6})?)?|(?-i:Z))`,
		"B": `(?P<B>Baisakh|Jestha|Ashadh|Shrawan|Bhadra|Ashwin|Kartik|Mangsir|Poush|Magh|Falgun|Chaitra")`,
		"A": `(?P<B>Sunday|Monday|Tuesday|Wednesday|Thursday|Friday|Saturday)`,
		// "b": obj.__seqToRE(EnglishChar.months, "b"),
		// "p": obj.__seqToRE(("AM", "PM",), "p"),
		// TODO: implement for the above commented directives
		"p": "(?P<p>AM|PM)",

		"%": "%",
	}

	return obj
}

// Handles conversion from format directives to regexes
func (obj *nepaliTimeRegex) pattern(format string) (string, error) {
	processedFormat := ""
	regexChars := regexp.MustCompile(`([\.^$*+?\(\){}\[\]|])`)
	format = regexChars.ReplaceAllString(format, `\$1`)
	whitespaceReplacement := regexp.MustCompile(`\s+`)
	format = whitespaceReplacement.ReplaceAllString(format, `\s+`)

	for {
		index := strings.Index(format, "%")
		// index = -1 means the sub string does not exist
		// in the string
		if index == -1 {
			break
		}

		directiveIndex := index + 1
		indexIncrement := 1

		if string(format[directiveIndex]) == "-" {
			indexIncrement = 2
		}

		directiveToCheck := string(format[directiveIndex : directiveIndex+indexIncrement])

		if val, ok := obj.PatternMap[directiveToCheck]; ok {
			processedFormat = fmt.Sprintf("%s%s%s", processedFormat, string(format[:directiveIndex-1]), val)
			format = string(format[directiveIndex+indexIncrement:])
		} else {
			return "", fmt.Errorf("the format '%%%s' isn't supported", directiveToCheck)
		}
	}

	tester := fmt.Sprintf("^%s%s$", processedFormat, format)

	return tester, nil
}

// handles regex compilation for format string
func (obj *nepaliTimeRegex) compile(format string) (*regexp.Regexp, error) {
	processedFormat, err := obj.pattern(format)

	if err != nil {
		return nil, err
	}

	// (?i) is for ignoring the case
	reg, err := regexp.Compile("(?i)" + processedFormat)

	if err != nil {
		return nil, err
	}

	return reg, nil
}
