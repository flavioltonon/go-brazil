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
}

func (t tituloEleitoral) Number(mask bool) string {
	var tNumber = t.number

	if mask {
		return string(tNumber[:4]) + " " + string(tNumber[4:8]) + " " + string(tNumber[8:])
	}
	return string(tNumber)
}

type tituloEleitoralNumber string

func ParseTituloEleitoral(number string) (tituloEleitoral, error) {
	var tNumber tituloEleitoralNumber

	number = regexp.MustCompile(`[^0-9]`).ReplaceAllString(number, "")

	if len(number) != 12 && len(number) != 14 {
		return tituloEleitoral{}, errIncorrectLenghtTituloEleitoralNumber
	}

	tNumber = tituloEleitoralNumber(number)

	if tNumber.isFalsePositive() {
		return tituloEleitoral{}, errInvalidTituloEleitoralNumber
	}

	if !tNumber.hasValidFirstDigit() {
		return tituloEleitoral{}, errInvalidTituloEleitoralNumber
	}

	if !tNumber.hasValidSecondDigit() {
		return tituloEleitoral{}, errInvalidTituloEleitoralNumber
	}

	return tituloEleitoral{
		number: tNumber,
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
	if firstDigit == 0 {
		if stateCodeString == "01" || stateCodeString == "02" {
			firstDigit = 1
		}
	}
	if firstDigit == 10 {
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
	if secondDigit == 0 {
		if stateCodeString == "01" || stateCodeString == "02" {
			secondDigit = 1
		}
	}
	if secondDigit == 10 {
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
	var sum, digit int

	for i := 0; i < 8; i++ {
		tituloEleitoralDigit, _ := strconv.Atoi(string(t[i]))
		sum += tituloEleitoralDigit * (i + 2)
	}
	digit = sum % 11
	if digit == 0 {
		stateCode, _ := strconv.Atoi(string(t[8:10]))
		switch stateCode {
		case 1:
			digit = 1
			break
		case 2:
			digit = 1
			break
		}
	}

	return string(t[10]) == strconv.Itoa(digit%10) || string(t[10]) == strconv.Itoa(digit)
}

func (t tituloEleitoralNumber) hasValidSecondDigit() bool {
	var sum, digit int

	for i := 8; i < 11; i++ {
		tituloEleitoralDigit, _ := strconv.Atoi(string(t[i]))
		sum += tituloEleitoralDigit * (i - 1)
	}
	digit = sum % 11
	if digit == 0 {
		stateCode, _ := strconv.Atoi(string(t[8:10]))
		switch stateCode {
		case 1:
			digit = 1
			break
		case 2:
			digit = 1
			break
		}
	}

	return string(t[11]) == strconv.Itoa(digit%10) || string(t[11]) == strconv.Itoa(digit)
}
