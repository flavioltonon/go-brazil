package main

import (
	"log"

	brazil "flavioltonon/go-brazil"
)

func main() {
	// ------------------------------------------------------------------------------------------------

	// Generates a new string mobile full number
	mobileFullNumber := brazil.RandomMobileFullNumber(false)

	// Creates a new mobile struct
	mobile, err := brazil.ParseMobile(mobileFullNumber)
	if err != nil {
		log.Println(err)
		log.Println(mobileFullNumber)
		return
	}

	// Returns mobile full number
	log.Println(mobile.FullNumber(true))

	// Returns mobile country code
	log.Println(mobile.CountryCode(false))

	// Returns mobile area code
	log.Println(mobile.AreaCode(false))

	// Returns mobile number
	log.Println(mobile.Number(false))

	// ------------------------------------------------------------------------------------------------

	// Generates a new CPF number in the string format XXX.XXX.XXX-XX
	cpfNumber := brazil.RandomCPFNumber(true)

	// Creates a new CPF struct
	cpf, err := brazil.ParseCPF(cpfNumber)
	if err != nil {
		log.Println(err)
		log.Println(cpfNumber)
		return
	}

	// Returns CPF number
	log.Println(cpf.Number(true))

	// ------------------------------------------------------------------------------------------------

	// Generates a new PIS number in the string format XXX.XXXXX.XX-X
	pisNumber := brazil.RandomPISNumber(true)

	// Creates a new PIS struct
	pis, err := brazil.ParsePIS(pisNumber)
	if err != nil {
		log.Println(err)
		log.Println(pisNumber)
		return
	}

	// Returns PIS number
	log.Println(pis.Number(true))

	// ------------------------------------------------------------------------------------------------

	// Generates a new Título Eleitoral number in the string format XXXXXXXXXXXX
	tituloEleitoralNumber := brazil.RandomTituloEleitoralNumber(true)

	// Creates a new TituloEleitoral struct
	titulo, err := brazil.ParseTituloEleitoral(tituloEleitoralNumber)
	if err != nil {
		log.Println(err)
		log.Println(tituloEleitoralNumber)
		return
	}

	// Returns Titulo Eleitoral number
	log.Println(titulo.Number(true))

	// ------------------------------------------------------------------------------------------------

	// Generates a new SUS number in the string format XXXXXXXXXXXX
	susNumber := brazil.RandomSUSNumber(true)

	// Creates a new SUS struct
	sus, err := brazil.ParseSUS(susNumber)
	if err != nil {
		log.Println(err)
		log.Println(susNumber)
		return
	}

	// Returns SUS number
	log.Println(sus.Number(true))

	// ------------------------------------------------------------------------------------------------

	// Generates a new time.Time date inside of a chosen range of years
	newDate, err := brazil.RandomDate(0, 1000)
	if err != nil {
		log.Println(err)
		return
	}

	// Creates a new Date struct
	date, err := brazil.ParseDate(newDate)
	if err != nil {
		log.Println(err)
		log.Println(newDate)
		return
	}

	// Returns date
	log.Println(date.Date())

	// Returns time
	log.Println(date.Time())

	// Validators - return true when their condition is matched
	log.Println(date.IsFuture())
	log.Println(date.IsToday())
	log.Println(date.IsPast())

	// Returns true if date is older than the input date object
	log.Println(date.IsOlderThan(date))

	// Returns true if year input is a leap year
	log.Println(brazil.IsLeapYear(date.Year()))

	// ------------------------------------------------------------------------------------------------
}
