package brazil

import (
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// ParseAlphanumericCNPJ validates CNPJ numbers in the alphanumeric format
// introduced by IN RFB n° 2.229/2024: positions 0-11 accept 0-9 or A-Z and
// the check digits (positions 12-13) remain numeric. Legacy numeric CNPJ
// numbers are also accepted. Lowercase letters are normalized to uppercase.
func ParseAlphanumericCNPJ(number string) (cnpj, error) {
	number = strings.ToUpper(regexp.MustCompile(`[./-]`).ReplaceAllString(number, ""))

	if len(number) != 14 {
		return cnpj{}, ErrIncorrectLenghtCNPJNumber
	}

	if !regexp.MustCompile(`^[0-9A-Z]{12}[0-9]{2}$`).MatchString(number) {
		return cnpj{}, ErrInvalidCNPJCharacter
	}

	if strings.Count(number, string(number[0])) == 14 {
		return cnpj{}, ErrRepeatedCNPJNumber
	}

	cnpjNumber := cnpjNumber(number)

	if !cnpjNumber.hasValidAlphanumericFirstDigit() {
		return cnpj{}, ErrInvalidCNPJFirstDigit
	}

	if !cnpjNumber.hasValidAlphanumericSecondDigit() {
		return cnpj{}, ErrInvalidCNPJSecondDigit
	}

	return cnpj{
		number: cnpjNumber,
		valid:  true,
	}, nil
}

func RandomAlphanumericCNPJNumber(mask bool) string {
	var multipliers = []int{6, 5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
	const charset = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	base := make([]byte, 12)
	for i := range base {
		base[i] = charset[r.Intn(len(charset))]
	}
	cnpjString := string(base)

	// Calculate first digit
	sum := 0
	for i := 0; i < 12; i++ {
		sum += alphanumericCharValue(cnpjString[i]) * multipliers[i+1]
	}
	firstDigit := 0
	if sum%11 >= 2 {
		firstDigit = 11 - sum%11
	}

	// Calculate second digit
	sum = 0
	for i := 0; i < 12; i++ {
		sum += alphanumericCharValue(cnpjString[i]) * multipliers[i]
	}
	sum += firstDigit * multipliers[12]
	secondDigit := 0
	if sum%11 >= 2 {
		secondDigit = 11 - sum%11
	}

	if mask {
		return cnpjString[:2] + "." + cnpjString[2:5] + "." + cnpjString[5:8] + "/" + cnpjString[8:12] + "-" + strconv.Itoa(firstDigit) + strconv.Itoa(secondDigit)
	}
	return cnpjString + strconv.Itoa(firstDigit) + strconv.Itoa(secondDigit)
}

// alphanumericCharValue converts a CNPJ character to its check digit value
// as defined by IN RFB n° 2.229/2024: value = ASCII(char) - 48, so that
// '0'-'9' map to 0-9 and 'A'-'Z' map to 17-42.
func alphanumericCharValue(c byte) int {
	return int(c) - 48
}

func (c cnpjNumber) hasValidAlphanumericFirstDigit() bool {
	var (
		multipliers = []int{5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
		sum         int
	)

	for i := 0; i < 12; i++ {
		sum += alphanumericCharValue(c[i]) * multipliers[i]
	}
	if sum%11 < 2 {
		return string(c[12]) == strconv.Itoa(0)
	}

	return string(c[12]) == strconv.Itoa(11-sum%11)
}

func (c cnpjNumber) hasValidAlphanumericSecondDigit() bool {
	var (
		multipliers = []int{6, 5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
		sum         int
	)

	for i := 0; i < 13; i++ {
		sum += alphanumericCharValue(c[i]) * multipliers[i]
	}

	if sum%11 < 2 {
		return string(c[13]) == strconv.Itoa(0)
	}
	return string(c[13]) == strconv.Itoa(11-sum%11)
}
