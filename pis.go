package brazil

import (
	"math/rand"
	"regexp"
	"strconv"
	"time"
)

// PIS struct
type pis struct {
	number pisNumber
}

func (p pis) Number(mask bool) string {
	var pNumber = p.number

	if mask {
		return string(pNumber[:3]) + "." + string(pNumber[3:8]) + "." + string(pNumber[8:10]) + "-" + string(pNumber[10:])
	}
	return string(pNumber)
}

func ParsePIS(number string) (pis, error) {
	var pNumber pisNumber

	number = regexp.MustCompile(`[^0-9]`).ReplaceAllString(number, "")

	if len(number) != 11 && len(number) != 13 {
		return pis{}, errIncorrectLenghtPisNumber
	}

	pNumber = pisNumber(number)

	if pNumber.isFalsePositive() {
		return pis{}, errInvalidPisNumber
	}

	if !pNumber.hasValidDigit() {
		return pis{}, errInvalidPisNumber
	}

	return pis{
		number: pNumber,
	}, nil
}

func RandomPISNumber(mask bool) string {
	var (
		source      = rand.NewSource(time.Now().UnixNano())
		multipliers = []int{3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
		sum         int
	)

	r := rand.New(source)
	pNumber := int(r.Int63n(8999999999) + 1000000000)
	pString := strconv.Itoa(pNumber)

	for i := 0; i < 10; i++ {
		number, _ := strconv.Atoi(string(pString[i]))
		sum = sum + number*multipliers[i]
	}
	digit := 11 - sum%11
	if digit >= 10 {
		digit = 0
	}

	if mask {
		return pString[:3] + "." + pString[3:8] + "." + pString[8:] + "-" + strconv.Itoa(digit)
	}
	return pString + strconv.Itoa(digit)
}

type pisNumber string

func (p pisNumber) isFalsePositive() bool {
	if string(p) == "00000000000" {
		return true
	}
	return false
}

func (p pisNumber) hasValidDigit() bool {
	var (
		multipliers = []int{3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
		sum         int
	)

	for i := 0; i < 10; i++ {
		pisDigit, _ := strconv.Atoi(string(p[i]))
		sum = sum + pisDigit*multipliers[i]
	}

	return string(p[10]) == strconv.Itoa((11-sum%11)%10) || string(p[10]) == strconv.Itoa((11-sum%11)%11)
}
