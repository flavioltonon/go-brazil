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
	valid       bool
}

func (m mobile) FullNumber(mask bool) string {
	if m.valid && mask {
		return "+" + string(m.countryCode) + "(" + string(m.areaCode) + ")" + string(m.number[:5]) + "-" + string(m.number[5:])
	}

	return string(m.countryCode) + string(m.areaCode) + string(m.number)
}

func (m mobile) CountryCode(mask bool) string {
	if m.valid && mask {
		return "+" + string(m.countryCode)
	}
	return string(m.countryCode)
}

func (m mobile) AreaCode(mask bool) string {
	if m.valid && mask {
		return "(" + string(m.areaCode) + ")"
	}
	return string(m.areaCode)
}

func (m mobile) Number(mask bool) string {
	if m.valid && mask {
		return string(m.number[:5]) + "-" + string(m.number[5:])
	}
	return string(m.number)
}

func ParseMobile(number string) (mobile, error) {
	number = regexp.MustCompile(`[^0-9]`).ReplaceAllString(number, "")
	if len(number) != 13 {
		return mobile{}, ErrIncorrectFormatMobileNumber
	}

	countryCode := countryCode(number[:2])
	areaCode := areaCode(number[2:4])
	mobileNumber := mobileNumber(number[4:])

	if !countryCode.isValid() {
		return mobile{}, ErrInvalidBrazilianCountryCode
	}

	if !areaCode.isValid() {
		return mobile{}, ErrInvalidBrazilianAreaCode
	}

	if !mobileNumber.isValid() {
		return mobile{}, ErrInvalidBrazilianMobileNumber
	}

	return mobile{
		countryCode: countryCode,
		areaCode:    areaCode,
		number:      mobileNumber,
		valid:       true,
	}, nil
}

func RandomMobileFullNumber(mask bool) string {
	mobile, err := ParseMobile(standardCountryCode + RandomAreaCode() + RandomNumber())
	if err != nil {
		panic("Internal Error")
	}
	return mobile.FullNumber(mask)
}

type countryCode string

var standardCountryCode = "55"

func (c countryCode) isValid() bool {
	return c == countryCode(standardCountryCode)
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

func RandomAreaCode() string {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	areaCodes := ListAreaCodes()

	return string(areaCodes[int(r.Int31n(int32(len(areaCodes))))])
}

type mobileNumber string

func (m mobileNumber) isValid() bool {
	if string(m[0]) != "9" {
		return false
	}
	return true
}

func RandomNumber() string {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	return string(mobileNumber(strconv.Itoa(900000000 + int(r.Int63n(99999999)))))
}
