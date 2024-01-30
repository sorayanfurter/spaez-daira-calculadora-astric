package env

import (
	ASTRICmodels "ASTRIC/BackEnd/ASTRIC/pModels"
	"os"
	"strconv"
)

// ConfigGenral configuracion general
var ConfigGenral ASTRICmodels.Config

// EnvGeneral enviroment general
var EnvGeneral ASTRICmodels.EnvGeneral

// EnvAPI enviromene API
var EnvAPI ASTRICmodels.EnvAPI

// EnvWS enviroment Websocket
var EnvWS ASTRICmodels.EnvWs

// EnvMongo enviroment mongo
var EnvMongo ASTRICmodels.EnvMongo

// EnvMysql enviroment mysql
var EnvMysql ASTRICmodels.EnvMysql

// EnvMails enviroment mails
var EnvMails ASTRICmodels.EnvMails

// EnvDoc enviroment documenttacion
var EnvDoc ASTRICmodels.EnvDoc

// SetTokenStatus forzar desactivado de TKN
func SetTokenStatus() {
	EnvGeneral.DisableToken = true
}

// GetConfigAPI debuelve las variable de API
func GetConfigAPI() ASTRICmodels.EnvAPI {
	return EnvAPI
}

// GetConfigWS debuelve las variable de Websocket
func GetConfigWS() ASTRICmodels.EnvWs {
	return EnvWS
}

// GetConfigGeneral debuelve las variable de documentascion
func GetConfigGeneral() ASTRICmodels.Config {
	return ConfigGenral
}

// GetEnvDoc debuelve las variable de documentascion
func GetEnvDoc() ASTRICmodels.EnvDoc {
	return EnvDoc
}

// GetEnvGeneral debulve las variables de entrono
func GetEnvGeneral() ASTRICmodels.EnvGeneral {
	return EnvGeneral
}

// GetEnvMongo debulve las variables de entrono
func GetEnvMongo() ASTRICmodels.EnvMongo {
	return EnvMongo
}

// GetEnvMysql debulve las variables de entrono
func GetEnvMysql() ASTRICmodels.EnvMysql {
	return EnvMysql
}

// GetEnvMails debulve las variables de entrono
func GetEnvMails() ASTRICmodels.EnvMails {
	return EnvMails
}

// Ifbool funcion de bool para enviroment
func Ifbool(target string) bool {
	retult, _ := strconv.ParseBool(os.Getenv(target))
	return retult
}

// Ifint funcion de int para enviroment
func Ifint(target string, source int) int {
	num, err := strconv.Atoi(os.Getenv(target))
	if err != nil {
		return source
	} else if num < 1 {
		return source
	}
	return num
}

// Ifstring funcion de string para enviroment
func Ifstring(target string, source string) string {
	if os.Getenv(target) != "" {
		return os.Getenv(target)
	}
	return source
}

// Default modelo de configuracion por defecto
type Default struct {
	APIEnanle bool
	DOCEnable bool
	WSEnable  bool
	MYSQLHost string
	MYSQLPort int
	MYSQLPass string
	MYSQLUser string
	MONGOURL  string
}
