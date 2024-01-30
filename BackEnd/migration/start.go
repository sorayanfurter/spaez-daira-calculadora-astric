package migration

import (
	"ASTRIC/BackEnd/ASTRIC/config/connection/database"
	"ASTRIC/BackEnd/ASTRIC/config/environment"
	"ASTRIC/BackEnd/ASTRIC/helper/env"
	"ASTRIC/BackEnd/ASTRIC/server"
	"ASTRIC/BackEnd/config"
	"fmt"
	"os"
)

// StartMigration inicia la configuracion de la BD
func StartMigration() {

	colorRed := "\033[31m"
	colorGreen := "\033[32m"

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

	con, cancel := database.SetMysqlCNORM(env.GetEnvGeneral().DBASTRIC)
	defer cancel()

	if err := createUserStructBD(con); err != nil {
		fmt.Println(string("\033[31m"), "Migrado de usuario"+err.Error())
		return
	}

	con, cancel = database.SetMysqlCNORM(env.GetEnvMysql().Db)
	defer cancel()

	if err := config.Migration(con); err != nil {
		fmt.Println(string("\033[31m"), "Migrado de generales"+err.Error())
		return
	}

	fmt.Println(string("\033[32m"), "Migracion finalizada con exito")

}
