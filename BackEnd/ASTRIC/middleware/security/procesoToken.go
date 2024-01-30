package security

import (
	"ASTRIC/BackEnd/ASTRIC/config/connection/database"
	"ASTRIC/BackEnd/ASTRIC/helper/env"
	e "ASTRIC/BackEnd/ASTRIC/helper/errors"
	"ASTRIC/BackEnd/config"
	"ASTRIC/BackEnd/shared/ep"
	"errors"

	"github.com/dgrijalva/jwt-go"
)

// Clain modelo de desencriptado del tocken
type Clain struct {
	Data config.Usuarios
	jwt.StandardClaims
}

/*
DesencriptarToken Funcion para desencriptar token
@tk: token en formato string
Returns:

	Clains: modelo de token completado
	error: errores producidos
*/
func (security *Security) DesencriptarToken() error {

	defer e.ErrorControl("Desencripacion de TOKEN")

	claims := &Clain{}
	miCleve := []byte(env.GetEnvGeneral().ClaveSecreta)

	//Valido el tkn y lo desencripto
	_, err := jwt.ParseWithClaims(security.tkn, claims, func(token *jwt.Token) (interface{}, error) {
		return miCleve, nil
	})
	if err != nil {
		return err
	}
	//Valido el hash y obtengo el usuario de la bd
	err = security.getUsr(*claims)
	if err != nil {
		return err
	}
	//Renuevo el tkn
	err = security.generarTockens()
	if err != nil {
		return err
	}

	return nil
}

func (security *Security) getUsr(datos Clain) error {

	con, _ := database.SetMysqlCNORM(env.GetEnvGeneral().DBASTRIC)
	msg := "El id o hash son invalidos o el hash cambio"

	result := con.Where("id = ? AND hash = ?", datos.Data.ID, datos.Data.Hash).Find(&security.usuario)

	if result.RowsAffected < 1 {
		if env.GetEnvGeneral().Produccion {
			msg = "token invalido"
		}
		return errors.New(msg)
	}

	ep.SetUsr(security.usuario)
	return result.Error
}
