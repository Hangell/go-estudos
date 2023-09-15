package main

import "fmt"

func main() {
	resultado := calculadora(soma, 10, 10)
	fmt.Println(resultado)
	resultado = calculadora(subtracao, 10, 2)
	fmt.Println(resultado)
	resultado = calculadora(multiplicacao, 10, 10)
	fmt.Println(resultado)
	resultado = calculadora(divisao, 10, 2)
	fmt.Println(resultado)
}

func soma(i ...int) int {
	total := 0
	for _, v := range i {
		total += v
	}
	return total
}

func subtracao(i ...int) int {
	// Inicie com o primeiro número
	total := i[0]
	for _, v := range i[1:] {
		total -= v
	}
	return total
}

func multiplicacao(i ...int) int {
	total := 1
	for _, v := range i {
		total *= v
	}
	return total
}

func divisao(i ...int) int {
	if len(i) == 0 {
		return 0
	}
	total := i[0]
	if total == 0 {
		return 0 // Evitar divisão por zero
	}
	for _, v := range i[1:] {
		total /= v
	}
	return total
}

func calculadora(funcao func(i ...int) int, numeros ...int) int {
	return funcao(numeros...)
}
