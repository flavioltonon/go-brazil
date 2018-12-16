package brazil

import (
	"math/rand"
	"regexp"
	"strconv"
	"time"
)

// Mobile struct
type mobile struct {
	countryCode countryCode
	areaCode    areaCode
	number      mobileNumber
}

func (m mobile) FullNumber(mask bool) string {
	if mask {
		return "+" + string(m.countryCode) + "(" + string(m.areaCode) + ")" + string(m.number[:5]) + "-" + string(m.number[5:])
	}
	return string(m.countryCode) + string(m.areaCode) + string(m.number)
}

func (m mobile) CountryCode(mask bool) string {
	var cCode = m.countryCode

	if mask {
		return "+" + string(cCode)
	}
	return string(cCode)
}

func (m mobile) AreaCode(mask bool) string {
	var aCode = m.areaCode

	if mask {
		return "(" + string(aCode) + ")"
	}
	return string(aCode)
}

func (m mobile) Number(mask bool) string {
	var mNumber = m.number
	if mask {
		return string(mNumber[:5]) + "-" + string(mNumber[5:])
	}
	return string(mNumber)
}

func ParseMobile(number string) (mobile, error) {
	var (
		cCode   countryCode
		aCode   areaCode
		mNumber mobileNumber
	)

	number = regexp.MustCompile(`[^0-9]`).ReplaceAllString(number, "")

	if len(number) == 13 {
		cCode = countryCode(number[:2])
		number = number[2:]
	}
	if len(number) == 11 {
		aCode = areaCode(number[:2])
		number = number[2:]
	}

	mNumber = mobileNumber(number)

	if len(mNumber) != 9 {
		return mobile{}, errIncorrectFormatMobileNumber
	}

	if len(cCode) > 0 && !cCode.isValid() {
		return mobile{}, errInvalidBrazilianCountryCode
	}

	if len(aCode) > 0 && !aCode.isValid() {
		return mobile{}, errInvalidBrazilianAreaCode
	}

	if len(mNumber) > 0 && !mNumber.isValid() {
		return mobile{}, errInvalidBrazilianMobileNumber
	}

	return mobile{
		countryCode: cCode,
		areaCode:    aCode,
		number:      mNumber,
	}, nil
}

func RandomMobileFullNumber(mask bool) string {
	var (
		cCode   = CountryCode(mask)
		aCode   = RandomAreaCode(mask)
		mNumber = RandomNumber(mask)
	)

	if mask {
		return "+" + cCode + "(" + aCode + ")" + mNumber
	}
	return cCode + aCode + mNumber
}

type countryCode string

func CountryCode(mask bool) string {
	if mask {
		return "+55"
	}
	return "55"
}

func (c countryCode) isValid() bool {
	return c == countryCode(CountryCode(false))
}

type areaCode string

func ListAreaCodes() []areaCode {
	return []areaCode{
		"11", "12", "13", "14", "15", "16", "17", "18", "19",
		"21", "22", "24", "27", "28",
		"31", "32", "33", "34", "35", "37", "38",
		"41", "42", "43", "44", "45", "46", "47", "48", "49",
		"51", "53", "54", "55",
		"61", "62", "63", "64", "65", "66", "67", "68", "69",
		"71", "73", "74", "75", "77", "79",
		"81", "82", "83", "84", "85", "86", "87", "88", "89",
		"91", "92", "93", "94", "95", "96", "97", "98", "99",
	}
}

func (a areaCode) isValid() bool {
	for _, areaCode := range ListAreaCodes() {
		if a == areaCode {
			return true
		}
	}
	return false
}

func RandomAreaCode(mask bool) string {
	var source = rand.NewSource(time.Now().UnixNano())

	r := rand.New(source)

	aCode := ListAreaCodes()[int(r.Int63n(59))]

	if mask {
		return "(" + string(aCode) + ")"
	}
	return string(aCode)
}

type mobileNumber string

func (m mobileNumber) isValid() bool {
	if string(m[0]) != "9" {
		return false
	}
	return true
}

func RandomNumber(mask bool) string {
	var (
		source  = rand.NewSource(time.Now().UnixNano())
		mNumber mobileNumber
	)

	r := rand.New(source)

	mNumber = mobileNumber(strconv.Itoa(900000000 + int(r.Int63n(99999999))))

	if mask {
		return string(mNumber[:5]) + "-" + string(mNumber[5:])
	}
	return string(mNumber)
}
