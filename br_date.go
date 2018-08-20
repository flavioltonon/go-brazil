package brazil

import (
	"time"
)

func ParseDate(number string) brDate {
	date, err := time.Parse("02/01/2006", number)
	if err != nil {
		return brDate{
			Date: date,
			Err:  err,
		}
	}
	return brDate{
		Date: date,
		Err:  nil,
	}
}

func (date brDate) IsPast() bool {
	t := date.Date
	now := time.Now()
	return t.Year() < now.Year() ||
		t.Year() == now.Year() && int(t.Month()) < int(now.Month()) ||
		t.Year() == now.Year() && int(t.Month()) == int(now.Month()) && t.Day() < now.Day()
}

func (date *brDate) IsToday() bool {
	t := date.Date
	now := time.Now()
	return t.Year() == now.Year() && int(t.Month()) == int(now.Month()) && t.Day() == now.Day()
}

func (date *brDate) IsFuture() bool {
	t := date.Date
	now := time.Now()
	return t.Year() > now.Year() ||
		t.Year() == now.Year() && int(t.Month()) > int(now.Month()) ||
		t.Year() == now.Year() && int(t.Month()) == int(now.Month()) && t.Day() > now.Day()
}
