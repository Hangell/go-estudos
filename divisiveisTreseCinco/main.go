package main

import "fmt"

func main() {
	i := 1
	for i <= 100 {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("Pin Pan")
		} else if i%3 == 0 {
			fmt.Println("Pin")
		} else if i%5 == 0 {
			fmt.Println("Pan")
		} else {
			fmt.Println(i)
		}
		i++
	}

	// Pausa a execução até que o usuário pressione Enter caso seja executado
	fmt.Println("Pressione ENTER para sair...")
	fmt.Scanln()
}
