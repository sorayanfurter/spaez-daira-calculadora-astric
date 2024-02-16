package routes

import (
	"ASTRIC/BackEnd/ASTRIC/helper/routes"
)

// RutasModulo implementación de rutas
func RutasModulo(rutastruct routes.RoutersStruct) {
	rutas(rutastruct.Ruta)
	acciones(rutastruct.Accion)
	services(rutastruct.Service)
}
