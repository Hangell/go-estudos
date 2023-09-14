package main

import "fmt"

func main() {
	i := 1
	for i <= 100 {
		if i%3 == 0 {
			fmt.Printf("O número %v é divisível por 3\n", i)
		}
		i++
	}

	// Pausa a execução até que o usuário pressione Enter caso seja executado
	fmt.Println("Pressione ENTER para sair...")
	fmt.Scanln()
}
