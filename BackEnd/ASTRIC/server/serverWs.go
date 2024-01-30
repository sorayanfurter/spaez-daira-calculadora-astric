package server

import (
	"ASTRIC/BackEnd/ASTRIC/helper/env"
	"ASTRIC/BackEnd/ASTRIC/helper/socket"
	"net/http"
	"time"

	"github.com/olahol/melody"
)

// StartWs Activa el servidor SOCKET
func StartWs(status chan string, msg chan string) {
	PORT := env.GetConfigWS().Port
	m := melody.New()
	m.Upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	var err error

	serverhttp := http.Server{
		Addr:              ":" + PORT,
		ReadHeaderTimeout: time.Second * time.Duration(env.GetConfigWS().Timeout),
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err = m.HandleRequest(w, r)
	})

	socket.NewSocket(m)

	msg <- "Ejecutando WEBSOCKET: URL = http://" + env.GetEnvGeneral().ServerIP + ":" + PORT + "/"
	err = serverhttp.ListenAndServe()

	if err != nil {
		status <- "Error en la ejecucion de WS puerto:" + err.Error()
	}

}
