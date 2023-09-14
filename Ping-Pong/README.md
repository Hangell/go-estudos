# Go Ping-Pong - Concorrência

## Descrição do Projeto

Este projeto é um exemplo simples, mas eficaz, de como utilizar canais e goroutines na linguagem de programação Go. O objetivo é demonstrar a alternância entre as palavras "ping" e "pong" através de concorrência.

## Objetivo Geral

Criar um programa na linguagem Go que trabalhe com dois pacotes e concorrência.

## Desafio

Escreva um código em Go utilizando todo o conhecimento adquirido até o momento! E nesse código você precisará, baseado em nossas aulas de concorrência, utilizar canais e goroutines para que o seu programa exiba, em alternância, as palavras ping e pong.

## Desenvolvido por

Este projeto foi desenvolvido por ***Rodrigo Rangel***.

## Estrutura do Código

O projeto contém três funções principais:

- `ping(c chan string)`: Envia a string "ping" para o canal.
- `pong(c chan string)`: Envia a string "pong" para o canal.
- `printer(c chan string)`: Recebe uma string do canal e imprime no terminal.

Cada uma dessas funções é executada como uma goroutine separada.

## Como Executar

1. Certifique-se de que você tenha Go instalado em seu sistema. Se não, você pode baixá-lo [aqui](https://golang.org/dl/).
2. Clone este repositório ou copie o código acima para um arquivo `.go`.
3. Abra um terminal e navegue até o diretório onde o arquivo `.go` está localizado.
4. Execute `go run <nome-do-arquivo>.go`.
5. O resultado será impresso no console.
