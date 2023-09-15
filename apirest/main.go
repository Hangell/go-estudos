package main

import (
	"apirest/src/database"
	"apirest/src/router"
	"log"
	"net/http"
)

func main() {
	// Inicializa o banco de dados
	err := database.InitializeDatabase()
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}

	r := router.Gerar()

	log.Fatal(http.ListenAndServe(":3000", r))
}
