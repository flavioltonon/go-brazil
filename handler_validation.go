package brazil

import (
	"fmt"
)

type DocumentType string

const (
	CNPJ            DocumentType = "cnpj"
	CPF             DocumentType = "cpf"
	MOBILE          DocumentType = "mobile"
	PIS             DocumentType = "pis"
	SUS             DocumentType = "sus"
	TITULOELEITORAL DocumentType = "tituloEleitoral"
)

func Validate(docType DocumentType, number string, mask bool) (string, error) {

	switch docType {
	case CNPJ:
		cnpj, err := ParseCNPJ(number)
		if err != nil {
			return "", err
		}
		return cnpj.Number(mask), err

	case CPF:
		cpf, err := ParseCPF(number)
		if err != nil {
			return "", err
		}
		return cpf.Number(mask), err

	case MOBILE:
		mobile, err := ParseMobile(number)
		if err != nil {
			return "", err
		}
		return mobile.FullNumber(mask), err

	case PIS:
		pis, err := ParsePIS(number)
		if err != nil {
			return "", err
		}
		return pis.Number(mask), err

	case SUS:
		sus, err := ParseSUS(number)
		if err != nil {
			return "", err
		}
		return sus.Number(mask), err

	case TITULOELEITORAL:
		titulo, err := ParseTituloEleitoral(number)
		if err != nil {
			return "", err
		}
		return titulo.Number(mask), err

	default:
		return "", fmt.Errorf("%s is not a valid docType", docType)
	}
}
