package main

import "fmt"

func Day(numero int) string {
	var diadasemana string
	switch {
	case numero == 1:
		diadasemana = "Domingo"
	case numero == 2:
		diadasemana = "Segunda"
	case numero == 3:
		diadasemana = "Terça"
	case numero == 4:
		diadasemana = "Quarta"
	case numero == 5:
		diadasemana = "Quinta"
	case numero == 6:
		diadasemana = "Sexta"
	case numero == 7:
		diadasemana = "Sabado"

	default:
		diadasemana = "numero invalido"
	}

	return diadasemana

}
func main() {
	fmt.Println("switch")
	dia := Day(2)
	fmt.Println(dia)

}
