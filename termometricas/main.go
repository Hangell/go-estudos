package main

import "fmt"

func main() {
	pontoDeEbulicaoDaAguaEmKelvin := 373.15

	pontoDeEbulicaoDaAguaEmCelsius := pontoDeEbulicaoDaAguaEmKelvin - 273.15

	fmt.Printf("A temperatura de ebulição da água em Kelvin é %g K, temperatura de ebulição em Celsius é %g ºC\n", pontoDeEbulicaoDaAguaEmKelvin, pontoDeEbulicaoDaAguaEmCelsius)

	// Pausa a execução até que o usuário pressione Enter
	fmt.Println("Pressione ENTER para sair...")
	fmt.Scanln()
}
