package brazil

import (
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func ParseCPF(n string) cpf {
	var c number
	c.number = n
	return cpf{
		number: c,
	}
}

func (c cpf) Number() string {
	return c.number.number
}

func (c *cpf) IsValid() bool {
	c.number.validation = c.numberIsValid()
	if c.number.validation.valid {
		c.valid = true
		return true
	}
	return false
}

func (c cpf) Errors() []error {
	var errors []error

	if c.valid {
		return nil
	}
	if !c.number.validation.valid {
		errors = append(errors, c.number.validation.reason)
	}

	return errors
}

func RandomCPF() string {
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

func (c cpf) numberIsValid() validation {
	if c.number.number == "" {
		return validation{
			valid:  false,
			reason: errFieldNumberIsRequired,
		}
	}

	cleanString := regexp.MustCompile(`[^0-9]`).ReplaceAllString(c.number.number, "")
	v := regexp.MustCompile(`^[0-9]{11}$`).MatchString(cleanString)
	if !v {
		return validation{
			valid:  false,
			reason: errIncorrectFormatCpfNumber,
		}
	}

	var sum, digit int

	firstDigit, _ := strconv.Atoi(string(cleanString[9]))
	secondDigit, _ := strconv.Atoi(string(cleanString[10]))

	// False positives
	numbers, _ := strconv.Atoi(cleanString)
	if numbers%11111111111 == 0 {
		return validation{
			valid:  false,
			reason: errInvalidCpfNumber,
		}
	}

	// First digit validation
	for i := 0; i < 9; i++ {
		number, _ := strconv.Atoi(string(cleanString[i]))
		sum = sum + number*(10-i)
	}
	digit = (sum * 10) % 11
	if digit == 10 {
		digit = 0
	}
	if digit != firstDigit {
		return validation{
			valid:  false,
			reason: errInvalidCpfNumber,
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
		return validation{
			valid:  false,
			reason: errInvalidCpfNumber,
		}
	}

	return validation{
		valid:  true,
		reason: errValidCpfNumber,
	}
}
