package main

import "fmt"

func main() {

	err := ""
	fmt.Println("if and else")

	if err != "" {
		fmt.Println(err)
	} else {
		fmt.Println("sem error")
	}

	numero := 10

	if outronumero := numero; outronumero > 0 {
		fmt.Println("o numero é positivo")
	}

}
