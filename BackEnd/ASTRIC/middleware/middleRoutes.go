package middleware

import (
	//"ASTRIC/BackEnd/ASTRIC/config/connection/database"

	dbCon "ASTRIC/BackEnd/ASTRIC/config/connection/database"
	"ASTRIC/BackEnd/ASTRIC/global"
	"ASTRIC/BackEnd/ASTRIC/helper/env"
	"ASTRIC/BackEnd/ASTRIC/helper/res"
	"ASTRIC/BackEnd/ASTRIC/helper/routes"
	dbSet "ASTRIC/BackEnd/ASTRIC/middleware/database"
	"ASTRIC/BackEnd/ASTRIC/middleware/security"
	"ASTRIC/BackEnd/config"
	"ASTRIC/BackEnd/shared/ep"
	"errors"
	"net/http"
)

type procesarRutas struct {
	next http.Handler
	w    http.ResponseWriter
	r    *http.Request
}

// ProcesarRutas Middleware de autorizacion de las rutas
func ProcesarRutas(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		routes := routes.NewRoutDesc(r)

		global.SetAppResponceWS(routes.App)
		global.SetEndPointResponceWS(routes.Modulo)

		//Ruta de login, tiene toda la seguridad anulada
		if routes.Ruta == "/ASTRIC/usuarioLoguin" {
			next.ServeHTTP(w, r)
			return
		}

		//Obtengo la bd a trabajar
		result := r.Header.Values("Database")
		if len(result) < 1 {
			res.Response("Seteo de BD", "Falta el parametor Database en el header", nil, w)
			return
		}

		db := string(result[0])
		dbSet.SetDb(db)

		rute := procesarRutas{next: next, r: r, w: w}

		if env.GetEnvGeneral().DisableToken {
			rute.disableTokn()
			return
		}

		auth := security.Security{R: r}
		err := auth.Autorisations()
		if err != nil {
			res.ResponseExpire("Atuorizacion", rute.w, err.Error())
			return
		}
		global.SetTKN(auth.GetTKN())
		next.ServeHTTP(w, r)
	})
}

func (p *procesarRutas) disableTokn() {

	result := p.r.Header.Values("Authorization")
	if len(result) < 1 {
		res.ResponseUnauthorizedMSG("Validacion de TOKEN", "Falta el parametor Authorization en el header", nil, p.w, "La seguridad de TOKEN esta desactivada debe enviar el usuario en ves del tocken")
		return
	}
	usr := string(result[0])
	err := p.setUserGlobal(usr)
	if err != nil {
		res.ResponseUnauthorizedMSG("Validacion de TOKEN", err.Error(), nil, p.w, "La seguridad de TOKEN esta desactivada debe enviar el usuario en ves del tocken")
		return
	}
	global.SetTKN("")
	p.next.ServeHTTP(p.w, p.r)
}

func (p *procesarRutas) setUserGlobal(usr string) error {

	var usuario config.Usuarios

	con, _ := dbCon.SetMysqlCNORM(env.GetEnvGeneral().DBASTRIC)
	result := con.Where(env.GetConfigGeneral().UserColumnName+" = ?", usr).Find(&usuario)

	if result.RowsAffected < 1 {
		return errors.New("usuario no encontrado en la BD")
	}

	ep.SetUsr(usuario)
	global.SetUsuarioResponceWS(usuario.Usuario)
	return result.Error
}
