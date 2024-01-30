package security

import (
	"ASTRIC/BackEnd/config"
	"errors"
	"net/http"
)

// Security Procesa el tkn
type Security struct {
	R       *http.Request
	usuario config.Usuarios
	tkn     string
}

// GetTKN retorna el token renovado
func (security *Security) GetTKN() string {
	return security.tkn
}

// Autorisations Middleware de autorizacion de las rutas
func (security *Security) Autorisations() error {

	//Obtengo el tkn
	err := security.getToken()
	if err != nil {
		return errors.New("obtener tkn: " + err.Error())
	}
	//Desencripto y valido el tkn y renueva el tkn
	err = security.DesencriptarToken()
	if err != nil {
		return errors.New("desencriptar tkn: " + err.Error())
	}
	return nil
}

// Obtiene el token del reques
func (security *Security) getToken() error {
	result := security.R.Header.Values("Authorization")
	if len(result) < 1 {
		return errors.New("falta paramtetro 'Authorization: token' en el header del reques")
	}
	security.tkn = string(result[0])
	return nil
}
