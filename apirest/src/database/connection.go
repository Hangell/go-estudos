package database

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var DBCon *sql.DB

func InitializeDatabase() error {
	// Connect to the SQLite database
	var err error
	DBCon, err = sql.Open("sqlite3", "./bancodedados.db")
	if err != nil {
		return err
	}

	// Create 'clientes' table
	createTableClientes := `
	CREATE TABLE IF NOT EXISTS clientes (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		nome TEXT,
		email TEXT,
		telefone TEXT,
		criadoEm TEXT
	);
	`
	_, err = DBCon.Exec(createTableClientes)
	if err != nil {
		return fmt.Errorf("could not create clientes table: %v", err)
	}

	// Create 'enderecos' table
	createTableEnderecos := `
	CREATE TABLE IF NOT EXISTS enderecos (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		clienteId INTEGER,
		logradouro TEXT,
		numero TEXT,
		complemento TEXT,
		bairro TEXT,
		cidade TEXT,
		estado TEXT,
		cep TEXT,
		FOREIGN KEY (clienteId) REFERENCES clientes (id)
	);
	`
	_, err = DBCon.Exec(createTableEnderecos)
	if err != nil {
		return fmt.Errorf("Não pode criar as tabelas: %v", err)
	}

	fmt.Println("Tabelas criadas com sucesso!")
	return nil
}
