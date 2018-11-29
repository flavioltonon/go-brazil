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

func ParseDate(d time.Time) date {
	return date{
		date: d.Truncate(24 * time.Hour),
	}
}

func (d *date) Date() string {
	return d.date.Format("02/01/2006")
}

func (d *date) IsValid() bool {
	if !d.notNull {
		d.validation = validation{
			valid:  false,
			reason: errNullDate,
		}
		return false
	}

	if d.date.IsZero() {
		d.validation = validation{
			valid:  false,
			reason: errIncorrectFormatDate,
		}
		return false
	}

	d.validation = validation{
		valid:  true,
		reason: nil,
	}
	return true
}

func (d *date) Errors() []error {
	if d.validation.valid {
		return nil
	}
	return []error{
		d.validation.reason,
	}
}

func RandomDate(minYear int32, maxYear int32) time.Time {
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

	return date
}

func IsLeapYear(year int) bool {
	return (year%4 == 0 && year%100 != 0) ||
		year%400 == 0
}

func (d date) IsPast() bool {
	_, offset := d.date.Zone()
	date := d.date.Add(time.Duration(offset) * time.Second).UTC().Truncate(24 * time.Hour)
	if date.IsZero() {
		return false
	}

	today := time.Now().Truncate(24 * time.Hour)
	return int(today.Sub(date).Hours()/24) > 0
}

func (d date) IsOlderThan(ref date) bool {
	return ref.date.Sub(d.date).Hours() > 0
}

func (d *date) IsToday() bool {
	_, offset := d.date.Zone()
	date := d.date.Add(time.Duration(offset) * time.Second).UTC().Truncate(24 * time.Hour)
	if date.IsZero() {
		return false
	}

	today := time.Now().Truncate(24 * time.Hour)
	return int((today.Sub(date).Hours())/24) == 0
}

func (d *date) IsFuture() bool {
	_, offset := d.date.Zone()
	date := d.date.Add(time.Duration(offset) * time.Second).UTC().Truncate(24 * time.Hour)
	if date.IsZero() {
		return false
	}

	today := time.Now().UTC().Truncate(24 * time.Hour)
	return int(today.Sub(date).Hours()/24) < 0
}
