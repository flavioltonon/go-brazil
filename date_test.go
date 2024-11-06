package brazil

import (
	"fmt"
	"testing"
)

func Test_normalizeMonth(t *testing.T) {
	var months = []string{"JANEIRO", "FEVEREIRO", "MARÇO", "ABRIL", "MAIO", "JUNHO", "JULHO", "AGOSTO", "SETEMBRO", "OUTUBRO", "NOVEMBRO", "DEZEMBRO"}
	t.Run("should normalize the month name to number", func(t *testing.T) {
		for idx, name := range months {
			want := fmt.Sprintf("%02d", idx+1)
			got := normalizeMonth(name)
			if got != want {
				t.Errorf("Expected %s, got %#v", want, got)
			}
		}
	})
	t.Run("should return an empty string when the month name is not valid", func(t *testing.T) {
		want := ""
		got := normalizeMonth("INVALID")
		if got != want {
			t.Errorf("Expected %s, got %#v", want, got)
		}
	})
}

func TestParseDate(t *testing.T) {
	t.Run("should return an error when the date is not valid", func(t *testing.T) {
		_, err := ParseDate("01 13 2020")
		if err == nil {
			t.Errorf("Expected an error, got nil")
		}
	})
	t.Run("should return a valid date", func(t *testing.T) {
		want := "01/06/1920"
		for _, value := range []string{
			"01 JUN/JUN 1920",
			"1 JUN/JUN 1920",
			"1.JUN/JUN.1920",
			"01/JUNHO/1920",
			"1/JUNHO/1920",
			"01/JUN/1920",
			"1 JUNHO 1920",
			"01-JUNHO-1920",
			"1 JUN 1920",
			"01/06/1920",
			"01-06-1920",
			"01.06.1920",
			"01 06 1920",
			"1 6 1920",
		} {
			got, err := ParseDate(value)
			if err != nil {
				t.Errorf("Expected nil, got %s", err)
			}
			if got.String() != want || got.Time().Format(DateFormatLong) != want {
				t.Errorf("Expected %s, got %#v", want, got)
			}
		}
	})
}
