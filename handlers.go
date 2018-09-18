package brazil

import "time"

// CPF struct
type CPF struct {
	number    Number
	fullName  string
	birthdate BrDate
	valid     bool
}

// Título de eleitor struct
type TituloEleitoral struct {
	number    Number
	zone      string
	section   string
	state     string
	city      string
	issueDate BrDate
	valid     bool
}

// PIS struct
type PIS struct {
	number       Number
	registration BrDate
	valid        bool
}

// Date struct
type BrDate struct {
	date       time.Time
	validation Validation
	notNull    bool
}

type Number struct {
	number     string
	validation Validation
}

type document interface {
	hasExpectedFormat() bool
	isValid() bool
}
