package brazil

import (
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func ParseTituloEleitoral(n string) tituloEleitoral {
	var t number
	t.number = n
	return tituloEleitoral{
		number: t,
	}
}

func (t tituloEleitoral) Number() string {
	return t.number.number
}

func (t *tituloEleitoral) IsValid() bool {
	t.number.validation = t.numberIsValid()
	if t.number.validation.valid {
		t.valid = true
		return true
	}
	return false
}

func (t *tituloEleitoral) Errors() []error {
	var errors []error

	if t.valid {
		return nil
	}
	if !t.number.validation.valid {
		errors = append(errors, t.number.validation.reason)
	}

	return errors
}

func RandomTituloEleitoral() string {
	var i, sum int
	var digitlessTitulo string

	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	random := int(r.Int31n(89999999) + 10000000)
	randomStr := strconv.Itoa(random)
	stateCode := int(r.Int31n(27) + 1)
	stateCodeStr := strconv.Itoa(stateCode)

	if stateCode < 10 {
		digitlessTitulo = randomStr + "0" + stateCodeStr
	} else {
		digitlessTitulo = randomStr + stateCodeStr
	}

	// First digit
	sum = 0
	for i = 0; i < 8; i++ {
		number, _ := strconv.Atoi(string(digitlessTitulo[i]))
		sum = sum + number*(i+2)
	}
	firstDigit := sum % 11
	if firstDigit == 10 {
		firstDigit = 0
	}
	if firstDigit == 0 {
		switch stateCode {
		case 1:
			firstDigit = 1
			break
		case 2:
			firstDigit = 1
			break
		}
	}
	if firstDigit == 10 {
		firstDigit = 0
	}

	// Second digit
	sum = 0
	for i = 8; i < 10; i++ {
		number, _ := strconv.Atoi(string(digitlessTitulo[i]))
		sum = sum + number*(i-1)
	}
	sum = sum + firstDigit*9
	secondDigit := sum % 11
	if secondDigit == 10 {
		secondDigit = 0
	}
	if secondDigit == 0 {
		switch stateCode {
		case 1:
			secondDigit = 1
			break
		case 2:
			secondDigit = 1
			break
		}
	}
	if secondDigit == 10 {
		secondDigit = 0
	}

	return strings.Join([]string{
		digitlessTitulo,
		strconv.Itoa(firstDigit),
		strconv.Itoa(secondDigit),
	}, "")
}

func (t tituloEleitoral) numberIsValid() validation {
	if t.number.number == "" {
		return validation{
			valid:  false,
			reason: errFieldNumberIsRequired,
		}
	}

	cleanString := regexp.MustCompile(`[^0-9]`).ReplaceAllString(t.number.number, "")
	v1 := regexp.MustCompile(`^[0-9]{12}$`).MatchString(cleanString)
	v2 := regexp.MustCompile(`^[0-9]{14}$`).MatchString(cleanString)
	if !v1 && !v2 {
		return validation{
			valid:  false,
			reason: errIncorrectFormatTituloEleitoralNumber,
		}
	}

	var sum, digit int

	firstDigit, _ := strconv.Atoi(string(cleanString[10]))
	secondDigit, _ := strconv.Atoi(string(cleanString[11]))

	// False positives
	numbers, _ := strconv.Atoi(cleanString)
	if numbers == 0 {
		return validation{
			valid:  false,
			reason: errInvalidTituloEleitoralNumber,
		}
	}

	// First digit validation
	for i := 0; i < 8; i++ {
		number, _ := strconv.Atoi(string(cleanString[i]))
		sum = sum + number*(i+2)
	}
	digit = sum % 11
	if digit == 0 {
		stateCode, _ := strconv.Atoi(string(cleanString[8:10]))
		switch stateCode {
		case 1:
			digit = 1
			break
		case 2:
			digit = 1
			break
		}
	}
	if digit == 10 {
		digit = 0
	}

	if digit != firstDigit {
		return validation{
			valid:  false,
			reason: errInvalidTituloEleitoralNumber,
		}
	}

	// Second digit validation
	sum = 0
	for i := 8; i < 11; i++ {
		number, _ := strconv.Atoi(string(cleanString[i]))
		sum = sum + number*(i-1)
	}
	digit = sum % 11
	if digit == 0 {
		stateCode, _ := strconv.Atoi(string(cleanString[8:10]))
		switch stateCode {
		case 1:
			digit = 1
			break
		case 2:
			digit = 1
			break
		}
	}
	if digit == 10 {
		digit = 0
	}

	if digit != secondDigit {
		return validation{
			valid:  false,
			reason: errInvalidTituloEleitoralNumber,
		}
	}

	return validation{
		valid:  true,
		reason: errValidTituloEleitoralNumber,
	}
}
