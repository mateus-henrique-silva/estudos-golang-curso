package main

import "fmt"

func main() {
	var array [5]int
	var slice []int
	var x = 1
	array = [5]int{0, 1, 2, 3, x} //[...] tamanho variavel de acordo com os elmentos recebido;
	slice = []int{5, 3, 5, 4}
	fmt.Println(array[1:3], slice)

	slice3 := make([]float32, 10, 11)
	fmt.Println(slice3)
	slice3 = append(slice3, 10)
	slice3 = append(slice3, 20)
	fmt.Println(cap(slice3))
	fmt.Println(len(slice3))
	fmt.Println(array)

}
