package routes

import (
	"ASTRIC/BackEnd/ASTRIC/users/controller/usuarios"
	"ASTRIC/BackEnd/shared/routes"
)

// RutasAdmin funcion de rutas
func rutasAdmin(ruta routes.TipoRuta) {

	//Rutas de gestion de usuarios
	ruta("/usuarioAlta", usuarios.Alta, "POST", "Crear nuevos usuarios")
	ruta("/Modificar", usuarios.Modificar, "PUT", "Modificar datos del usuario")
}
