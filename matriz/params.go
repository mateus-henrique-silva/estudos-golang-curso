package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Params(value int) {
	var array [5][5]int
	var i int
	var j int

	for i = 1; i < 5; i++ {
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
		for j = 1; j < 5; j++ {
			array[i][j] = value
		}
	}

	for i = 1; i < 5; i++ {
		for j = 1; j < 5; j++ {
			fmt.Print(" ", array[i][j])
		}

		fmt.Printf("\n")
	}

}
