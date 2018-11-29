package brazil

import (
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var areaCodes = []int{
	11, 12, 13, 14, 15, 16, 17, 18, 19,
	21, 22, 24, 27, 28,
	31, 32, 33, 34, 35, 37, 38,
	41, 42, 43, 44, 45, 46, 47, 48, 49,
	51, 53, 54, 55,
	61, 62, 63, 64, 65, 66, 67, 68, 69,
	71, 73, 74, 75, 77, 79,
	81, 82, 83, 84, 85, 86, 87, 88, 89,
	91, 92, 93, 94, 95, 96, 97, 98, 99,
}

func ParsePhoneNumber(phoneNumber string) phone {
	var (
		fullNumber            number
		countryCode, areaCode string
	)

	fullNumber.number = phoneNumber

	phoneNumber = regexp.MustCompile(`[^0-9]`).ReplaceAllString(phoneNumber, "")

	if len(phoneNumber) == 13 {
		countryCode = phoneNumber[:2]
		phoneNumber = phoneNumber[2:]
	}
	if len(phoneNumber) == 11 {
		areaCode = phoneNumber[:2]
		phoneNumber = phoneNumber[2:]
	}
	return phone{
		fullNumber:  fullNumber,
		countryCode: countryCode,
		areaCode:    areaCode,
		number:      phoneNumber,
	}
}

func (p phone) FullNumber() string {
	return p.fullNumber.number
}

func (p phone) CountryCode() string {
	return p.countryCode
}

func (p phone) AreaCode() string {
	return p.areaCode
}

func (p phone) Number() string {
	return p.number
}

func (p *phone) IsValid() bool {
	p.fullNumber.validation = p.numberIsValid()
	if p.fullNumber.validation.valid {
		p.valid = true
		return true
	}
	return false
}

func (p phone) Errors() []error {
	var errors []error

	if p.valid {
		return nil
	}
	if !p.fullNumber.validation.valid {
		errors = append(errors, p.fullNumber.validation.reason)
	}

	return errors
}

func RandomMobileNumber() string {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	countryCode := "55"

	randomAreaCode := areaCodes[int(r.Int63n(59))]
	randomAreaCodeStr := strconv.Itoa(randomAreaCode)

	randomNumber := int(r.Int63n(99999999) + 900000000)
	randomNumberStr := strconv.Itoa(randomNumber)

	return strings.Join([]string{
		"+",
		countryCode,
		"(",
		randomAreaCodeStr,
		")",
		randomNumberStr[:5],
		"-",
		randomNumberStr[5:],
	}, "")
}

func (p phone) numberIsValid() validation {
	if p.fullNumber.number == "" {
		return validation{
			valid:  false,
			reason: errFieldFullNumberIsRequired,
		}
	}

	v := regexp.MustCompile(`^(\d{9}|^(\d{11}|^(\d{13}))$`).MatchString(p.fullNumber.number)
	if !v {
		return validation{
			valid:  false,
			reason: errIncorrectFormatPhoneNumber,
		}
	}

	return validation{
		valid:  true,
		reason: errValidDate,
	}
}
