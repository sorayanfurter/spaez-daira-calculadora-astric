package models

// Cliente estructura de cliente
type Cliente struct {
	Nombre   string `json:"cliente_nombre"`
	Apellido string `json:"cliente_apellido"`
	Edad     int    `json:"cliente_edad"`
}
