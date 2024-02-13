package routers

import (
	"ASTRIC/BackEnd/ASTRIC/helper/routes"
	"ASTRIC/BackEnd/ASTRIC/middleware"
	rClientes "ASTRIC/BackEnd/api/clientes/routes"

	"github.com/gorilla/mux"
)

// RutasPrincipales Manejador de rutas principales, donde se declaran los prefijos
func RutasPrincipales(rout *mux.Router) {

	routClintes := routes.NewPrefix(rout, "/clientes", "Modulo de Clientes")
	routClintes.Use(middleware.ProcesarRutas)
	routersInformes := routes.NewRouter(routClintes)
	rClientes.RutasModulo(routersInformes)

}
