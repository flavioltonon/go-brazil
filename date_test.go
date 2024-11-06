package brazil

import (
	"errors"
	"fmt"
	"testing"
)

func Test_normalizeMonth(t *testing.T) {
	t.Run("must normalize the month name to the number", func(t *testing.T) {
		tableTests := []struct {
			name   string
			values []string
		}{
			{
				name:   "short",
				values: []string{"JAN", "FEV", "MAR", "ABR", "MAI", "JUN", "JUL", "AGO", "SET", "OUT", "NOV", "DEZ"},
			},
			{
				name:   "full",
				values: []string{"JANEIRO", "FEVEREIRO", "MARÇO", "ABRIL", "MAIO", "JUNHO", "JULHO", "AGOSTO", "SETEMBRO", "OUTUBRO", "NOVEMBRO", "DEZEMBRO"},
			},
		}
		for _, tt := range tableTests {
			for idx, name := range tt.values {
				want := fmt.Sprintf("%02d", idx+1)
				got := normalizeMonth(name)
				if got != want {
					t.Errorf("Expected %s, got %#v", want, got)
				}
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
		tableTests := []struct {
			name  string
			value string
			want  error
		}{
			{
				value: "01 13 2020",
				want:  ErrInvalidDate,
			},
			{
				value: "01 AGO",
				want:  ErrInvalidDate,
			},
		}
		for _, tt := range tableTests {
			_, err := ParseDate(tt.value)
			if err == nil {
				t.Errorf("Expected an error, got nil")
			}
			if !errors.Is(err, tt.want) {
				t.Errorf("Expected %s, got %#v", tt.want, err)
			}
		}
	})
	t.Run("should return a valid date", func(t *testing.T) {
		for value, want := range map[string]string{
			"27 de AG0STO de 1994": "27/08/1994",
			"9 JUL/JUL 1932":       "09/07/1932",
			"19/ABRIL/1943":        "19/04/1943",
			"15.NOVEMBRO.1889":     "15/11/1889",
			"11-SET-01":            "11/09/2001",
			"1 JU1HO 2024":         "01/07/2024",
			"01-JUNHO-1920":        "01/06/1920",
			"1/1/1992":             "01/01/1992",
			"1-06-1920":            "01/06/1920",
			"21.09.2012":           "21/09/2012",
			"1,5,1889":             "01/05/1889",
			"20 Nov 1695":          "20/11/1695",
			"9 MAI/MAY 1988":       "09/05/1988",
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
