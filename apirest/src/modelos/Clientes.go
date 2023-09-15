package modelos

import "database/sql"

// Cliente representa um cliente utilizando a API
type Cliente struct {
	ID       uint64         `json:"id,omitempty"`
	Nome     string         `json:"nome,omitempty"`
	Email    string         `json:"email,omitempty"`
	Telefone string         `json:"telefone,omitempty"`
	Endereco []Endereco     `json:"endereco,omitempty"`
	CriadoEm sql.NullString `json:"criadoEm,omitempty"`
}

// Endereco representa um endereço de um cliente
type Endereco struct {
	ID          uint64 `json:"id,omitempty"`
	ClienteID   uint64 `json:"clienteId,omitempty"`
	Logradouro  string `json:"logradouro,omitempty"`
	Numero      string `json:"numero,omitempty"`
	Complemento string `json:"complemento,omitempty"`
	Bairro      string `json:"bairro,omitempty"`
	Cidade      string `json:"cidade,omitempty"`
	Estado      string `json:"estado,omitempty"`
	CEP         string `json:"cep,omitempty"`
}
