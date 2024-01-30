package usuarios

import (
	"ASTRIC/BackEnd/ASTRIC/config/connection/database"
	"ASTRIC/BackEnd/ASTRIC/helper/env"
	"ASTRIC/BackEnd/config"
	"ASTRIC/BackEnd/shared/ep"
	"errors"

	"encoding/json"
	"net/http"
)

// Alta Realiza la alta de un nuevo usuario
func Alta(w http.ResponseWriter, r *http.Request) {
	// swagger:operation POST /admin/usuarioAlta Usuarios usuarioAlta
	// ---
	// summary: Alta de nuevo usuario
	// parameters:
	// - name: request
	//   in: body
	//   description: Request body
	//   required: true
	//   schema:
	//     "$ref": "#/definitions/Usuario"
	// responses:
	//   default:
	//     description: Respuesta por defecto
	//     schema:
	//       "$ref": "#/definitions/Response"

	defer ep.ErrorControlResponse("admin/controlleUsuario/Alta", w, r)

	var user config.Usuarios

	user.Reset = true

	usr := env.GetConfigGeneral()
	userErr := json.NewDecoder(r.Body).Decode(&user)

	if userErr != nil {
		panic("Error en el formato del json " + userErr.Error())
	}

	erroes, err := ep.ValidateStruct(user)

	if err {
		ep.Response("Alta de usuario", erroes, user, w)
		return
	}

	er := datoExstenteString(user.Usuario, usr.UserColumnName)
	if er != nil {
		ep.Response("Alta de usuario", er.Error(), user, w)
		return
	}

	con, cancel := database.SetMysqlCNORM(env.GetEnvGeneral().DBASTRIC)
	defer cancel()
	result := con.Create(&user)
	if result.RowsAffected < 1 {
		ep.Response("Alta de usuario", "El usuario no pudo crearse", user, w)
		return
	}

	if result.Error != nil {
		ep.Response("Alta de usuario", result.Error, user, w)
		return
	}

	ep.Response("Alta de usuario", nil, user, w)

}

func datoExstenteString(dato string, proiedad string) error {

	con, cancel := database.SetMysqlCNORM(env.GetEnvGeneral().DBASTRIC)
	defer cancel()
	result := con.Where(proiedad+" = ?", dato)
	if result.RowsAffected > 0 {
		return errors.New("el usuario ya existe en la BD")
	}

	return nil
}
