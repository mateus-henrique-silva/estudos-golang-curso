package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var array [2][2]int
	var i int
	var j int

	for i = 0; i < 2; i++ {

		for j = 0; j < 2; j++ {
			reader := bufio.NewReader(os.Stdin)
			input, _ := reader.ReadString('\n')

			// Remove o caractere de nova linha da entrada do usuário
			input = input[:len(input)-1]

			// Converte a entrada do usuário em um valor inteiro
			value, err := strconv.Atoi(input)
			if err != nil {
				fmt.Println("Erro ao converter entrada do usuário em um inteiro:", err)
				return
			}
			array[i][j] = value
		}
	}

	for i = 0; i < 2; i++ {
		for j = 0; j < 2; j++ {
			fmt.Print(" ", array[i][j])
		}

		fmt.Printf("\n")
	}

}
