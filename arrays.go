package main

import "fmt"

//START OMIT
func main() {
	var lista [5]int
	fmt.Println("Lista: ", lista)

	lista[3] = 10
	fmt.Println("Lista atualizada: ", lista)
	fmt.Println("lista[3]", lista[3])

	fmt.Println("Tamanho: ", len(lista))

	lista2 := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Lista2[3] ", lista2[3])

	var lista2d [3][3]int
	for i := 0; i < 3; i++ {
		for x := 0; x < 3; x++ {
			lista2d[i][x] = i + x
		}
	}
	fmt.Println(lista2d)
}

//STOP OMIT
