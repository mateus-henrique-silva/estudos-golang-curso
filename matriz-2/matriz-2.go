package main

import "fmt"

func main() {
	var array [3][2]int

	var i int
	var j int

	for i = 0; i < 3; i++ {
		for j = 0; j < 2; j++ {
			array[i][j] = (i + j)
		}
	}
	for i = 0; i < 3; i++ {
		for j = 0; j < 2; j++ {
			fmt.Print(" ", array[i][j])

		}
		fmt.Print("\n")
	}
	fmt.Print("\n")
	fmt.Print("Matriz invertida")
	for i = 0; i < 3; i++ {
		for j = 0; j < 2; j++ {
			fmt.Print(" ", array[i][j])

		}
		fmt.Print("\n")
	}

}
