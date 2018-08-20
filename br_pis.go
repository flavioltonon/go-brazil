package brazil

import (
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func ParsePIS(number string) pis {
	return pis{
		Number: number,
	}
}

func EvaluatePIS(p pis) Validation {
	var v Validation

	v = p.hasExpectedFormat()
	if v.Valid {
		v = p.isValid()
	}
	return Validation{
		Valid:  v.Valid,
		Reason: v.Reason,
	}
}

func GeneratePIS() string {
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

func (p pis) hasExpectedFormat() Validation {
	var valid bool

	pattern1 := regexp.MustCompile(`^[0-9]{11}$`)
	valid = pattern1.MatchString(p.Number)
	if valid {
		return Validation{
			Valid:  true,
			Reason: nil,
		}
	}

	pattern2 := regexp.MustCompile(`^[0-9]{3}[\.][0-9]{5}[\.][0-9]{2}[-][0-9]{1}$`)
	valid = pattern2.MatchString(p.Number)
	if valid {
		return Validation{
			Valid:  true,
			Reason: nil,
		}
	}

	return Validation{
		Valid:  false,
		Reason: errIncorrectFormatPis,
	}
}

func (p pis) isValid() Validation {
	var sum int
	var digit int

	var multipliers = []int{3, 2, 9, 8, 7, 6, 5, 4, 3, 2}

	pattern := regexp.MustCompile(`[^0-9]`)
	cleanString := pattern.ReplaceAllString(p.Number, "")

	onlyDigit, _ := strconv.Atoi(string(cleanString[10]))

	// False positives
	numbers, _ := strconv.Atoi(cleanString)
	if numbers == 0 {
		return Validation{
			Valid:  false,
			Reason: errInvalidPis,
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
			Reason: errInvalidPis,
		}
	}

	// Success
	return Validation{
		Valid:  true,
		Reason: nil,
	}
}
