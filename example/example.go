package main

import (
	brazil "go-brazil"
	"log"
)

func main() {
	log.Println(brazil.EvaluateCPF(brazil.ParseCPF("XXX.XXX.XXX-XX")))

	log.Println(brazil.EvaluateTituloEleitoral(brazil.ParseTituloEleitoral("XXXXXXXXXXXX")))

	log.Println(brazil.EvaluatePIS(brazil.ParsePIS("XXX.XXXXX.XX-XX")))

	date := brazil.ParseDate("DD/MM/YYYY")
	log.Println(date.IsFuture())
	log.Println(date.IsToday())
	log.Println(date.IsPast())

	log.Println(brazil.GenerateCPF())
	log.Println(brazil.EvaluateCPF(brazil.ParseCPF(brazil.GenerateCPF())))

	log.Println(brazil.GeneratePIS())
	log.Println(brazil.EvaluatePIS(brazil.ParsePIS(brazil.GeneratePIS())))

	log.Println(brazil.GenerateTituloEleitoral())
	log.Println(brazil.EvaluateTituloEleitoral(brazil.ParseTituloEleitoral(brazil.GenerateTituloEleitoral())))
}
