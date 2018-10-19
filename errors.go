package brazil

import "errors"

var (
	errFieldNumberIsRequired                = errors.New("Field 'number' is required.")
	errFieldDateIsRequired                  = errors.New("Field 'date' is required.")
	errIncorrectFormatCpfNumber             = errors.New("CPF format input is different than expected. Expected formats: XXXXXXXXXXX, XXX.XXX.XXX-XX")
	errInvalidCpfNumber                     = errors.New("CPF number input is not valid.")
	errValidCpfNumber                       = errors.New("CPF number input is valid.")
	errIncorrectFormatTituloEleitoralNumber = errors.New("Título de eleitor format input is different than expected. Expected formats: XXXXXXXXXXXX, XXXXXXXXXXXXXX")
	errInvalidTituloEleitoralNumber         = errors.New("Título de eleitor number input is not valid.")
	errValidTituloEleitoralNumber           = errors.New("Título de eleitor number input is valid.")
	errIncorrectFormatPisNumber             = errors.New("PIS format input is different than expected. Expected formats: XXXXXXXXXXX, XXX.XXXXX.XX-X")
	errCorrectFormatPisNumber               = errors.New("PIS format input is correct.")
	errInvalidPisNumber                     = errors.New("PIS number input is not valid.")
	errValidPisNumber                       = errors.New("PIS number input is valid.")
	errFieldDateNotRequired                 = errors.New("Field date is not required.")
	errIncorrectFormatDate                  = errors.New("Date format input is different than expected. Expected format: DD/MM/YYYY")
	errNotPastDate                          = errors.New("Date must be in the past.")
	errInvalidDate                          = errors.New("Date input is not valid.")
	errValidDate                            = errors.New("Date input is valid.")
	errNullDate                             = errors.New("Date has not been input.")
	errIncorrectFormatPhoneNumber           = errors.New("Phone format input is different than expected. Expected formats: XXXXXXXXXXXXX, +XX(XX)XXXXX-XXXX, (XX)XXXXX-XXXX, etc")
	errFieldFullNumberIsRequired            = errors.New("Field 'fullNumber' is required.")
)
