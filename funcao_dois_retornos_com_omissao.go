package main

import "fmt"

//START OMIT
func main() {
	_, status := resultado("iTerior")

	if status {
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
