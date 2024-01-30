package routes

import "ASTRIC/BackEnd/ASTRIC/helper/routes"

//RutasModul Implementacion de rutas
func RutasModul(rutastruct routes.RoutersStruct) {

	rutas(rutastruct.Ruta)
	acciones(rutastruct.Accion)
	services(rutastruct.Service)

}
