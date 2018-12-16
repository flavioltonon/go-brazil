package brazil

import "errors"

var (
	errIncorrectLenghtCpfNumber = errors.New("CPF numbers must contain 11 numbers")
	errInvalidCpfNumber         = errors.New("CPF number input is not valid")

	errIncorrectLenghtTituloEleitoralNumber = errors.New("CPF numbers must contain 12 or 14 numbers")
	errInvalidTituloEleitoralNumber         = errors.New("Título de eleitor number input is not valid")

	errIncorrectLenghtPisNumber = errors.New("PIS numbers must contain 11 or 13 numbers")
	errInvalidPisNumber         = errors.New("PIS number input is not valid")

	errIncorrectFormatDate = errors.New("Date format input is different than expected. Expected format: DD/MM/YYYY")
	errInvalidYearLimits   = errors.New("Date minYear must be lower than maxYear and higher than zero")
	errNotPastDate         = errors.New("Date must be in the past")
	errInvalidDate         = errors.New("Date input is not valid")
	errNullDate            = errors.New("Date has not been input")

	errIncorrectFormatMobileNumber  = errors.New("Mobile number format input is different than expected. Expected formats: XXXXXXXXXXXXX, +XX(XX)XXXXX-XXXX, (XX)XXXXX-XXXX, etc")
	errInvalidBrazilianCountryCode  = errors.New("Brazilian mobile numbers should have a brazilian country code (ex: 55)")
	errInvalidBrazilianAreaCode     = errors.New("Brazilian mobile numbers should have a valid brazilian area code (ex: 11, 21, 53, 68, etc)")
	errInvalidBrazilianMobileNumber = errors.New("Brazilian mobile numbers should have a valid number (ex: 9xxxx-xxxx")

	errIncorrectLenghtSusNumber = errors.New("SUS numbers must contain 15 numbers")
	errInvalidSusNumber         = errors.New("SUS number input is not valid")
)
