package controllers

import (
	"ASTRIC/BackEnd/api/clientes/models"
	"ASTRIC/BackEnd/shared/ep"
	"net/http"
)

// Prueba esta es una funcion de prueba
func Prueba(w http.ResponseWriter, r *http.Request) {

	defer ep.ErrorControlResponse("/clientes/prueba", w, r)

	res := ep.NewResponse("Prueba de funcionamiento", w)

	var cliene models.Cliente

	cliene.Nombre = "Nombre"
	cliene.Edad = 45
	cliene.Apellido = "Apellido"

	res.DatoSend(cliene)

}
