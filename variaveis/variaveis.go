package main

import "fmt"

func main() {
	var mateus int = 10
	var luis *int = &mateus
	mateus++
	fmt.Println(*luis, mateus)
}
