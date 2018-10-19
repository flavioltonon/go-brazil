package brazil

import (
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func ParsePIS(n string) pis {
	var p number
	p.number = n
	return pis{
		number: p,
	}
}

func (p pis) Number() string {
	return p.number.number
}

func (p *pis) IsValid() bool {
	p.number.validation = p.numberIsValid()
	if p.number.validation.valid {
		p.valid = true
		return true
	}
	return false
}

func (p *pis) Errors() []error {
	var errors []error

	if p.valid {
		return nil
	}
	if !p.number.validation.valid {
		errors = append(errors, p.number.validation.reason)
	}

	return errors
}

func RandomPIS() string {
	var i, sum int
	var multipliers = []int{3, 2, 9, 8, 7, 6, 5, 4, 3, 2}

	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	random := int(r.Int63n(8999999999) + 1000000000)
	randomStr := strconv.Itoa(random)

	// Only digit
	sum = 0
	for i = 0; i < 10; i++ {
		number, _ := strconv.Atoi(string(randomStr[i]))
		sum = sum + number*multipliers[i]
	}
	onlyDigit := 11 - sum%11
	switch onlyDigit {
	case 10:
		onlyDigit = 0
		break
	case 11:
		onlyDigit = 0
		break
	}

	return strings.Join([]string{
		randomStr[:3],
		".",
		randomStr[3:8],
		".",
		randomStr[8:],
		"-",
		strconv.Itoa(onlyDigit),
	}, "")
}

func (p pis) numberIsValid() validation {
	if p.number.number == "" {
		return validation{
			valid:  false,
			reason: errFieldNumberIsRequired,
		}
	}

	cleanString := regexp.MustCompile(`[^0-9]`).ReplaceAllString(p.number.number, "")
	v1 := regexp.MustCompile(`^[0-9]{11}$`).MatchString(cleanString)
	v2 := regexp.MustCompile(`^[0-9]{13}$`).MatchString(cleanString)
	if !v1 && !v2 {
		return validation{
			valid:  false,
			reason: errIncorrectFormatPisNumber,
		}
	}

	var sum, digit int
	var multipliers = []int{3, 2, 9, 8, 7, 6, 5, 4, 3, 2}

	onlyDigit, _ := strconv.Atoi(string(cleanString[10]))

	// False positives
	numbers, _ := strconv.Atoi(cleanString)
	if numbers == 0 {
		return validation{
			valid:  false,
			reason: errInvalidPisNumber,
		}
	}

	// Only digit validation
	for i := 0; i < 10; i++ {
		number, _ := strconv.Atoi(string(cleanString[i]))
		sum = sum + number*multipliers[i]
	}
	digit = 11 - sum%11
	switch digit {
	case 10:
		digit = 0
		break
	case 11:
		digit = 0
		break
	}
	if digit != onlyDigit {
		return validation{
			valid:  false,
			reason: errInvalidPisNumber,
		}
	}

	return validation{
		valid:  true,
		reason: errValidPisNumber,
	}
}
