package routes

import (
	"ASTRIC/BackEnd/ASTRIC/global"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type endpoint func(http.ResponseWriter, *http.Request)

// TipoRuta tipo ruta
type TipoRuta func(string, endpoint, string, string)

// TipoAccion tipo accion
type TipoAccion func(string, endpoint, string, string)

// TipoService tipo service
type TipoService func(string, endpoint, string)

// RoutersStruct estructura de rutas
type RoutersStruct struct {
	Router *mux.Router
}

// NewRouter Contructor de rutas
func NewRouter(rout *mux.Router) RoutersStruct {
	return RoutersStruct{Router: rout}
}

/*
Ruta Crea una ruta: Las RUTAS requieresn bd y poseen seguridad
Las mismas comprenden el menu del usuario
*/
func (r *RoutersStruct) Ruta(ruta string, endpoint endpoint, method string, description string) {
	r.Router.HandleFunc(ruta, endpoint).Methods(method)
}

/*
Accion crea una accion: Las acciones son rutas llamadas generalmente por
botosnes, estas no impactan en el menu del usuario
*/
func (r *RoutersStruct) Accion(ruta string, endpoint endpoint, method string, description string) {
	r.Router.HandleFunc(SetRuteAccion(ruta), endpoint).Methods(method)
}

/*
Service crea un service: estos son rutas sin seguridad solamente requieren que
el usuario este logeado
*/
func (r *RoutersStruct) Service(ruta string, endpoint endpoint, method string) {
	r.Router.HandleFunc(DisableSecurity(ruta), endpoint).Methods(method)
}

/*
DisableSecurity Desactiva la seguridad de la ruta
@ruta: la ruta propiamente dicha
*/
func DisableSecurity(ruta string) string {
	global.SetSecurity(strings.TrimLeft(ruta, "/"))
	return ruta
}

/*
DisableSucrsal desactiva el requerimiento de la sucursal en la ruta
@ruta: La direccion de la ruta propiamente dicha
*/
func DisableSucrsal(ruta string) string {
	global.SetDisableSucursal(strings.TrimLeft(ruta, "/"))
	return ruta
}

/*
SetRuteAccion seta la ruta en modo accion para que no impacte en
el manu y cambie su mensage
@ruta: la ruta propiamente dicha
*/
func SetRuteAccion(ruta string) string {
	global.SetRuteAccion(strings.TrimLeft(ruta, "/"))
	return ruta
}
