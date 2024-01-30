package routes

import (
	"ASTRIC/BackEnd/ASTRIC/helper/routes"
	"ASTRIC/BackEnd/ASTRIC/users/controller/usuarios"
)

// RutasAdmin funcion de rutas
func servisesAdmin(service routes.TipoService) {
	//service("/usuarioGetInfo", usuarios.GetInfo, "GET")

	//Listado de usuarios
	service("/Listado", usuarios.Listado, "GET")
	service("/FiltrarDatosUsuario/{ID}", usuarios.FiltrarDatosUsuario, "GET")
}
