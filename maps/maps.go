package main

import "fmt"

func main() {
	fmt.Println("Maps")

	usuario := map[string]string{
		//chave [] - tipo dos valores
		"nome":  "mateus",
		"idade": "10",
	}

	fmt.Println(usuario["nome"])

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "mateus",
			"ultimo":   "Henrique",
		},
	}
	fmt.Println(usuario2["nome.primeiro"])
}
