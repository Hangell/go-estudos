package main

import "testing"

func TestSumCorrect(t *testing.T) {
	result := calculadora(soma, 10, 10)
	if result != 20 {
		t.Errorf("Resultado inválido: %d. Esperado: %d", result, 20)
	}
}

func TestSumIncorrect(t *testing.T) {
	result := calculadora(soma, 10, 8)
	if result != 20 {
		t.Errorf("Resultado inválido: %d. Esperado: %d", result, 20)
	}
}

func TestSubtractCorrect(t *testing.T) {
	result := calculadora(subtracao, 10, 2)
	if result != 8 {
		t.Errorf("Resultado inválido: %d. Esperado: %d", result, 8)
	}
}

func TestSubtractIncorrect(t *testing.T) {
	result := calculadora(subtracao, 10, 3)
	if result != 8 {
		t.Errorf("Resultado inválido: %d. Esperado: %d", result, 8)
	}
}

func TestMultiplyCorrect(t *testing.T) {
	result := calculadora(multiplicacao, 10, 10)
	if result != 100 {
		t.Errorf("Multiplicação incorreta, resultado: %d, esperado: %d.", result, 100)
	}
}

func TestMultiplyIncorrect(t *testing.T) {
	result := calculadora(multiplicacao, 10, 8)
	if result != 100 {
		t.Errorf("Multiplicação incorreta, resultado: %d, esperado: %d.", result, 100)
	}
}

func TestDivideCorrect(t *testing.T) {
	result := calculadora(divisao, 10, 2)
	if result != 5 {
		t.Errorf("Divisão incorreta, resultado: %d, esperado: %d.", result, 5)
	}
}

func TestDivideIncorrect(t *testing.T) {
	result := calculadora(divisao, 10, 3)
	if result != 5 {
		t.Errorf("Divisão incorreta, resultado: %d, esperado: %d.", result, 5)
	}
}
