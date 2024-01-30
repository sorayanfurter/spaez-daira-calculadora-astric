package security

import (
	"ASTRIC/BackEnd/ASTRIC/helper/env"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// generarTockens Genera el tocken del usuario
func (security *Security) generarTockens() error {

	clave := []byte(env.GetEnvGeneral().ClaveSecreta)
	expire := env.GetEnvGeneral().TiempoToken

	payload := jwt.MapClaims{
		"data":   security.usuario,
		"expire": time.Now().Add(time.Minute * time.Duration(expire)),
		"exp":    time.Now().Add(time.Minute * time.Duration(expire)).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(clave)
	if err != nil {
		return err

	}

	security.tkn = tokenStr
	return nil
}
