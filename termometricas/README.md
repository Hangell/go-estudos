# Conversão de Escalas Termométricas

## Desafio da Escola DIO Tech Education

Este projeto é um resultado de um desafio proposto por um professor de ensino médio que quis incentivar seus alunos a desenvolverem habilidades em programação. O objetivo geral é criar um programa em Go que converta o ponto de ebulição da água de Kelvin para Celsius.

### Descrição do Desafio

Um professor de ensino médio solicitou aos seus alunos que desenvolvessem um programa para converter o valor do ponto de ebulição da água de Kelvin para graus Celsius. O ponto de ebulição da água em Kelvin é de aproximadamente 373.15 K, e a conversão pode ser realizada com a fórmula:

\[
C = K - 273.15
\]

### Resolução

O código em Go abaixo faz a conversão necessária:

```go
package main

import "fmt"

func main() {
 pontoDeEbulicaoDaAguaEmKelvin := 373.15

 pontoDeEbulicaoDaAguaEmCelsius := pontoDeEbulicaoDaAguaEmKelvin - 273.15

 fmt.Printf("A temperatura de ebulição da água em Kelvin é %g K, temperatura de ebulição em Celsius é %g ºC", pontoDeEbulicaoDaAguaEmKelvin, pontoDeEbulicaoDaAguaEmCelsius)
}
```

O código começa por definir o ponto de ebulição da água em Kelvin (`373.15`) e, em seguida, utiliza a fórmula para converter esse valor para Celsius. Finalmente, o programa imprime ambas as temperaturas no console.

### Como Executar

1. Certifique-se de que você tenha Go instalado em seu sistema. Se não, você pode baixá-lo [aqui](https://golang.org/dl/).
2. Clone este repositório ou copie o código acima para um arquivo `.go`.
3. Abra um terminal e navegue até o diretório onde o arquivo `.go` está localizado.
4. Execute `go run <nome-do-arquivo>.go`.
5. O resultado da conversão será impresso no console.

### Contribuições

Este é um projeto educacional e todas as contribuições são bem-vindas. Sinta-se à vontade para abrir um issue ou criar um pull request.

---

Desenvolvido por Rodrigo Rangel.
