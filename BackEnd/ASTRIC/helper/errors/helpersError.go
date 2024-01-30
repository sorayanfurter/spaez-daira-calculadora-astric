package errors

import (
	"ASTRIC/BackEnd/ASTRIC/helper/env"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

/*
ErrorControlResponse Controla los errores en los controllers de puntos finales
@pila: Corresponde a donde se esta llamando el controlador de errores. Generalmente en el nombre de la funcion contenedora
@w: El ResponseWriter ( w http.ResponseWriter)
@r: El Request (r *http.Request)
*/
func ErrorControlResponse(pila string, w http.ResponseWriter, r *http.Request) {

	err := r.Body.Close()
	if err != nil {
		panic("Error en cerrar Body CODE=ER3:" + err.Error())
	}
	if env.GetEnvGeneral().Panic {
		reco := recover()
		if reco != nil {
			errorSend(reco, pila, w, r)
		}
	}
}

/*
ErrorControl Controla los errores en las funciones
@pila: Corresponde a donde se esta llamando el controlador de errores. Generalmente en el nombre de la funcion contenedora
@w: El ResponseWriter ( w http.ResponseWriter)
@r: El Request (r *http.Request)
*/
func ErrorControl(pila string) {
	if env.GetEnvGeneral().Panic {
		reco := recover()
		if reco != nil {
			errorSend(reco, pila, nil, nil)
		}
	}
}

/*
ErrorDeferFunction !!!NO ESTA PROBADA funcion que recibe mas de una funcion en un Defer para ejecurtarlas y controlar errores
@function: todas las funciones sepradas por coma
*/
func ErrorDeferFunction(function ...func()) {
	var fun func()
	for _, fun = range function {
		fun()
	}
}

func errorSend(reco interface{}, pila string, w http.ResponseWriter, r *http.Request) {

	var ruta, adres string

	if r != nil {
		ruta = " - RUTA: " + r.RequestURI
		adres = " - ADRES " + r.RemoteAddr
	}

	if w != nil {
		w.Header().Set("context-type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		err := json.NewEncoder(w).Encode("Erro interno en : ")
		if err != nil {
			panic("Error al decodificar jason para enviar a Discord " + err.Error())
		}
	}

	msj := "APP: " + " ---- PILA: " + pila + " - ERROR: " + fmt.Sprintf("%v", reco) + ruta + adres
	log.Printf("%v", msj)

}
