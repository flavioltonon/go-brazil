package brazil

import "time"

// CPF struct
type cpf struct {
	number number
	valid  bool
}

// Título de eleitor struct
type tituloEleitoral struct {
	number number
	valid  bool
}

// PIS struct
type pis struct {
	number number
	valid  bool
}

// Phone struct
type phone struct {
	fullNumber  number
	countryCode string
	areaCode    string
	number      string
	valid       bool
}

type number struct {
	number     string
	validation validation
}

type validation struct {
	valid  bool
	reason error
}

// Date struct
type date struct {
	date       time.Time
	validation validation
	notNull    bool
}

type document interface {
	hasExpectedFormat() bool
	isvalid() bool
}
