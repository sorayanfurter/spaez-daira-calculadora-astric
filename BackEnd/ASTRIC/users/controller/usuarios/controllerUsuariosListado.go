package usuarios

import (
	"ASTRIC/BackEnd/ASTRIC/config/connection/database"
	"ASTRIC/BackEnd/ASTRIC/helper/env"
	"ASTRIC/BackEnd/config"
	"ASTRIC/BackEnd/shared/ep"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Listado Lista los datos de los usuarios registrados
func Listado(w http.ResponseWriter, r *http.Request) {
	// swagger:operation GET /admin/usuariosListado Usuarios usuariosListado
	// ---
	// summary: Listado de usuarios
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

	defer ep.ErrorControlResponse("admin/controllerUsuario/Listado", w, r)

	var user []config.Usuarios

	env := env.GetEnvGeneral()

	con, cancel := database.SetMysqlCNORM(env.DBASTRIC)

	result := con.Find(&user)
	if result.RowsAffected < 1 {
		ep.Response("Listado de usuarios", "No se encontraron usuarios activos", user, w)
		return
	}

	if result.Error != nil {
		ep.Response("Listado de usuarios", result.Error, user, w)
		return
	}

	defer cancel()

	ep.Response("Listado de usuarios", nil, user, w)
}

// FiltrarDatosUsuario Filtra los datos del usuario por ID
func FiltrarDatosUsuario(w http.ResponseWriter, r *http.Request) {
	// swagger:operation GET /admin/usuariosListado Usuarios FiltrarDatosUsuario
	// ---
	// summary: Lista los datos del usuario por ID
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

	defer ep.ErrorControlResponse("admin/controllerUsuario/FiltrarDatosUsuario", w, r)

	var user config.Usuarios

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["ID"])

	if err != nil {
		ep.Response("Listar datos del usuario", err.Error(), user, w)
		return
	}

	env := env.GetEnvGeneral()

	con, cancel := database.SetMysqlCNORM(env.DBASTRIC)

	result := con.Where("id = ?", id).Find(&user)
	if result.RowsAffected < 1 {
		ep.Response("Listar datos del usuario", "No se encontraron usuarios con el ID asignado", user, w)
		return
	}

	if result.Error != nil {
		ep.Response("Listar datos del usuario", result.Error, user, w)
		return
	}

	defer cancel()

	ep.Response("Listar datos del usuario", nil, user, w)
}
