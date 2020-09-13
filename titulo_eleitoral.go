package brazil

import (
	"math/rand"
	"regexp"
	"strconv"
	"time"
)

// Título de eleitor struct
type tituloEleitoral struct {
	number tituloEleitoralNumber
	valid  bool
}

func (t tituloEleitoral) Number(mask bool) string {
	if t.valid && mask {
		return string(t.number[:4]) + " " + string(t.number[4:8]) + " " + string(t.number[8:])
	}
	return string(t.number)
}

type tituloEleitoralNumber string

func ParseTituloEleitoral(number string) (tituloEleitoral, error) {
	number = regexp.MustCompile(`[^0-9]`).ReplaceAllString(number, "")
	switch len(number) {
	case 11:
		number = "0" + number
	case 12:
		// OK
	default:
		return tituloEleitoral{}, ErrIncorrectLenghtTituloEleitoralNumber
	}

	tituloNumber := tituloEleitoralNumber(number)

	if tituloNumber.isFalsePositive() {
		return tituloEleitoral{}, ErrInvalidTituloEleitoralNumber
	}

	if !tituloNumber.hasValidFirstDigit() {
		return tituloEleitoral{}, ErrInvalidTituloEleitoralNumber
	}

	if !tituloNumber.hasValidSecondDigit() {
		return tituloEleitoral{}, ErrInvalidTituloEleitoralNumber
	}

	return tituloEleitoral{
		number: tituloNumber,
		valid:  true,
	}, nil
}

func RandomTituloEleitoralNumber(mask bool) string {
	var (
		source = rand.NewSource(time.Now().UnixNano())
		sum    int
	)

	r := rand.New(source)

	tNumber := int(r.Int31n(89999999) + 10000000)
	tString := strconv.Itoa(tNumber)

	stateCode := int(r.Int31n(27) + 1)
	stateCodeString := strconv.Itoa(stateCode)
	if stateCode < 10 {
		stateCodeString = "0" + stateCodeString
	}

	digitlessTitulo := tString + stateCodeString

	// First digit
	for i := 0; i < 8; i++ {
		number, _ := strconv.Atoi(string(digitlessTitulo[i]))
		sum += number * (i + 2)
	}

	firstDigit := sum % 11
	switch firstDigit {
	case 0:
		if stateCodeString == "01" || stateCodeString == "02" {
			firstDigit = 1
		}
	case 10:
		firstDigit = 0
	}

	// Second digit
	sum = 0
	for i := 8; i < 10; i++ {
		number, _ := strconv.Atoi(string(digitlessTitulo[i]))
		sum += number * (i - 1)
	}
	sum = sum + firstDigit*9
	secondDigit := sum % 11
	switch secondDigit {
	case 0:
		if stateCodeString == "01" || stateCodeString == "02" {
			secondDigit = 1
		}
	case 10:
		secondDigit = 0
	}

	if mask {
		return digitlessTitulo[:4] + " " + digitlessTitulo[4:8] + " " + digitlessTitulo[8:] + strconv.Itoa(firstDigit) + strconv.Itoa(secondDigit)
	}
	return digitlessTitulo + strconv.Itoa(firstDigit) + strconv.Itoa(secondDigit)
}

func (t tituloEleitoralNumber) isFalsePositive() bool {
	if string(t) == "000000000000" {
		return true
	}
	return false

}

func (t tituloEleitoralNumber) hasValidFirstDigit() bool {
	var sum int

	for i := 0; i < 8; i++ {
		tituloEleitoralDigit, _ := strconv.Atoi(string(t[i]))
		sum += tituloEleitoralDigit * (i + 2)
	}

	digit := sum % 11
	if digit != 0 {
		return string(t[10]) == strconv.Itoa(digit%10)
	}

	stateCode := string(t[8:10])
	if stateCode == "01" || stateCode == "02" {
		return string(t[10]) == strconv.Itoa(1)
	}

	return string(t[10]) == strconv.Itoa(digit)
}

func (t tituloEleitoralNumber) hasValidSecondDigit() bool {
	var sum int

	for i := 8; i < 11; i++ {
		tituloEleitoralDigit, _ := strconv.Atoi(string(t[i]))
		sum += tituloEleitoralDigit * (i - 1)
	}

	digit := sum % 11
	if digit != 0 {
		return string(t[11]) == strconv.Itoa(digit%10)
	}

	stateCode := string(t[8:10])
	if stateCode == "01" || stateCode == "02" {
		return string(t[11]) == strconv.Itoa(1)
	}

	return string(t[11]) == strconv.Itoa(digit)
}
