package brazil

import "errors"

var (
	errIncorrectFormatCpf             = errors.New("CPF format input is different than expected. Expected formats: XXXXXXXXXXX, XXX.XXX.XXX-XX")
	errInvalidCpf                     = errors.New("CPF number input is not valid.")
	errIncorrectFormatTituloEleitoral = errors.New("Título de eleitor format input is different than expected. Expected formats: XXXXXXXXXXXX, XXXXXXXXXXXXXX")
	errInvalidTituloEleitoral         = errors.New("Título de eleitor number input is not valid.")
	errIncorrectFormatPis             = errors.New("PIS format input is different than expected. Expected formats: XXXXXXXXXXX, XXX.XXXXX.XX-X")
	errInvalidPis                     = errors.New("PIS number input is not valid.")
	errIncorrectFormatDate            = errors.New("Date format input is different than expected. Expected format: DD/MM/YYYY")
)
