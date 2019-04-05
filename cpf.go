package brazil

import (
	"math/rand"
	"regexp"
	"strconv"
	"time"
)

// CPF struct
type cpf struct {
	number cpfNumber
	valid  bool
}

func (c cpf) Number(mask bool) string {
	if c.valid && mask {
		return string(c.number[:3]) +
			"." +
			string(c.number[3:6]) +
			"." +
			string(c.number[6:9]) +
			"-" +
			string(c.number[9:])
	}
	return string(c.number)
}

func ParseCPF(number string) (cpf, error) {
	number = regexp.MustCompile(`[^0-9]`).ReplaceAllString(number, "")

	if len(number) != 11 {
		return cpf{}, errIncorrectLenghtCPFNumber
	}

	cpfNumber := cpfNumber(number)

	if cpfNumber.isFalsePositive() {
		return cpf{}, errInvalidCPFNumber
	}

	if !cpfNumber.hasValidFirstDigit() {
		return cpf{}, errInvalidCPFNumber
	}

	if !cpfNumber.hasValidSecondDigit() {
		return cpf{}, errInvalidCPFNumber
	}

	return cpf{
		number: cpfNumber,
		valid:  true,
	}, nil
}

func RandomCPFNumber(mask bool) string {
	var (
		source = rand.NewSource(time.Now().UnixNano())
		sum    int
	)

	r := rand.New(source)
	cNumber := int(r.Int31n(899999999) + 100000000)
	cString := strconv.Itoa(cNumber)

	// Calculate first digit
	for i := 0; i < 9; i++ {
		number, _ := strconv.Atoi(string(cString[i]))
		sum += number * (10 - i)
	}
	firstDigit := ((sum * 10) % 11) % 10

	// Calculate second digit
	sum = 0
	for i := 0; i < 9; i++ {
		number, _ := strconv.Atoi(string(cString[i]))
		sum += number * (11 - i)
	}
	sum += firstDigit * 2
	secondDigit := ((sum * 10) % 11) % 10

	if mask {
		return cString[:3] + "." + cString[3:6] + "." + cString[6:] + "-" + strconv.Itoa(firstDigit) + strconv.Itoa(secondDigit)
	}
	return cString + strconv.Itoa(firstDigit) + strconv.Itoa(secondDigit)
}

type cpfNumber string

func (c cpfNumber) isFalsePositive() bool {
	cpf, _ := strconv.Atoi(string(c))
	if cpf%11111111111 == 0 {
		return true
	}
	return false
}

func (c cpfNumber) hasValidFirstDigit() bool {
	var sum int

	for i := 0; i < 9; i++ {
		cpfDigit, _ := strconv.Atoi(string(c[i]))
		sum += cpfDigit * (10 - i)
	}

	return string(c[9]) == strconv.Itoa(((sum*10)%11)%10)
}

func (c cpfNumber) hasValidSecondDigit() bool {
	var sum int

	for i := 0; i < 10; i++ {
		cpfDigit, _ := strconv.Atoi(string(c[i]))
		sum += cpfDigit * (11 - i)
	}

	return string(c[10]) == strconv.Itoa(((sum*10)%11)%10)
}
