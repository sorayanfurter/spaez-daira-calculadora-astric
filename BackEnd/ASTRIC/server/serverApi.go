package server

import (
	"ASTRIC/BackEnd/ASTRIC/helper/env"
	a "ASTRIC/BackEnd/ASTRIC/users/routes"
	"ASTRIC/BackEnd/config"
	r "ASTRIC/BackEnd/routers"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// StartAPI Activa el servidor HTTP APIS Rest
func StartAPI(status chan string, msg chan string) {
	rout := mux.NewRouter().StrictSlash(false)
	rout.HandleFunc("/", config.RutaRaiz).Methods("GET")
	r.RutasPrincipales(rout)
	a.RutasModul(rout)
	routes := cors.AllowAll().Handler(rout)

	PORT := env.GetConfigAPI().Port

	msg <- "Ejecutando API ASTRIC: URL = http://" + env.GetEnvGeneral().ServerIP + ":" + PORT + "/"

	server := http.Server{
		Addr:              ":" + PORT,
		ReadHeaderTimeout: time.Second * time.Duration(env.GetConfigAPI().Timeout),
		Handler:           routes,
	}

	err := server.ListenAndServe()
	if err != nil {
		status <- "Error en la ejecucion de APIS puerto:" + PORT + " Err " + err.Error()
	}

}
