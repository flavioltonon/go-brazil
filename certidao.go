package brazil

import (
	"math/rand"
	"regexp"
	"strconv"
	"time"
)

// Certidão struct - Validates Certidão de Casamento, Nascimento and Óbito numbers
type certidao struct {
	number certidaoNumber
	valid  bool
}

type certidaoKind string

const (
	Nascimento         certidaoKind = "nascimento"
	Casamento          certidaoKind = "casamento"
	CasamentoReligioso certidaoKind = "casamento-religioso"
	Obito              certidaoKind = "obito"
	Natimorto          certidaoKind = "natimorto"
	Proclamas          certidaoKind = "proclamas" // anúncio necessário antes de qualquer casamento entre noivos de diferentes distritos ou cidades
	Especial           certidaoKind = "especial"  // averbações; nascimentos, casamentos e óbitos ocorridos no exterior
	Emancipacao        certidaoKind = "emancipacao"
	Interdicao         certidaoKind = "interdicao"
	None               certidaoKind = "none"
)

func (c certidaoKind) string() string {
	return string(c)
}

func (c certidaoKind) number() int {
	switch c {
	case Nascimento:
		return 1
	case Casamento:
		return 2
	case CasamentoReligioso:
		return 3
	case Obito:
		return 4
	case Natimorto:
		return 5
	case Proclamas:
		return 6
	case Especial:
		return 7
	case Emancipacao:
		return 8
	case Interdicao:
		return 9
	default:
		return 0
	}
}

func (c certidao) Number(mask bool) string {
	if c.valid && mask {
		return string(c.number[:6]) +
			" " +
			string(c.number[6:8]) +
			" " +
			string(c.number[8:10]) +
			" " +
			string(c.number[10:14]) +
			" " +
			string(c.number[14:15]) +
			" " +
			string(c.number[15:20]) +
			" " +
			string(c.number[20:23]) +
			" " +
			string(c.number[23:30]) +
			"-" +
			string(c.number[30:])
	}
	return string(c.number)
}

func (c certidao) Kind() string {
	switch int(c.number[14]) {
	case Nascimento.number():
		return Nascimento.string()
	case Casamento.number():
		return Casamento.string()
	case CasamentoReligioso.number():
		return CasamentoReligioso.string()
	case Obito.number():
		return Obito.string()
	case Natimorto.number():
		return Natimorto.string()
	case Proclamas.number():
		return Proclamas.string()
	case Especial.number():
		return Especial.string()
	case Emancipacao.number():
		return Emancipacao.string()
	case Interdicao.number():
		return Interdicao.string()
	default:
		return None.string()
	}
}

func ParseCertidao(number string) (certidao, error) {
	number = regexp.MustCompile(`[^0-9]`).ReplaceAllString(number, "")

	if len(number) != 32 {
		return certidao{}, errIncorrectLenghtCertidaoNumber
	}

	certidaoNumber := certidaoNumber(number)

	if !certidaoNumber.hasValidYear() {
		return certidao{}, errInvalidCertidaoYear
	}

	if !certidaoNumber.hasValidFirstDigit() {
		return certidao{}, errInvalidCertidaoFirstDigit
	}

	if !certidaoNumber.hasValidSecondDigit() {
		return certidao{}, errInvalidCertidaoSecondDigit
	}

	return certidao{
		number: certidaoNumber,
		valid:  true,
	}, nil
}

func RandomCertidaoNumber(mask bool, kind certidaoKind) string {
	var multipliers = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	source := rand.NewSource(time.Now().UnixNano())

	r := rand.New(source)

	cns := r.Intn(899999) + 100000
	ca := r.Intn(89) + 10
	tsp := r.Intn(89) + 10
	year := r.Intn(time.Now().Year()-2010) + 2010
	tl := r.Intn(8) + 1
	nl := r.Intn(89999) + 10000
	nf := r.Intn(899) + 100
	nt := r.Intn(8999999) + 1000000

	if kind != None {
		tl = kind.number()
	}

	certidaoString := strconv.Itoa(cns) +
		strconv.Itoa(ca) +
		strconv.Itoa(tsp) +
		strconv.Itoa(year) +
		strconv.Itoa(tl) +
		strconv.Itoa(nl) +
		strconv.Itoa(nf) +
		strconv.Itoa(nt)

	// Calculate first digit
	sum := 0
	for i := 0; i < 30; i++ {
		number, _ := strconv.Atoi(string(certidaoString[i]))
		sum += number * multipliers[i+1]
	}
	firstDigit := sum % 11
	if sum%11 == 10 {
		firstDigit = 1
	}

	// Calculate second digit
	sum = 0
	for i := 0; i < 30; i++ {
		number, _ := strconv.Atoi(string(certidaoString[i]))
		sum += number * multipliers[i]
	}
	sum += firstDigit * multipliers[30]
	secondDigit := sum % 11
	if sum%11 == 10 {
		secondDigit = 1
	}

	if mask {
		return string(certidaoString[:6]) +
			" " +
			string(certidaoString[6:8]) +
			" " +
			string(certidaoString[8:10]) +
			" " +
			string(certidaoString[10:14]) +
			" " +
			string(certidaoString[14:15]) +
			" " +
			string(certidaoString[15:20]) +
			" " +
			string(certidaoString[20:23]) +
			" " +
			string(certidaoString[23:30]) +
			"-" +
			strconv.Itoa(firstDigit) +
			strconv.Itoa(secondDigit)
	}
	return certidaoString + strconv.Itoa(firstDigit) + strconv.Itoa(secondDigit)
}

type certidaoNumber string

func (c certidaoNumber) hasValidYear() bool {
	year, _ := strconv.Atoi(string(c[10:14]))
	return year >= 2010 && year <= time.Now().Year()
}

func (c certidaoNumber) hasValidFirstDigit() bool {
	var (
		multipliers = []int{2, 3, 4, 5, 6, 7, 8, 9, 10, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
		sum         int
	)

	for i := 0; i < 30; i++ {
		certidaoDigit, _ := strconv.Atoi(string(c[i]))
		sum += certidaoDigit * multipliers[i]
	}

	remainder := sum % 11
	if remainder == 10 {
		return string(c[30]) == strconv.Itoa(1)
	}

	return string(c[30]) == strconv.Itoa(remainder)
}

func (c certidaoNumber) hasValidSecondDigit() bool {
	var (
		multipliers = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
		sum         int
	)

	for i := 0; i < 31; i++ {
		certidaoDigit, _ := strconv.Atoi(string(c[i]))
		sum += certidaoDigit * multipliers[i]
	}

	remainder := sum % 11
	if remainder == 10 {
		return string(c[31]) == strconv.Itoa(1)
	}

	return string(c[31]) == strconv.Itoa(remainder)
}
