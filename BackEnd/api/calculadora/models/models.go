package models

//Estructura Calculadora Request
type CalculadoraRequest struct {
	N1        float64 `json:"n1"`
	N2        float64 `json:"n2"`
	Operacion string  `json:"operacion"`
}

//Estructura Calculadora Response
type CalculadoraResponse struct {
	Resultado float64        `json:"resultado"`
	History   []HistoryEntry `json:"history"`
}

//Estructura entrada historial
type HistoryEntry struct {
	Fecha     string  `json:"fecha"`
	Operacion string  `json:"operacion"`
	Resultado float64 `json:"resultado"`
}
