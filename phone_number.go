package brazil

import (
	"regexp"
)

func ParseBrPhoneNumber(phoneNumber string) Phone {
	var (
		countryCode, areaCode, number Number
	)

	phoneNumber = regexp.MustCompile(`[^0-9]`).ReplaceAllString(phoneNumber, "")
	if len(phoneNumber) == 13 {
		countryCode.number = phoneNumber[:1]
		phoneNumber = phoneNumber[1:]
	}
	if len(phoneNumber) == 11 {
		areaCode.number = phoneNumber[:1]
		phoneNumber = phoneNumber[1:]
	}
	if len(phoneNumber) == 9 {
		number.number = phoneNumber
	}
	return Phone{
		countryCode: countryCode,
		areaCode:    areaCode,
		number:      number,
	}
}

func (p Phone) CountryCode() string {
	return p.countryCode.number
}

func (p Phone) AreaCode() string {
	return p.areaCode.number
}

func (p Phone) Number() string {
	return p.number.number
}

func (p *Phone) IsValid() bool {
	return false
}

func (p Phone) Errors() []error {
	var errors []error
	return errors
}

func RandomPhoneNumber() string {
	return ""
}

func (p Phone) numberIsValid() Validation {
	return Validation{
		Valid:  true,
		Reason: errValidDate,
	}
}
