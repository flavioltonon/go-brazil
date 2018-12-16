package brazil

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type date struct {
	year  year
	month month
	day   day
}

type year string
type month string
type day string

func (d date) Date() string {
	return d.Day() + "/" + d.Month() + "/" + d.Year()
}

func (d date) Time() time.Time {
	var (
		day, _      = strconv.Atoi(string(d.day))
		month, _    = strconv.Atoi(string(d.month))
		year, _     = strconv.Atoi(string(d.year))
		location, _ = time.LoadLocation("Local")
	)

	return time.Date(
		year,
		time.Month(month),
		day,
		0,
		0,
		0,
		0,
		location,
	)
}

func (d date) Year() string {
	return string(d.year)
}

func (d date) Month() string {
	return string(d.month)
}

func (d date) Day() string {
	return string(d.day)
}

func ParseDate(d string) (date, error) {
	var _, err = time.Parse("02/01/2006", d)
	if err != nil {
		return date{}, errIncorrectFormatDate
	}

	return date{
		year:  year(strings.Split(d, "/")[2]),
		month: month(strings.Split(d, "/")[1]),
		day:   day(strings.Split(d, "/")[0]),
	}, nil
}

func RandomDate(minYear int32, maxYear int32) (string, error) {
	var (
		source      = rand.NewSource(time.Now().UnixNano())
		maxDay      int32
		location, _ = time.LoadLocation("Local")
	)

	if minYear > maxYear || minYear < 0 {
		return "", errInvalidYearLimits
	}

	r := rand.New(source)

	// Generate random month
	month := int(r.Int31n(11) + 1)

	// Generate random year
	year := int(r.Int31n(maxYear-minYear) + minYear)

	// Generate random day
	maxDay = days[strconv.Itoa(month)]
	if month == 2 && IsLeapYear(strconv.Itoa(year)) {
		maxDay = days["2b"]
	}
	day := int(r.Int31n(maxDay-1) + 1)

	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, location).Format("02/01/2006"), nil
}

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

func IsLeapYear(year string) bool {
	var yInt, _ = strconv.Atoi(year)
	return yInt%4 == 0 && (yInt%100 != 0 || yInt%400 == 0)
}

func (d date) IsPast() bool {
	date := d.Time()
	if date.IsZero() {
		return false
	}

	today := time.Now().Truncate(24 * time.Hour)
	return int(today.Sub(date).Hours()/24) > 0
}

func (d date) IsOlderThan(ref date) bool {
	return ref.Time().Truncate(24*time.Hour).Sub(d.Time()).Hours() > 0
}

func (d date) IsToday() bool {
	date := d.Time()
	if date.IsZero() {
		return false
	}

	today := time.Now().Truncate(24 * time.Hour)
	return int((today.Sub(date).Hours())/24) == 0
}

func (d date) IsFuture() bool {
	date := d.Time()
	if date.IsZero() {
		return false
	}

	today := time.Now().Truncate(24 * time.Hour)
	return int(today.Sub(date).Hours()/24) < 0
}
