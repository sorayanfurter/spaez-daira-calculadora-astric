package usuarios

import (
	"ASTRIC/BackEnd/ASTRIC/config/connection/database"
	"ASTRIC/BackEnd/ASTRIC/global"
	"ASTRIC/BackEnd/ASTRIC/helper/env"
	"ASTRIC/BackEnd/ASTRIC/helper/res"
	"ASTRIC/BackEnd/ASTRIC/users/helper"
	"ASTRIC/BackEnd/config"
	"ASTRIC/BackEnd/shared/ep"
	"encoding/json"
	"errors"
	"net/http"
)

// Login realiza el logueo del usuario generando un token
func Login(w http.ResponseWriter, r *http.Request) {
	// swagger:operation POST /admin/usuarioLoguin Usuarios Login
	// ---
	// summary: Ruta de login del usuario
	// parameters:
	// - name: request
	//   in: body
	//   description: Request body
	//   required: true
	//   schema:
	//     "$ref": "#/definitions/Login"
	// responses:
	//   default:
	//     description: Respuesta por defecto
	//     schema:
	//       "$ref": "#/definitions/Response"

	defer ep.ErrorControlResponse("admin/controlleUsuario/Login", w, r)

	var usuario config.Usuarios

	err := json.NewDecoder(r.Body).Decode(&usuario)
	if err != nil {
		res.Response("Login", err.Error(), nil, w)
	}

	err = buscarUsuario(&usuario)
	if err != nil {
		res.ResponseUnauthorized("Login", err.Error(), nil, w)
		return
	}

	tkn, err := helper.GenerarTockens(usuario)
	if err != nil {
		res.Response("Login", err.Error(), nil, w)
	}
	global.SetTKN(tkn)
	res.ResponseLogin("Login", w)
}

func buscarUsuario(usuario *config.Usuarios) error {

	usr := env.GetConfigGeneral()
	env := env.GetEnvGeneral()

	con, cancel := database.SetMysqlCNORM(env.DBASTRIC)
	defer cancel()
	result := con.Where(usr.UserColumnName+" = ? AND "+usr.UserColumnClave+" = ?", usuario.Usuario, usuario.Clave).Find(usuario)
	if result.RowsAffected < 1 {
		return errors.New("nombre de usuario o clave incorrecta")
	}

	result = con.Model(&usuario).Update("hash", helper.GetHash())
	if result.RowsAffected < 1 {
		return errors.New("error al guardar nuevo hash")
	}

	return result.Error

}
