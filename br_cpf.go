package brazil

import (
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func ParseCPF(number string) cpf {
	return cpf{
		Number: number,
	}
}

func EvaluateCPF(c cpf) Validation {
	var v Validation

	v = c.hasExpectedFormat()
	if v.Valid {
		v = c.isValid()
	}
	return Validation{
		Valid:  v.Valid,
		Reason: v.Reason,
	}
}

func GenerateCPF() string {
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

func (c cpf) hasExpectedFormat() Validation {
	var valid bool

	pattern1 := regexp.MustCompile(`^[0-9]{11}$`)
	valid = pattern1.MatchString(c.Number)

	if valid {
		return Validation{
			Valid:  true,
			Reason: nil,
		}
	}

	pattern2 := regexp.MustCompile(`^[0-9]{3}[\.][0-9]{3}[\.][0-9]{3}[-][0-9]{2}$`)
	valid = pattern2.MatchString(c.Number)

	if valid {
		return Validation{
			Valid:  true,
			Reason: nil,
		}
	}

	return Validation{
		Valid:  false,
		Reason: errIncorrectFormatCpf,
	}
}

func (c cpf) isValid() Validation {
	var sum int
	var digit int

	pattern := regexp.MustCompile(`[^0-9]`)
	cleanString := pattern.ReplaceAllString(c.Number, "")

	firstDigit, _ := strconv.Atoi(string(cleanString[9]))
	secondDigit, _ := strconv.Atoi(string(cleanString[10]))

	// False positives
	numbers, _ := strconv.Atoi(cleanString)
	if numbers%11111111111 == 0 {
		return Validation{
			Valid:  false,
			Reason: errInvalidCpf,
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
			Reason: errInvalidCpf,
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
			Reason: errInvalidCpf,
		}
	}

	// Success
	return Validation{
		Valid:  true,
		Reason: nil,
	}
}
