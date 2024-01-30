package usuarios

import (
	"ASTRIC/BackEnd/ASTRIC/config/connection/database"
	"ASTRIC/BackEnd/ASTRIC/helper/env"
	"ASTRIC/BackEnd/config"
	"ASTRIC/BackEnd/shared/ep"
	"encoding/json"
	"net/http"
)

// Modificar Modifica los datos de los usuarios
func Modificar(w http.ResponseWriter, r *http.Request) {
	// swagger:operation PUT /admin/Modificar Usuarios Modificar
	// ---
	// summary: Modificar datos del usuario
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

	defer ep.ErrorControlResponse("admin/controllerUsuario/Modificar", w, r)

	var user config.Usuarios

	user.Reset = true

	userErr := json.NewDecoder(r.Body).Decode(&user)
	if userErr != nil {
		panic("Error en el formato json" + userErr.Error())
	}

	errores, err := ep.ValidateStruct(user)
	if err {
		ep.Response("Modificar datos del usuario", errores, user, w)
		return
	}

	con, cancel := database.SetMysqlCNORM(env.GetEnvGeneral().DBASTRIC)

	result := con.Model(&user).Updates(map[string]interface{}{
		"nombre": user.Nombre, "apellido": user.Apellido, "mail": user.Mail,
		"usuario": user.Usuario, "clave": user.Clave,
		"activo": user.Activo, "admin": user.Admin})

	if result.RowsAffected < 1 {
		ep.Response("Modificar datos del usuario", "Los datos del usuario no pudieron modificarse", user, w)
		return
	}

	if result.Error != nil {
		ep.Response("Modificar datos del usuario", result.Error, user, w)
		return
	}

	defer cancel()

	ep.Response("Modificar datos del usuario", nil, user, w)
}
