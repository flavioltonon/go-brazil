package brazil

import "errors"

var (
	ErrIncorrectLenghtCertidaoNumber = errors.New("Certidão numbers must contain 32 numbers")
	ErrInvalidCertidaoYear           = errors.New("Certidão year input is not valid")
	ErrInvalidCertidaoFirstDigit     = errors.New("Certidão number first digit input is not valid")
	ErrInvalidCertidaoSecondDigit    = errors.New("Certidão number second digit input is not valid")

	ErrIncorrectLenghtCNPJNumber = errors.New("CNPJ numbers must contain 14 numbers")
	ErrInvalidCNPJFirstDigit     = errors.New("CNPJ number first digit input is not valid")
	ErrInvalidCNPJSecondDigit    = errors.New("CNPJ number second digit input is not valid")

	ErrIncorrectLenghtCPFNumber = errors.New("CPF numbers must contain 11 numbers")
	ErrInvalidCPFNumber         = errors.New("CPF number input is not valid")

	ErrIncorrectLenghtTituloEleitoralNumber = errors.New("Título de eleitor numbers must contain 12 numbers")
	ErrInvalidTituloEleitoralNumber         = errors.New("Título de eleitor number input is not valid")

	ErrIncorrectLenghtPISNumber = errors.New("PIS numbers must contain 11 or 13 numbers")
	ErrInvalidPISNumber         = errors.New("PIS number input is not valid")

	ErrIncorrectFormatDate = errors.New("Date format input is different than expected. Expected format: DD/MM/YYYY")
	ErrInvalidYearLimits   = errors.New("Date minYear must be lower than maxYear and higher than zero")
	ErrNotPastDate         = errors.New("Date must be in the past")
	ErrInvalidDate         = errors.New("Date input is not valid")
	ErrNullDate            = errors.New("Date has not been input")

	ErrIncorrectFormatMobileNumber  = errors.New("Mobile number format input is different than expected. Expected formats: XXXXXXXXXXXXX, +XX(XX)XXXXX-XXXX, (XX)XXXXX-XXXX, etc")
	ErrInvalidBrazilianCountryCode  = errors.New("Brazilian mobile numbers should have a brazilian country code (ex: 55)")
	ErrInvalidBrazilianAreaCode     = errors.New("Brazilian mobile numbers should have a valid brazilian area code (ex: 11, 21, 53, 68, etc)")
	ErrInvalidBrazilianMobileNumber = errors.New("Brazilian mobile numbers should have a valid number (ex: 9xxxx-xxxx")

	ErrIncorrectLenghtSUSNumber = errors.New("SUS numbers must contain 15 numbers")
	ErrInvalidSUSNumber         = errors.New("SUS number input is not valid")
)
