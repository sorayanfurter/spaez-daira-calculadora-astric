package server

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"

	// Register plantilla de swagger
	_ "ASTRIC/BackEnd/ASTRIC/config/swagger"
	"ASTRIC/BackEnd/ASTRIC/helper/env"

	httpSwagger "github.com/swaggo/http-swagger"
)

// StartSWA Inicializa el servidor de documentacion de la api
func StartSWA(status chan string, msg chan string) {

	config := env.GetEnvDoc()
	routes := mux.NewRouter().StrictSlash(false)
	routes.PathPrefix("/" + config.URL).Handler(httpSwagger.WrapHandler)

	msg <- "Ejecutando SWAGGER: URL = http://" + env.GetEnvGeneral().ServerIP + ":" + env.GetEnvDoc().Port + "/" + config.URL

	server := http.Server{
		Addr:              ":" + env.GetEnvDoc().Port,
		ReadHeaderTimeout: time.Second * time.Duration(env.GetEnvDoc().Timeout),
		Handler:           routes,
	}

	err := server.ListenAndServe()
	if err != nil {
		status <- "Error en la ejecucion de APIS puerto:" + env.GetEnvDoc().Port
	}

}
