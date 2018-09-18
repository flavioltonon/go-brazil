package main

import (
	"fmt"
	"log"

	brazil "github.com/flavioltonon/go-brazil"
)

func main() {
	// ------------------------------------------------------------------------------------------------

	// Creates a new CPF struct
	var cpf brazil.CPF

	// Generates a new CPF number in the string format XXX.XXX.XXX-XX
	cpfNumber := brazil.RandomCPFNumber()

	// Sets cpf number
	cpf.SetNumber(cpfNumber)

	// Returns true if the CPF struct input is valid
	log.Println(fmt.Sprintf("%+v", cpf.IsValid()))

	// Returns all errors associated to the CPF struct
	log.Println(fmt.Sprintf("%+v", cpf.Errors()))

	// ------------------------------------------------------------------------------------------------

	// Generates a new PIS number in the string format XXX.XXXXX.XX-X
	pisNumber := brazil.RandomPISNumber()

	// Sets pis number
	pis := brazil.NewPIS(pisNumber)

	// Returns true if the PIS struct input is valid
	log.Println(fmt.Sprintf("%+v", pis.IsValid()))

	// Returns all errors associated to the PIS struct
	log.Println(fmt.Sprintf("%+v", pis.Errors()))

	// ------------------------------------------------------------------------------------------------

	// Creates a new TituloEleitoral struct
	var titulo brazil.TituloEleitoral
	var erros []error
	var valid bool

	for len(titulo.Errors()) == 0 {
		// Generates a new Título Eleitoral number in the string format XXXXXXXXXXXX
		tituloNumber := brazil.RandomTituloEleitoralNumber()

		// Sets titulo eleitoral number
		titulo.Number(tituloNumber)

		valid = titulo.IsValid()
		erros = titulo.Errors()
		fmt.Println(erros)
	}

	log.Println(fmt.Sprintf("%+v", titulo.GetNumber()))

	// Returns true if the TituloEleitoral struct input is valid
	log.Println(fmt.Sprintf("%+v", valid))

	// Returns all errors associated to the TituloEleitoral struct
	log.Println(fmt.Sprintf("%+v", erros))

	// ------------------------------------------------------------------------------------------------

	var date brazil.BrDate

	// Generates a new time.Time date inside of a chosen range of years
	newDate := brazil.RandomDate(0, 9999)

	// Sets date
	date.SetDate(newDate)
	log.Println(date.GetDate())

	// Validators - return true when their condition is matched
	log.Println(date.IsFuture())
	log.Println(date.IsToday())
	log.Println(date.IsPast())

	// Returns true if year input is a leap year
	log.Println(brazil.IsLeapYear(newDate.Year()))

	// ------------------------------------------------------------------------------------------------
}
