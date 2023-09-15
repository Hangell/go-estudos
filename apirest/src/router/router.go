package router

import (
	"apirest/src/router/rotas"

	"github.com/gorilla/mux"
)

// Vai retornar um ponteiro para um objeto do tipo mux.Router
func Gerar() *mux.Router {
	r := mux.NewRouter()
	return rotas.Configurar(r)
}
