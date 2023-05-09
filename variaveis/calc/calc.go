package calc

import "fmt"

func Calcular(nome string, number1 int, number2 int) {
	total := number1 + number2
	fmt.Println(nome, total)
	go calcularTest()

}
func calcularTest() {
	fmt.Println("teste")
}
