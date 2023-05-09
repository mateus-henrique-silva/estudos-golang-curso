package main

import "fmt"

type usuario struct {
	nome     string
	idade    int64
	endereco endereco
}
type endereco struct {
	rua    string
	numero int64
}

func main() {
	fmt.Println("Arquivo Struct")
	var u usuario

	u.nome = "Mateus"
	u.idade = 19

	fmt.Println(u)

	usuario2 := usuario{nome: "mateus", idade: 19}

	fmt.Println(usuario2)
	end := endereco{"rua dos alfeneros", 10}
	usuario3 := usuario{"mateus", 19, end}

	fmt.Println(usuario3)
}
