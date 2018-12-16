package brazil

import (
	"math/rand"
	"regexp"
	"strconv"
	"time"
)

// SUS struct
type sus struct {
	number susNumber
}

func (s sus) Number(mask bool) string {
	var sNumber = s.number
	if mask {
		return string(sNumber[:3]) + " " + string(sNumber[3:7]) + " " + string(sNumber[7:11]) + " " + string(sNumber[11:])
	}
	return string(sNumber)
}

func ParseSUS(number string) (sus, error) {
	var sNumber susNumber

	number = regexp.MustCompile(`[^0-9]`).ReplaceAllString(number, "")

	if len(number) != 15 {
		return sus{}, errIncorrectLenghtSusNumber
	}

	sNumber = susNumber(number)

	if sNumber.isFalsePositive() {
		return sus{}, errInvalidSusNumber
	}

	if !sNumber.isValid() {
		return sus{}, errInvalidSusNumber
	}

	return sus{
		number: sNumber,
	}, nil
}

func RandomSUSNumber(mask bool) string {
	var (
		source               = rand.NewSource(time.Now().UnixNano())
		possibleFirstNumbers = []int{1, 2, 7, 8, 9}
		numbers              string
	)

	r := rand.New(source)

	firstNumber := possibleFirstNumbers[r.Int31n(4)]
	firstNumberString := strconv.Itoa(firstNumber)
	if firstNumber == 1 || firstNumber == 2 {
		numbers = strconv.FormatInt(r.Int63n(8999999999)+1000000000, 10) + "00"
	} else {
		numbers = strconv.FormatInt(r.Int63n(899999999999)+100000000000, 10)
	}

	sum := firstNumber * 15
	for i := 0; i < 12; i++ {
		number, _ := strconv.Atoi(string(numbers[i]))
		sum += number * (14 - i)
	}
	remainder := sum % 11

	penultimate := 0
	ultimate := remainder

	if remainder%11 > 0 {
		penultimate = 0
		ultimate = 11 - remainder
		if (11 - remainder) == 10 {
			penultimate = 1
			ultimate = 11 - (sum+2*penultimate)%11
			if (sum+2*penultimate)%11 == 0 {
				ultimate = 0
			}
		}
	}

	pString := strconv.Itoa(penultimate)
	uString := strconv.Itoa(ultimate)

	sString := firstNumberString + numbers + pString + uString
	if mask {
		return sString[:3] + " " + sString[3:7] + " " + sString[7:11] + " " + sString[11:]
	}
	return sString
}

type susNumber string

func (s susNumber) isFalsePositive() bool {
	var regex = regexp.MustCompile(`(^[1-2]\d{10}00[0-1]\d{1}$)|(^[7-9]\d{14}$)`)
	if regex.MatchString(string(s)) {
		return false
	}
	return true
}

func (s susNumber) isValid() bool {
	var sum int

	for i := 0; i < 15; i++ {
		susDigit, _ := strconv.Atoi(string(s[i]))
		sum += susDigit * (15 - i)
	}

	return sum%11 == 0
}
