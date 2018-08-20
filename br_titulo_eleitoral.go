package brazil

import (
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func ParseTituloEleitoral(number string) tituloEleitoral {
	return tituloEleitoral{
		Number: number,
	}
}

func EvaluateTituloEleitoral(t tituloEleitoral) Validation {
	var v Validation

	v = t.hasExpectedFormat()
	if v.Valid {
		v = t.isValid()
	}
	return Validation{
		Valid:  v.Valid,
		Reason: v.Reason,
	}
}

func GenerateTituloEleitoral() string {
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

func (t tituloEleitoral) hasExpectedFormat() Validation {
	var valid bool

	// First pattern: 12 digits
	pattern1 := regexp.MustCompile(`^[0-9]{12}$`)
	valid = pattern1.MatchString(t.Number)
	if valid {
		return Validation{
			Valid:  true,
			Reason: nil,
		}
	}

	// Second pattern: 14 digits
	pattern2 := regexp.MustCompile(`^[0-9]{14}$`)
	valid = pattern2.MatchString(t.Number)
	if valid {
		return Validation{
			Valid:  true,
			Reason: nil,
		}
	}

	return Validation{
		Valid:  false,
		Reason: errIncorrectFormatTituloEleitoral,
	}
}

func (t tituloEleitoral) isValid() Validation {
	var sum int
	var digit int

	pattern := regexp.MustCompile(`[^0-9]`)
	cleanString := pattern.ReplaceAllString(t.Number, "")

	firstDigit, _ := strconv.Atoi(string(cleanString[10]))
	secondDigit, _ := strconv.Atoi(string(cleanString[11]))

	// False positives
	numbers, _ := strconv.Atoi(cleanString)
	if numbers == 0 {
		return Validation{
			Valid:  false,
			Reason: errInvalidTituloEleitoral,
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
			Reason: errInvalidTituloEleitoral,
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
			Reason: errInvalidTituloEleitoral,
		}
	}

	// Success
	return Validation{
		Valid:  true,
		Reason: nil,
	}
}
