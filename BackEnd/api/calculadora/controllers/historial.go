package controllers

import (
	"ASTRIC/BackEnd/shared/ep"
	"net/http"
)

func HistoryHandler(w http.ResponseWriter, r *http.Request) {

	defer ep.ErrorControlResponse("/calculadora/Historial", w, r)
	response := ep.NewResponse("Historial operaciones", w)
	response.DatoSend(history)
}
