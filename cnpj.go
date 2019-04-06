package brazil

import (
	"math/rand"
	"regexp"
	"strconv"
	"time"
)

// CNPJ struct
type cnpj struct {
	number cnpjNumber
	valid  bool
}

func (c cnpj) Number(mask bool) string {
	if c.valid && mask {
		return string(c.number[:2]) +
			"." +
			string(c.number[2:5]) +
			"." +
			string(c.number[5:8]) +
			"/" +
			string(c.number[8:12]) +
			"-" +
			string(c.number[12:])
	}
	return string(c.number)
}

func ParseCNPJ(number string) (cnpj, error) {
	number = regexp.MustCompile(`[^0-9]`).ReplaceAllString(number, "")

	if len(number) != 14 {
		return cnpj{}, errIncorrectLenghtCNPJNumber
	}

	cnpjNumber := cnpjNumber(number)

	if !cnpjNumber.hasValidFirstDigit() {
		return cnpj{}, errInvalidCNPJFirstDigit
	}

	if !cnpjNumber.hasValidSecondDigit() {
		return cnpj{}, errInvalidCNPJSecondDigit
	}

	return cnpj{
		number: cnpjNumber,
		valid:  true,
	}, nil
}

func RandomCNPJNumber(mask bool) string {
	var multipliers = []int{6, 5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}

	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	cnpjNumber := int(r.Int63n(899999999999) + 100000000000)
	cnpjString := strconv.Itoa(cnpjNumber)

	// Calculate first digit
	sum := 0
	for i := 0; i < 12; i++ {
		number, _ := strconv.Atoi(string(cnpjString[i]))
		sum += number * multipliers[i+1]
	}
	firstDigit := 0
	if sum%11 >= 2 {
		firstDigit = 11 - sum%11
	}

	// Calculate second digit
	sum = 0
	for i := 0; i < 12; i++ {
		number, _ := strconv.Atoi(string(cnpjString[i]))
		sum += number * multipliers[i]
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

type cnpjNumber string

func (c cnpjNumber) hasValidFirstDigit() bool {
	var (
		multipliers = []int{5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
		sum         int
	)

	for i := 0; i < 12; i++ {
		cnpjDigit, _ := strconv.Atoi(string(c[i]))
		sum += cnpjDigit * multipliers[i]
	}
	if sum%11 < 2 {
		return string(c[12]) == strconv.Itoa(0)
	}

	return string(c[12]) == strconv.Itoa(11-sum%11)
}

func (c cnpjNumber) hasValidSecondDigit() bool {
	var (
		multipliers = []int{6, 5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
		sum         int
	)

	for i := 0; i < 13; i++ {
		cnpjDigit, _ := strconv.Atoi(string(c[i]))
		sum += cnpjDigit * multipliers[i]
	}

	if sum%11 < 2 {
		return string(c[13]) == strconv.Itoa(0)
	}
	return string(c[13]) == strconv.Itoa(11-sum%11)
}
