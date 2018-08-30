package brazil

import (
	"math/rand"
	"strconv"
	"time"
)

var days = map[string]int32{
	"1":  31,
	"2":  28,
	"2b": 29,
	"3":  31,
	"4":  30,
	"5":  31,
	"6":  30,
	"7":  31,
	"8":  31,
	"9":  30,
	"10": 31,
	"11": 30,
	"12": 31,
}

func ParseDate(number string) brDate {
	date, err := time.Parse("02/01/2006", number)
	return brDate{
		Date: date,
		Err:  err,
	}
}

func GenerateRandomDate(minYear int32, maxYear int32) string {
	var maxDay int32
	var day, month, year int

	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	// Generate random month
	month = int(r.Int31n(11) + 1)

	// Generate random year
	year = int(r.Int31n(maxYear-minYear) + minYear)

	// Generate random day
	if month == 2 && IsLeapYear(year) {
		maxDay = days["2b"]
	} else {
		maxDay = days[strconv.Itoa(month)]
	}
	day = int(r.Int31n(maxDay-1) + 1)

	date := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)

	return date.Format("02/01/2006")
}

func IsLeapYear(year int) bool {
	return (year%4 == 0 && year%100 != 0) ||
		year%400 == 0
}

func (date brDate) IsPast() bool {
	_, offset := date.Date.Zone()
	d := date.Date.Add(time.Duration(offset) * time.Second).UTC().Truncate(24 * time.Hour)
	if d.IsZero() {
		return false
	}

	today := time.Now().Truncate(24 * time.Hour)
	return int(today.Sub(d).Hours()/24) > 0
}

func (date *brDate) IsToday() bool {
	_, offset := date.Date.Zone()
	d := date.Date.Add(time.Duration(offset) * time.Second).UTC().Truncate(24 * time.Hour)
	if d.IsZero() {
		return false
	}

	today := time.Now().Truncate(24 * time.Hour)
	return int((today.Sub(d).Hours())/24) == 0
}

func (date *brDate) IsFuture() bool {
	_, offset := date.Date.Zone()
	d := date.Date.Add(time.Duration(offset) * time.Second).UTC().Truncate(24 * time.Hour)
	if d.IsZero() {
		return false
	}

	today := time.Now().UTC().Truncate(24 * time.Hour)
	return int(today.Sub(d).Hours()/24) < 0
}
