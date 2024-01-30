package global

import (
	ASTRICmodels "ASTRIC/BackEnd/ASTRIC/pModels"
)

var config ASTRICmodels.RoutesConfig

// SetDisableSucursal setea si la ruta requiere sucursal
func SetDisableSucursal(ruta string) {
	config.Dissucursal = append(config.Dissucursal, ruta)
}

// SetSecurity setea si la ruta requieres seguridad de los roles
func SetSecurity(ruta string) {
	config.RexetnasSeg = append(config.RexetnasSeg, ruta)
}

// SetRuteAccion setea que la reuta es un accion
func SetRuteAccion(ruta string) {
	config.Acciones = append(config.Acciones, ruta)
}

// GetRouteConfig Retorna la confugracion de la ruta
func GetRouteConfig() ASTRICmodels.RoutesConfig {
	return config
}
