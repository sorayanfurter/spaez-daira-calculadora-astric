package controllers

import (
	"ASTRIC/BackEnd/api/calculadora/models"
	"ASTRIC/BackEnd/shared/ep"
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"net/http"
	"time"
)

var history []models.HistoryEntry

func CalculadoraHandler(w http.ResponseWriter, r *http.Request) {

	var request models.CalculadoraRequest

	//Decodifico el body en el request que recibe un puntero.

	json.NewDecoder(r.Body).Decode(&request)
	response := ep.NewResponse("Resultado de calculo", w)

	defer ep.ErrorControlResponse("/calculadora/Calcular", w, r)

	switch request.Operacion {
	case "+", "-", "x", "÷":
		result, err := performCalculation(request.N1, request.N2, request.Operacion)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Guardar en historial
		historyEntry := models.HistoryEntry{
			Fecha:     time.Now().Format("2006-01-02"),
			Operacion: fmt.Sprintf("%.2f %s %.2f", request.N1, request.Operacion, request.N2),
			Resultado: result,
		}
		history = append(history, historyEntry)

		var resultado models.CalculadoraResponse

		resultado.Resultado = result
		resultado.History = history
		response.DatoSend(resultado)

	default:
		response.ErrSend("operación no válida")

	}
}

func performCalculation(n1, n2 float64, operacion string) (float64, error) {
	result := 0.0
	switch operacion {
	case "+":
		result = n1 + n2
	case "-":
		result = n1 - n2
	case "x":
		result = n1 * n2
	case "÷":
		if n2 != 0 {
			result = n1 / n2
		} else {
			// Manejar división por cero con un error
			return 0, errors.New("división por cero")
		}
	default:
		return 0, errors.New("operación no válida")
	}

	// Redondeo a 2 decimales
	result = math.Round(result*100) / 100

	return result, nil
}
