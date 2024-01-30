package routes

import "github.com/gorilla/mux"

// NewPrefix Contructor de prefijos de rutas
func NewPrefix(rout *mux.Router, prefix string, description string) *mux.Router {
	return rout.PathPrefix(prefix).Subrouter()
}
