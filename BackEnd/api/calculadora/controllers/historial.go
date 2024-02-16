package controllers

import (
	"ASTRIC/BackEnd/api/calculadora/models"
	"ASTRIC/BackEnd/shared/ep"
	"net/http"
)

func HistoryHandler(w http.ResponseWriter, r *http.Request) {

	defer ep.ErrorControlResponse("/calculadora/Historial", w, r)
	response := ep.NewResponse("Historial operaciones", w)
	var history []models.HistoryEntry
	response.DatoSend(history)
}
