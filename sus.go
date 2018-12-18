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
		possibleFirstNumbers = []string{"1", "2", "7", "8", "9"}
		susNumber            string
		numbers              string
		sum                  int
	)

	r := rand.New(source)

	firstNumber := possibleFirstNumbers[r.Int31n(4)]
	switch firstNumber {
	case "1", "2":
		numbers = firstNumber + strconv.FormatInt(r.Int63n(8999999999)+1000000000, 10)
		for i := 0; i < 11; i++ {
			susDigit, _ := strconv.Atoi(string(numbers[i]))
			sum += susDigit * (15 - i)
		}
		if sum%11 == 1 {
			susNumber = numbers + "001" + strconv.Itoa(11-(sum+2)%11)
		} else {
			susNumber = numbers + "000" + strconv.Itoa((11-sum%11)%11)
		}
	case "7", "8", "9":
		numbers = firstNumber + strconv.FormatInt(r.Int63n(899999999999)+100000000000, 10)
		for i := 0; i < 13; i++ {
			susDigit, _ := strconv.Atoi(string(numbers[i]))
			sum += susDigit * (15 - i)
		}
		if sum%11 == 1 {
			susNumber = numbers + "1" + strconv.Itoa(11-(sum+2)%11)
		} else {
			susNumber = numbers + "0" + strconv.Itoa((11-sum%11)%11)
		}
	}

	if mask {
		return susNumber[:3] + " " + susNumber[3:7] + " " + susNumber[7:11] + " " + susNumber[11:]
	}
	return susNumber
}

type susNumber string

func (s susNumber) isValid() bool {
	var sum int

	if !regexp.MustCompile(`(^[1-2]\d{10}00[0-1]\d{1}$)|(^[7-9]\d{14}$)`).MatchString(string(s)) {
		return false
	}

	switch string(s[0]) {
	case "1", "2":
		for i := 0; i < 11; i++ {
			susDigit, _ := strconv.Atoi(string(s[i]))
			sum += susDigit * (15 - i)
		}
		if sum%11 == 1 {
			return string(s[11:15]) == "001"+strconv.Itoa(11-(sum+2)%11)
		}
		return string(s[11:15]) == "000"+strconv.Itoa((11-sum%11)%11)
	case "7", "8", "9":
		for i := 0; i < 15; i++ {
			susDigit, _ := strconv.Atoi(string(s[i]))
			sum += susDigit * (15 - i)
		}
		return sum%11 == 0
	default:
		return false
	}
}
