package routes

import (
	"ASTRIC/BackEnd/api/CAMBIARNOMBRE/controllers"
	"ASTRIC/BackEnd/shared/routes"
)

func rutas(ruta routes.TipoRuta) {
	ruta("/prueba", controllers.Prueba, "GET", "Ruta de prueba")
}
