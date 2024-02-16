package routes

import (
	"ASTRIC/BackEnd/api/calculadora/controllers"
	"ASTRIC/BackEnd/shared/routes"
)

func rutas(ruta routes.TipoRuta) {
	ruta("/Calcular", controllers.CalculadoraHandler, "POST", "Calcular operacion")
	ruta("/Historial", controllers.HistoryHandler, "GET", "Historial")
}
