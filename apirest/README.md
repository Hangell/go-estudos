# API Rest para Armazenamento de Dados de Clientes

Este é um projeto de API Rest escrito em Go, que permite que a vovó armazene os dados dos seus clientes. Ele utiliza o pacote [Gorilla Mux](https://github.com/gorilla/mux) para gerenciar as rotas HTTP.

## Estrutura de Pastas

O projeto segue a seguinte estrutura de pastas:

```shell
:\apirest
│   bancodedados.db
│   go.mod
│   go.sum
│   main.go
│
└───src
    ├───database
    │       connection.go
    │
    ├───modelos
    │       Clientes.go
    │
    └───router
        │   router.go
        │
        ├───controllers
        │       clientes.go
        │
        └───rotas
                clientes.go
                rotas.go
```

- `bancodedados.db`: O banco de dados SQLite usado para armazenar os dados dos clientes.
- `go.mod` e `go.sum`: Arquivos de controle de dependências do Go.
- `main.go`: O arquivo principal que inicia o servidor da API.
- `src`: Pasta contendo o código-fonte do projeto.
- `database`: Pasta com o código relacionado à conexão com o banco de dados.
- `modelos`: Pasta com as definições das estruturas de dados dos clientes.
- `router`: Pasta com o código responsável por configurar as rotas da API.
- `controllers`: Pasta com os controladores que lidam com as solicitações HTTP.
- `rotas`: Pasta com as definições das rotas da API.

## Pré-requisitos

Antes de executar a API, certifique-se de ter o pacote Gorilla Mux instalado:

```shell
go get -u github.com/gorilla/mux
```

## Executando a API

Para iniciar a API, basta executar o arquivo main.go:

```shell
go run main.go
```

A API estará disponível em http://localhost:3000, a porta configurada no arquivo main.go (definida como default 3000).

## Possiveis Erros

Caso ocorra o erro

1- Não pôde criar a tabela: Binary was compiled with 'CGO_ENABLED=0', go-sqlite3 requires cgo to work. This is a stub
exit status 1

***Windows***
* Na raiz do projeto digite no console
```shell
set CGO_ENABLED=1
```

***Linux ou macOS***
```shell
export CGO_ENABLED=1
```

2 - # runtime/cgo cgo: C compiler "gcc" not found: exec: "gcc": executable file not found in %PATH%

***Windows***
* No powershell
```shell
choco install mingw -y
```

***Linux ou macOS***
```shell
apt-get install build-essential
```

## Endpoints da API

A API oferece os seguintes endpoints:

* [C] [POST](http://localhost:3000/clientes) /clientes: Cria um novo cliente.
* [R] [GET](http://localhost:3000/clientes) /clientes: Retorna todos os clientes.
* [R] [GET](http://localhost:3000/clientes/:id) /clientes/{id}: Retorna um cliente pelo ID.
* [U] [PUT](http://localhost:3000/clientes/:id) /clientes/{id}: Atualiza um cliente pelo ID.
* [D] [DELETE](http://localhost:3000/clientes/:id) /clientes/{id}: Exclui um cliente pelo ID.


## DTO para o Post /clientes

```json
{
  "nome": "Nome Sobrenome",
  "email": "user@email.com",
  "telefone": "+5511123456789",
  "endereco": [
    {
      "logradouro": "Avenida Paulista",
      "numero": "0",
      "complemento": "Apto 1",
      "bairro": "Bela Vista",
      "cidade": "São Paulo",
      "estado": "SP",
      "cep": "01000-000"
    }
  ]
}

```

## Response Get /clientes
```json
[
    {
        "id": 1,
        "nome": "Beltrano da Silva",
        "email": "beltrano@email",
        "telefone": "+5511987654321",
        "criadoEm": {
            "String": "2023-09-14 19:39:25.4527294-03:00",
            "Valid": true
        }
    },
    {
        "id": 2,
        "nome": "Ciclano da Silva",
        "email": "ciclano@email",
        "telefone": "+5511987654321",
        "criadoEm": {
          "String": "2023-09-14 19:39:25.4527294-03:00",
          "Valid": true
        }
    },  
]
```
## Response Get /clientes/{id}
```json
{
  "id": 1,
  "nome": "Beltrano da Silva",
  "email": "beltrano@email.com",
  "telefone": "+5511999999999",
  "endereco": [
    {
      "id": 1,
      "clienteId": 1,
      "logradouro": "Rua dos Bobos",
      "numero": "0",
      "complemento": "casa engraçada",
      "bairro": "Ninguem",
      "cidade": "Nunca",
      "estado": "NG",
      "cep": "01000-000"
    }
  ],
  "criadoEm": {
    "String": "2023-09-14T20:18:28-03:00",
    "Valid": true
  }
}
```
## Body Put /clientes/{id}
```json
{
  "nome": "Beltrano da Silva",
  "email": "beltrano@email.com",
  "telefone": "+5511999999999",
  "endereco": [
    {
      "logradouro": "Rua dos Bobos",
      "numero": "0",
      "complemento": "casa engraçada",
      "bairro": "Ninguem",
      "cidade": "Nunca",
      "estado": "NG",
      "cep": "01000-000"
    }
  ],
  "criadoEm": {
    "String": "2023-09-14T20:18:28-03:00",
    "Valid": true
  }
}
```

## Response Delete /clientes/{id}

```json
{"message":"Cliente excluído com sucesso"}
```

## Licença

Este projeto está licenciado sob a Licença MIT.

## Criado por
* Pix: ***rodrigo@hangell.org***
* Cryptocurrencies or nft MetaMask: <b style="color: green;"><small>0xEd4d1be72F807Faa358C966a8eF63367c200130F</small></b></p>

![Created By](https://media.licdn.com/dms/image/C4D03AQFip79uCsQQhQ/profile-displayphoto-shrink_200_200/0/1606340700270?e=1700092800&v=beta&t=gDe0fvZs8XwfqneqLGa-c8wBsP7CXJPeNDyEQi3WdgY)

<a href="https://www.linkedin.com/in/rodrigo-rangel-a80810170/">Linkedin</a>

<a href="https://play.google.com/store/apps/dev?id=5606456325281613718&hl=pt">PlayStore Hangell</a>
<a href="https://hangell.org">Site Hangell</a>
