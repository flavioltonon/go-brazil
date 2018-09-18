package brazil

import (
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func (c *CPF) SetNumber(n string) {
	c.number.number = n
	c.number.validation = Validation{}
	c.valid = false
}

func (c *CPF) IsValid() bool {
	c.number.validation = c.numberIsValid()
	c.birthdate.validation = c.dateIsValid()
	if c.number.validation.Valid && c.birthdate.validation.Valid {
		c.valid = true
		return true
	}
	return false
}

func (c CPF) Errors() []error {
	var errors []error

	if c.valid {
		return nil
	}
	if !c.number.validation.Valid {
		errors = append(errors, c.number.validation.Reason)
	}
	if !c.birthdate.validation.Valid {
		errors = append(errors, c.birthdate.validation.Reason)
	}

	return errors
}

func RandomCPFNumber() string {
	var i, sum int

	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	random := int(r.Int31n(899999999) + 100000000)
	randomStr := strconv.Itoa(random)

	// First digit
	sum = 0
	for i = 0; i < 9; i++ {
		number, _ := strconv.Atoi(string(randomStr[i]))
		sum = sum + number*(10-i)
	}
	firstDigit := (sum * 10) % 11
	if firstDigit == 10 {
		firstDigit = 0
	}

	// Second digit
	sum = 0
	for i = 0; i < 9; i++ {
		number, _ := strconv.Atoi(string(randomStr[i]))
		sum = sum + number*(11-i)
	}
	sum = sum + firstDigit*2
	secondDigit := (sum * 10) % 11
	if secondDigit == 10 {
		secondDigit = 0
	}

	return strings.Join([]string{
		randomStr[:3],
		".",
		randomStr[3:6],
		".",
		randomStr[6:],
		"-",
		strconv.Itoa(firstDigit),
		strconv.Itoa(secondDigit),
	}, "")
}

func (c CPF) numberIsValid() Validation {
	if c.number.number == "" {
		return Validation{
			Valid:  false,
			Reason: errFieldNumberIsRequired,
		}
	}

	cleanString := regexp.MustCompile(`[^0-9]`).ReplaceAllString(c.number.number, "")
	v := regexp.MustCompile(`^[0-9]{11}$`).MatchString(cleanString)
	if !v {
		return Validation{
			Valid:  false,
			Reason: errIncorrectFormatCpfNumber,
		}
	}

	var sum, digit int

	firstDigit, _ := strconv.Atoi(string(cleanString[9]))
	secondDigit, _ := strconv.Atoi(string(cleanString[10]))

	// False positives
	numbers, _ := strconv.Atoi(cleanString)
	if numbers%11111111111 == 0 {
		return Validation{
			Valid:  false,
			Reason: errInvalidCpfNumber,
		}
	}

	// First digit Validation
	for i := 0; i < 9; i++ {
		number, _ := strconv.Atoi(string(cleanString[i]))
		sum = sum + number*(10-i)
	}
	digit = (sum * 10) % 11
	if digit == 10 {
		digit = 0
	}
	if digit != firstDigit {
		return Validation{
			Valid:  false,
			Reason: errInvalidCpfNumber,
		}
	}

	// Second digit validation
	sum = 0
	for i := 0; i < 10; i++ {
		number, _ := strconv.Atoi(string(cleanString[i]))
		sum = sum + number*(11-i)
	}
	digit = (sum * 10) % 11
	if digit == 10 {
		digit = 0
	}
	if digit != secondDigit {
		return Validation{
			Valid:  false,
			Reason: errInvalidCpfNumber,
		}
	}

	return Validation{
		Valid:  true,
		Reason: errValidCpfNumber,
	}
}

func (c CPF) dateIsValid() Validation {
	if !c.birthdate.notNull {
		return Validation{
			Valid:  true,
			Reason: errFieldDateNotRequired,
		}
	}

	if !c.birthdate.IsValid() {
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
