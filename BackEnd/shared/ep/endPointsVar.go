package ep

import (
	"ASTRIC/BackEnd/ASTRIC/global"
	"ASTRIC/BackEnd/config"
)

// RequestInfo toda la informacion del usuario
var RequestInfo = global.GetRequestVar

// Usurio Variable global de usuarios
var Usurio config.Usuarios

// SetUsr Setea los datos del usuario
func SetUsr(usr config.Usuarios) {
	Usurio = usr
}
