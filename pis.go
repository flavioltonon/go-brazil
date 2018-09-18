package brazil

import (
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func (p *PIS) SetNumber(n string) {
	p.number.number = n
	p.number.validation = Validation{}
	p.valid = false
}

func (p *PIS) SetRegistration(r BrDate) {
	p.registration = r
	p.registration.validation = Validation{}
	p.valid = false
}

func (p *PIS) IsValid() bool {
	p.number.validation = p.numberIsValid()
	p.registration.validation = p.dateIsValid()
	if p.number.validation.Valid && p.registration.validation.Valid {
		p.valid = true
		return true
	}
	return false
}

func (p *PIS) Errors() []error {
	var errors []error

	if p.valid {
		return nil
	}
	if !p.number.validation.Valid {
		errors = append(errors, p.number.validation.Reason)
	}
	if !p.registration.validation.Valid {
		errors = append(errors, p.registration.validation.Reason)
	}

	return errors
}

func RandomPISNumber() string {
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

func (p PIS) numberIsValid() Validation {
	if p.number.number == "" {
		return Validation{
			Valid:  false,
			Reason: errFieldNumberIsRequired,
		}
	}

	cleanString := regexp.MustCompile(`[^0-9]`).ReplaceAllString(p.number.number, "")
	v1 := regexp.MustCompile(`^[0-9]{11}$`).MatchString(cleanString)
	v2 := regexp.MustCompile(`^[0-9]{13}$`).MatchString(cleanString)
	if !v1 && !v2 {
		return Validation{
			Valid:  false,
			Reason: errIncorrectFormatPisNumber,
		}
	}

	var sum, digit int
	var multipliers = []int{3, 2, 9, 8, 7, 6, 5, 4, 3, 2}

	onlyDigit, _ := strconv.Atoi(string(cleanString[10]))

	// False positives
	numbers, _ := strconv.Atoi(cleanString)
	if numbers == 0 {
		return Validation{
			Valid:  false,
			Reason: errInvalidPisNumber,
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
		return Validation{
			Valid:  false,
			Reason: errInvalidPisNumber,
		}
	}

	return Validation{
		Valid:  true,
		Reason: errValidPisNumber,
	}
}

func (p PIS) dateIsValid() Validation {
	if !p.registration.notNull {
		return Validation{
			Valid:  true,
			Reason: errFieldDateNotRequired,
		}
	}

	if !p.registration.IsValid() {
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
