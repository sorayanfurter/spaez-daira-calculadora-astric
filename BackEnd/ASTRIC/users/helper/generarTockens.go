package helper

import (
	"ASTRIC/BackEnd/ASTRIC/helper/env"
	"ASTRIC/BackEnd/config"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// GenerarTockens Genera el tocken del usuario
func GenerarTockens(usuario config.Usuarios) (string, error) {

	clave := []byte(env.GetEnvGeneral().ClaveSecreta)
	expire := env.GetEnvGeneral().TiempoToken

	payload := jwt.MapClaims{
		"data":   usuario,
		"expire": time.Now().Add(time.Minute * time.Duration(expire)),
		"exp":    time.Now().Add(time.Minute * time.Duration(expire)).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(clave)
	if err != nil {
		return "", err

	}
	return tokenStr, nil
}
