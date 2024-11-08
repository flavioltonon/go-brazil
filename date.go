// Package brazil provides utilities for parsing and formatting dates specific to Brazilian formats.
package brazil

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const (
	// dateDelimiter is the delimiter used in date strings.
	dateDelimiter = "/"
	// dateDelimiterPattern is the regex pattern for date delimiters.
	dateDelimiterPattern = `[ ]{0,1}%s[ ]{0,1}`
	// digitsPattern is the regex pattern for digits.
	digitsPattern = `\d+`
	// DateFormatShort is the short date format (DD/MM/YY).
	DateFormatShort = "02/01/06"
	// DateFormatLong is the long date format (DD/MM/YYYY).
	DateFormatLong = "02/01/2006"
)

var (
	// delimiters is a list of possible date delimiters.
	delimiters = []string{`\/`, `\.`, `\-`, `[ ]`, `,`, `de`}
	// months is a map of month abbreviations to their corresponding month numbers.
	months = monthsOfYear{
		"JAN": 1, "FEV": 2, "MAR": 3, "ABR": 4, "MAI": 5, "JUN": 6, "JUL": 7, "AGO": 8, "SET": 9, "OUT": 10, "NOV": 11, "DEZ": 12,
	}
	// monthPattern is the regex pattern for matching month names or numbers.
	monthPattern = `(?:` + digitsPattern + `|(` + strings.Join(months.getMonths(), "|") + `)[\D!ç]{0,7})`
	// datePattern is the regex pattern for matching dates.
	datePattern = `(?i)` + digitsPattern + dateDelimiterPattern + monthPattern + dateDelimiterPattern + digitsPattern
)

type monthsOfYear map[string]int8

// getMonths returns the months of the year as a slice of strings.
func (m monthsOfYear) getMonths() []string {
	keys := make([]string, 0, len(m))
	for name := range m {
		keys = append(keys, name)
	}
	return keys
}

// date struct represents a date value.
type date struct {
	value time.Time
}

// Time returns the date as a time.Time.
func (d date) Time() time.Time {
	return d.value
}

// String returns the date as a string with a DateFormatLong format.
func (d date) String() string {
	return d.value.Format(DateFormatLong)
}

// ParseDate parses a date from a string and returns a date struct.
// Here are some possibilities for usage:
// - "5 de Abril de 1999"
// - "05 Abril 1999"
// - "5-4-1999"
// - "05.ABR.1999"
// - "5,4,1999"
//
// The date can be separated by a space, a dot, a hyphen, a comma, or the word "de".
func ParseDate(value string) (date, error) {
	var (
		matches []string
		err     error
		time    *time.Time
	)

	// Create a pattern for each delimiter and find all matches.
	for _, delimiter := range delimiters {
		pattern := fmt.Sprintf(datePattern, delimiter, delimiter)
		matches = append(matches, regexp.MustCompile(pattern).FindAllString(value, -1)...)
	}

	// Try to get the time from the matches.
	for _, match := range matches {
		time, err = getTime(match)
		// If there has been an error, return it.
		if err != nil {
			return date{}, err
		}
	}

	// If the time is nil, return an error.
	if time == nil {
		return date{}, ErrInvalidDate
	}

	// Return the date.
	return date{value: *time}, nil
}

// getTime parses a date string and returns a time.Time pointer.
func getTime(value string) (*time.Time, error) {
	// Replace all non-word characters with the date delimiter.
	r := regexp.MustCompile(`(?:\W|(`+strings.Join(delimiters, "|")+`)\W)+`).ReplaceAllString(value, dateDelimiter)
	// Split the string by the date delimiter.
	arr := strings.Split(r, dateDelimiter)

	// If the array has less than 3 elements, return an error.
	if len(arr) < 3 {
		return nil, ErrInvalidDate
	}

	// Set the day, month, and year.
	day := arr[0]
	month := normalizeMonth(arr[1])
	year := arr[len(arr)-1:][0]

	// Set the default date format and year digits.
	dateFormat := DateFormatLong
	yearDigits := "%04d"

	// If the year has 2 digits, change the date format and the year digits.
	if len(year) == 2 {
		dateFormat = DateFormatShort
		yearDigits = "%02d"
	}

	// Parse the date with the format obtained.
	layout := fmt.Sprintf("%02d/%02d/"+yearDigits, parseToint(day), parseToint(month), parseToint(year))

	date, err := time.Parse(dateFormat, layout)
	if err != nil {
		return nil, fmt.Errorf("parsing time with layout %q: %w", layout, ErrInvalidDate)
	}

	// Return the date.
	return &date, nil
}

// normalizeMonth converts a month name to its corresponding month number.
func normalizeMonth(text string) string {
	// If the text is empty or has less than 3 characters, return it.
	if len(text) < 3 {
		return text
	}

	// Get the first 3 characters of the text and convert them to uppercase.
	month := strings.ToUpper(text)[0:3]

	// Check if the month is in the map, if not return an empty string.
	value, ok := months[month]
	if !ok {
		return ""
	}

	// Return the month number with 2 digits.
	return fmt.Sprintf("%02d", value)
}

// parseToint converts a string to an integer.
func parseToint(value string) int {
	if value == "" {
		return 0
	}
	i, err := strconv.Atoi(value)
	if err != nil {
		return 0
	}
	return i
}
