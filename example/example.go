package main

import (
	"fmt"
	"log"

	brazil "flavioltonon/go-brazil"
)

func main() {
	// ------------------------------------------------------------------------------------------------

	// Generates a new CPF number in the string format +
	phoneNumber := brazil.RandomMobileNumber()

	// Creates a new Phone struct
	mobile := brazil.ParsePhoneNumber(phoneNumber)

	// Returns mobile full number
	log.Println(mobile.FullNumber())

	// Returns mobile country code
	log.Println(mobile.CountryCode())

	// Returns mobile area code
	log.Println(mobile.AreaCode())

	// Returns mobile number
	log.Println(mobile.Number())

	// ------------------------------------------------------------------------------------------------

	// Generates a new CPF number in the string format XXX.XXX.XXX-XX
	cpfNumber := brazil.RandomCPF()

	// Creates a new CPF struct
	cpf := brazil.ParseCPF(cpfNumber)

	// Returns CPF number
	log.Println(cpf.Number())

	// Returns true if the CPF struct input is valid
	log.Println(fmt.Sprintf("%+v", cpf.IsValid()))

	// Returns all errors associated to the CPF struct
	log.Println(fmt.Sprintf("%+v", cpf.Errors()))

	// ------------------------------------------------------------------------------------------------

	// Generates a new PIS number in the string format XXX.XXXXX.XX-X
	pisNumber := brazil.RandomPIS()

	// Creates a new PIS struct
	pis := brazil.ParsePIS(pisNumber)

	// Returns PIS number
	log.Println(pis.Number())

	// Returns true if the PIS struct input is valid
	log.Println(fmt.Sprintf("%+v", pis.IsValid()))

	// Returns all errors associated to the PIS struct
	log.Println(fmt.Sprintf("%+v", pis.Errors()))

	// ------------------------------------------------------------------------------------------------

	// Generates a new Título Eleitoral number in the string format XXXXXXXXXXXX
	tituloNumber := brazil.RandomTituloEleitoral()

	// Creates a new TituloEleitoral struct
	titulo := brazil.ParseTituloEleitoral(tituloNumber)

	// Returns Titulo Eleitoral number
	log.Println(titulo.Number())

	// Returns true if the TituloEleitoral struct input is valid
	log.Println(fmt.Sprintf("%+v", titulo.IsValid()))

	// Returns all errors associated to the TituloEleitoral struct
	log.Println(fmt.Sprintf("%+v", titulo.Errors()))

	// ------------------------------------------------------------------------------------------------

	// Generates a new time.Time date inside of a chosen range of years
	newDate := brazil.RandomDate(0, 9999)

	// Creates a new TituloEleitoral struct
	date := brazil.ParseDate(newDate)

	// Returns date
	log.Println(date.Date())

	// Validators - return true when their condition is matched
	log.Println(date.IsFuture())
	log.Println(date.IsToday())
	log.Println(date.IsPast())

	// Returns true if year input is a leap year
	log.Println(brazil.IsLeapYear(newDate.Year()))

	// ------------------------------------------------------------------------------------------------
}
