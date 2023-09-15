package controllers

import (
	"apirest/src/database"
	"apirest/src/modelos"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// CriarCliente cria um cliente
func CriarCliente(w http.ResponseWriter, r *http.Request) {
	// Decode o corpo da requisição
	var novoCliente modelos.Cliente
	err := json.NewDecoder(r.Body).Decode(&novoCliente)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Pegar o horário atual e formatá-lo como uma string
	criadoEm := time.Now().Format(time.RFC3339) // Ajuste o formato conforme necessário

	// Iniciar uma transação
	db := database.DBCon
	tx, err := db.Begin()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Inserir o cliente
	res, err := tx.Exec("INSERT INTO clientes (nome, email, telefone, criadoEm) VALUES (?, ?, ?, ?)",
		novoCliente.Nome, novoCliente.Email, novoCliente.Telefone, criadoEm)
	if err != nil {
		tx.Rollback()
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	clienteID, err := res.LastInsertId()
	if err != nil {
		tx.Rollback()
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Inserir endereços, se houver
	for _, endereco := range novoCliente.Endereco {
		_, err := tx.Exec("INSERT INTO enderecos (clienteId, logradouro, numero, complemento, bairro, cidade, estado, cep) VALUES (?, ?, ?, ?, ?, ?, ?, ?)",
			clienteID, endereco.Logradouro, endereco.Numero, endereco.Complemento, endereco.Bairro, endereco.Cidade, endereco.Estado, endereco.CEP)
		if err != nil {
			tx.Rollback()
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	// Commit da transação
	if err := tx.Commit(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// mensagem de sucesso
	mensagemSucesso := fmt.Sprintf("Cliente %s criado com sucesso!", novoCliente.Nome)
	resposta := map[string]string{"message": mensagemSucesso}

	// Converter a resposta para JSON
	respostaJSON, err := json.Marshal(resposta)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Enviar a resposta
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	w.Write(respostaJSON)
}

// BuscarClientes busca todos os clientes
func BuscarClientes(w http.ResponseWriter, r *http.Request) {
	db := database.DBCon
	rows, err := db.Query("SELECT * FROM clientes")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var clientes []modelos.Cliente
	for rows.Next() {
		var c modelos.Cliente
		if err := rows.Scan(&c.ID, &c.Nome, &c.Email, &c.Telefone, &c.CriadoEm); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		clientes = append(clientes, c)
	}

	// Verifique se o array está vazio
	if len(clientes) == 0 {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("[]"))
		return
	}

	// Converta o array de clientes para JSON e retorne a resposta
	jsonClientes, err := json.Marshal(clientes)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonClientes)
}

// BuscarCliente busca um cliente pelo ID
func BuscarCliente(w http.ResponseWriter, r *http.Request) {
	//  ID do cliente a partir da URL
	vars := mux.Vars(r)
	id := vars["id"]

	// Conecte-se ao banco de dados
	db := database.DBCon

	// Consulta para buscar um cliente pelo ID
	var cliente modelos.Cliente
	err := db.QueryRow("SELECT id, nome, email, telefone, criadoEm FROM clientes WHERE id = ?", id).Scan(&cliente.ID, &cliente.Nome, &cliente.Email, &cliente.Telefone, &cliente.CriadoEm)
	if err != nil {
		if err == sql.ErrNoRows {
			// Caso em que o registro não foi encontrado
			w.WriteHeader(http.StatusNotFound)
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"message": "Id não encontrado"}`))
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Consulta para buscar endereços do cliente pelo ID
	rows, err := db.Query("SELECT id, clienteId, logradouro, numero, complemento, bairro, cidade, estado, cep FROM enderecos WHERE clienteId = ?", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var endereco modelos.Endereco
		if err := rows.Scan(&endereco.ID, &endereco.ClienteID, &endereco.Logradouro, &endereco.Numero, &endereco.Complemento, &endereco.Bairro, &endereco.Cidade, &endereco.Estado, &endereco.CEP); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		cliente.Endereco = append(cliente.Endereco, endereco)
	}

	// Transforma a struct do cliente em JSON
	clienteJSON, err := json.Marshal(cliente)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Envie a resposta
	w.Header().Set("Content-Type", "application/json")
	w.Write(clienteJSON)
}

// AtualizarCliente atualiza um cliente
func AtualizarCliente(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var atualizarCliente modelos.Cliente
	err := json.NewDecoder(r.Body).Decode(&atualizarCliente)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Iniciar uma transação
	db := database.DBCon
	tx, err := db.Begin()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Atualizar o cliente
	resultado, err := tx.Exec("UPDATE clientes SET nome = ?, email = ?, telefone = ? WHERE id = ?", atualizarCliente.Nome, atualizarCliente.Email, atualizarCliente.Telefone, id)
	if err != nil {
		tx.Rollback()
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Verificar se o cliente foi encontrado e atualizado
	linhasAfetadas, err := resultado.RowsAffected()
	if err != nil {
		tx.Rollback()
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if linhasAfetadas == 0 {
		// Cliente não encontrado, rollback e envie mensagem de erro
		tx.Rollback()
		w.WriteHeader(http.StatusNotFound)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"message": "Id não encontrado"}`))
		return
	}

	// Remover todos os endereços antigos (isso pode ser otimizado dependendo do caso de uso)
	_, err = tx.Exec("DELETE FROM enderecos WHERE clienteId = ?", id)
	if err != nil {
		tx.Rollback()
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Adicionar novos endereços
	for _, endereco := range atualizarCliente.Endereco {
		_, err := tx.Exec("INSERT INTO enderecos (clienteId, logradouro, numero, complemento, bairro, cidade, estado, cep) VALUES (?, ?, ?, ?, ?, ?, ?, ?)",
			id, endereco.Logradouro, endereco.Numero, endereco.Complemento, endereco.Bairro, endereco.Cidade, endereco.Estado, endereco.CEP)
		if err != nil {
			tx.Rollback()
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	// Commit da transação
	if err := tx.Commit(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Criar a mensagem de sucesso
	mensagemSucesso := fmt.Sprintf("Cliente %s atualizado com sucesso!", atualizarCliente.Nome)
	resposta := map[string]string{"message": mensagemSucesso}

	// Converter a resposta para JSON
	respostaJSON, err := json.Marshal(resposta)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Enviar a resposta
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(respostaJSON)
}

// DeletarCliente deleta um cliente
func DeletarCliente(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	// Iniciar uma transação
	db := database.DBCon
	tx, err := db.Begin()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Deleta todos os endereços relacionados ao cliente
	_, err = tx.Exec("DELETE FROM enderecos WHERE clienteId = ?", id)
	if err != nil {
		tx.Rollback()
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Deletar o cliente
	resultado, err := tx.Exec("DELETE FROM clientes WHERE id = ?", id)
	if err != nil {
		tx.Rollback()
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Verificar se o cliente foi encontrado e excluído
	linhasAfetadas, err := resultado.RowsAffected()
	if err != nil {
		tx.Rollback()
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if linhasAfetadas == 0 {
		// Cliente não encontrado, rollback e envie mensagem de erro
		tx.Rollback()
		w.WriteHeader(http.StatusNotFound)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"message": "Id não encontrado"}`))
		return
	}

	// Commit da transação
	if err := tx.Commit(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Criar a mensagem de sucesso
	mensagemSucesso := fmt.Sprintf("Cliente excluído com sucesso")
	resposta := map[string]string{"message": mensagemSucesso}

	// Converter a resposta para JSON
	respostaJSON, err := json.Marshal(resposta)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Enviar a resposta
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(respostaJSON)
}
