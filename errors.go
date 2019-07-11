package brazil

import "errors"

var (
	errIncorrectLenghtCertidaoNumber = errors.New("Certidão numbers must contain 30 numbers")
	errInvalidCertidaoYear           = errors.New("Certidão year input is not valid")
	errInvalidCertidaoFirstDigit     = errors.New("Certidão number first digit input is not valid")
	errInvalidCertidaoSecondDigit    = errors.New("Certidão number second digit input is not valid")

	errIncorrectLenghtCNPJNumber = errors.New("CNPJ numbers must contain 14 numbers")
	errInvalidCNPJFirstDigit     = errors.New("CNPJ number first digit input is not valid")
	errInvalidCNPJSecondDigit    = errors.New("CNPJ number second digit input is not valid")

	errIncorrectLenghtCPFNumber = errors.New("CPF numbers must contain 11 numbers")
	errInvalidCPFNumber         = errors.New("CPF number input is not valid")

	errIncorrectLenghtTituloEleitoralNumber = errors.New("Título de eleitor numbers must contain 12 numbers")
	errInvalidTituloEleitoralNumber         = errors.New("Título de eleitor number input is not valid")

	errIncorrectLenghtPISNumber = errors.New("PIS numbers must contain 11 or 13 numbers")
	errInvalidPISNumber         = errors.New("PIS number input is not valid")

	errIncorrectFormatDate = errors.New("Date format input is different than expected. Expected format: DD/MM/YYYY")
	errInvalidYearLimits   = errors.New("Date minYear must be lower than maxYear and higher than zero")
	errNotPastDate         = errors.New("Date must be in the past")
	errInvalidDate         = errors.New("Date input is not valid")
	errNullDate            = errors.New("Date has not been input")

	errIncorrectFormatMobileNumber  = errors.New("Mobile number format input is different than expected. Expected formats: XXXXXXXXXXXXX, +XX(XX)XXXXX-XXXX, (XX)XXXXX-XXXX, etc")
	errInvalidBrazilianCountryCode  = errors.New("Brazilian mobile numbers should have a brazilian country code (ex: 55)")
	errInvalidBrazilianAreaCode     = errors.New("Brazilian mobile numbers should have a valid brazilian area code (ex: 11, 21, 53, 68, etc)")
	errInvalidBrazilianMobileNumber = errors.New("Brazilian mobile numbers should have a valid number (ex: 9xxxx-xxxx")

	errIncorrectLenghtSUSNumber = errors.New("SUS numbers must contain 15 numbers")
	errInvalidSUSNumber         = errors.New("SUS number input is not valid")
)
