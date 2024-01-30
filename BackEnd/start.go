package backend

import (
	"ASTRIC/BackEnd/ASTRIC/config/environment"
	"ASTRIC/BackEnd/ASTRIC/helper/env"
	"ASTRIC/BackEnd/ASTRIC/server"
	messages "ASTRIC/BackEnd/Messages"
	"fmt"
	"log"
	"os"
)

// StartAPI Funcion principal
func StartAPI(modo string) {

	colorRed := "\033[31m"
	colorGreen := "\033[32m"
	colorYellow := "\033[33m"
	colorReset := "\033[0m"

	verifiErr := server.VerificarSO()
	if verifiErr != nil {
		fmt.Println(string(colorRed), verifiErr.Error())
		os.Exit(0)
	}

	errEnv := environment.LoadEnv()
	server.ClearConsole()
	environment.LoadEnvDefault()

	if errEnv != nil {
		fmt.Println(string(colorRed), errEnv.Error())
		fmt.Println(string(colorGreen), "Se cargaron las variables de entorno por defecto")
	}

	if modo != "BeckEndTKN" {
		env.SetTokenStatus()
	}

	status := make(chan string)
	msg := make(chan string)
	var ms string

	if env.GetConfigAPI().Enable {
		go server.StartAPI(status, msg)
		ms = <-msg
		fmt.Println(string(colorGreen), ms)
	}

	if env.GetConfigWS().Enable {
		go server.StartWs(status, msg)
		ms = <-msg
		fmt.Println(string(colorGreen), ms)
	}

	if env.GetEnvDoc().Enable {
		go server.StartSWA(status, msg)
		ms = <-msg
		fmt.Println(string(colorGreen), ms)
	}

	if !env.GetEnvGeneral().Produccion && !env.GetEnvGeneral().DisableConsola {

		messages := messages.GetMessagesIni()

		messages["DEBELOPEN"] = "La api esta ejecutandose en modo DESARROLLO"
		if env.GetEnvGeneral().DisableToken {
			messages["!IMPORTANTE¡"] = "Se ecuentra desactivada la validacion de TOKEN por lo tanto en Autorisations se debe enviar el nombre del usuarios en ves de TOKEN"
		}
		if env.GetEnvDoc().Enable {
			messages["COMANDO SWAGGER"] = "swagger generate spec -o ./swagger.json --scan-models"
		}

		for title, m := range messages {
			fmt.Print(string(colorYellow), title+" ---")
			fmt.Println(string(colorReset), m)
		}
	}

	err := <-status
	if err != "" {
		fmt.Println(string(colorRed), "Se produgeron errores al ejecutar algunos servidores:")
		log.Fatal(err)
	}

}
