package main

import "fmt"

func main() {
	pontoDeEbulicaoDaAguaEmKelvin := 373.15

	pontoDeEbulicaoDaAguaEmCelsius := pontoDeEbulicaoDaAguaEmKelvin - 273.15
	// Exibir o resultado
	fmt.Printf("A temperatura de ebulição da água em Kelvin é %g K, temperatura de ebulição em Celsius é %g ºC", pontoDeEbulicaoDaAguaEmKelvin, pontoDeEbulicaoDaAguaEmCelsius)
}
