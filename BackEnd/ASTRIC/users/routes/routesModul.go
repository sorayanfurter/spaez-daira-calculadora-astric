package routes

import (
	"ASTRIC/BackEnd/ASTRIC/helper/routes"
	"ASTRIC/BackEnd/ASTRIC/middleware"
	"ASTRIC/BackEnd/ASTRIC/users/controller/usuarios"

	"github.com/gorilla/mux"
)

// RutasModul Implementacion de rutas
func RutasModul(rout *mux.Router) {

	//Rutas de Usuario
	ruotAdmin := routes.NewPrefix(rout, "/ASTRIC", "Modulo de gestion de usuarios")
	//ruotersAdmin.Use(database.ChequeBD)
	ruotAdmin.Use(middleware.ProcesarRutas)
	routersAdmin := routes.NewRouter(ruotAdmin)

	rutasAdmin(routersAdmin.Ruta)
	accionesAdmin(routersAdmin.Accion)
	servisesAdmin(routersAdmin.Service)

	//Ruta de login totalmente exenta de seguridad y tokens
	ruotAdmin.HandleFunc("/usuarioLoguin", usuarios.Login).Methods("POST")

}
