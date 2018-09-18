package brazil

import (
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func (t *TituloEleitoral) Number(n string) {
	t.number.number = n
	t.number.validation = Validation{}
	t.valid = false
}

func (t *TituloEleitoral) IssueDate(r BrDate) {
	t.issueDate = r
	t.issueDate.validation = Validation{}
	t.valid = false
}

func (t TituloEleitoral) GetNumber() string {
	return t.number.number
}

func (t *TituloEleitoral) IsValid() bool {
	t.number.validation = t.numberIsValid()
	t.issueDate.validation = t.dateIsValid()
	if t.number.validation.Valid && t.issueDate.validation.Valid {
		t.valid = true
		return true
	}
	return false
}

func (t *TituloEleitoral) Errors() []error {
	var errors []error

	if t.valid {
		return nil
	}
	if !t.number.validation.Valid {
		errors = append(errors, t.number.validation.Reason)
	}
	if !t.issueDate.validation.Valid {
		errors = append(errors, t.issueDate.validation.Reason)
	}

	return errors
}

func RandomTituloEleitoralNumber() string {
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

func (t TituloEleitoral) numberIsValid() Validation {
	if t.number.number == "" {
		return Validation{
			Valid:  false,
			Reason: errFieldNumberIsRequired,
		}
	}

	cleanString := regexp.MustCompile(`[^0-9]`).ReplaceAllString(t.number.number, "")
	v1 := regexp.MustCompile(`^[0-9]{12}$`).MatchString(cleanString)
	v2 := regexp.MustCompile(`^[0-9]{14}$`).MatchString(cleanString)
	if !v1 && !v2 {
		return Validation{
			Valid:  false,
			Reason: errIncorrectFormatTituloEleitoralNumber,
		}
	}

	var sum, digit int

	firstDigit, _ := strconv.Atoi(string(cleanString[10]))
	secondDigit, _ := strconv.Atoi(string(cleanString[11]))

	// False positives
	numbers, _ := strconv.Atoi(cleanString)
	if numbers == 0 {
		return Validation{
			Valid:  false,
			Reason: errInvalidTituloEleitoralNumber,
		}
	}

	// First digit Validation
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
		return Validation{
			Valid:  false,
			Reason: errInvalidTituloEleitoralNumber,
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
		return Validation{
			Valid:  false,
			Reason: errInvalidTituloEleitoralNumber,
		}
	}

	return Validation{
		Valid:  true,
		Reason: errValidTituloEleitoralNumber,
	}
}

func (t TituloEleitoral) dateIsValid() Validation {
	if !t.issueDate.notNull {
		return Validation{
			Valid:  true,
			Reason: errFieldDateNotRequired,
		}
	}

	if !t.issueDate.IsValid() {
		return Validation{
			Valid:  false,
			Reason: errIncorrectFormatDate,
		}
	}

	return Validation{
		Valid:  true,
		Reason: errValidDate,
	}
}
