package main

import (
	"fmt"
	"log"

	brazil "github.com/flavioltonon/go-brazil"
)

func main() {
	var validation brazil.Validation

	// ------------------------------------------------------------------------------------------------

	// Generates a new CPF number in the string format XXX.XXX.XXX-XX
	newCPF := brazil.GenerateCPF()

	// Parses CPF number into a cpf struct
	cpf := brazil.ParseCPF(newCPF)

	// Evaluates a cpf struct number and returns a valid status and a reason in case this status equals false
	validation = brazil.EvaluateCPF(cpf)
	log.Println(fmt.Sprintf("%+v", validation))

	// ------------------------------------------------------------------------------------------------

	// Generates a new PIS number in the string format XXX.XXXXX.XX-X
	newPIS := brazil.GeneratePIS()

	// Parses CPF number into a cpf struct
	pis := brazil.ParsePIS(newPIS)

	// Evaluates a cpf struct number and returns a valid status and a reason in case this status equals false
	validation = brazil.EvaluatePIS(pis)
	log.Println(fmt.Sprintf("%+v", validation))

	// ------------------------------------------------------------------------------------------------

	// Generates a new Título de Eleitor number in the string format XXXXXXXXXXXX
	newTitulo := brazil.GenerateTituloEleitoral()

	// Parses CPF number into a cpf struct
	titulo := brazil.ParseTituloEleitoral(newTitulo)

	// Evaluates a tituloEleitoral struct number and returns a valid status and a reason in case this status equals false
	validation = brazil.EvaluateTituloEleitoral(titulo)
	log.Println(fmt.Sprintf("%+v", validation))

	// ------------------------------------------------------------------------------------------------

	// Generates a new date inside of a chosen range of years in the string format DD/MM/YYYY
	newDate := brazil.GenerateRandomDate(0, 9999)

	// Parses date into a brDate struct
	date := brazil.ParseDate(newDate)

	// Validators - return true when their condition is matched
	log.Println(date.IsFuture())
	log.Println(date.IsToday())
	log.Println(date.IsPast())

	// ------------------------------------------------------------------------------------------------
}
