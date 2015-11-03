package main

import "fmt"

//START OMIT
func main() {
	valor, status := resultado("iTerior")

	if valor == 1 && status {
		fmt.Println("Ativo")
	} else {
		fmt.Println("Inativo")
	}
}

func resultado(nome string) (int, bool) {
	if nome == "iTerior" {
		return 1, true
	} else {
		return 0, false
	}
}

//STOP OMIT
