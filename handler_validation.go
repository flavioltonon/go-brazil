package brazil

import (
	"fmt"
)

type documentType string

const (
	CNPJ             documentType = "cnpj"
	CPF              documentType = "cpf"
	MOBILE           documentType = "mobile"
	PIS              documentType = "pis"
	SUS              documentType = "sus"
	TITULO_ELEITORAL documentType = "tituloEleitoral"
	CERTIDAO         documentType = "certidao"
)

func Validate(t documentType, number string, mask bool) (string, error) {
	switch t {
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

	case TITULO_ELEITORAL:
		titulo, err := ParseTituloEleitoral(number)
		if err != nil {
			return "", err
		}
		return titulo.Number(mask), err

	case CERTIDAO:
		certidao, err := ParseCertidao(number)
		if err != nil {
			return "", err
		}
		return certidao.Number(mask), err

	default:
		return "", fmt.Errorf("%s is not a valid document type", t)
	}
}
