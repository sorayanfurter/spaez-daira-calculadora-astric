package environment

import (
	"ASTRIC/BackEnd/ASTRIC/helper/env"
	"ASTRIC/BackEnd/config"
	"net/http"
	"time"
)

// LoadEnvDefault Carga las variebles de entorno
func LoadEnvDefault() {

	//GENERALES
	env.EnvGeneral.Produccion = env.Ifbool("PRODUCCION")                        //Pone en produccion la API
	env.EnvGeneral.Panic = env.Ifbool("PANIC")                                  //Activa el control de errores
	env.EnvGeneral.DisableConsola = env.Ifbool("DISABLE_CONSOLA")               //Desactiva la consola
	env.EnvGeneral.Notificaciones = env.Ifbool("DISABLENOTIFICACIONES")         //Desabilita las notificaciones creadas en /Messages
	env.EnvGeneral.ContextTime = env.Ifint("CONTEXT_TIME", 60)                  //60 segundos para el contexto de conexiones con las bases de datos
	env.EnvGeneral.ClaveSecreta = env.Ifstring("CLAVE_SECRET", "calvedeprueba") //Clave de encriptacion, token, hash, etc.
	env.EnvGeneral.TiempoToken = env.Ifint("EXPIRE_TOKEN", 259200)              //1 año --- Tempo de caducidad de token
	env.EnvGeneral.DisableToken = env.Ifbool("DISABLE_TOKEN")                   //Desactiva el control de token (se envia usuario)
	env.EnvGeneral.DBASTRIC = env.Ifstring("DB_ASTRIC", "ASTRIC")               //Base de datos del Frameworks en Mysql
	env.EnvGeneral.TimeoutClient = env.Ifint("TIMEOUT_CLIENT", 3)               //Tiempo de inactividad del cliente en minutos
	env.EnvGeneral.TimeoutPingDB = env.Ifint("TIMEOUT_PING_DB", 8)              //Tiemo de prueba la realizar el PING a la BD
	env.EnvGeneral.ServerIP = env.Ifstring("SERVER_IP", "localhost")            //IP del servidor

	//API
	env.EnvAPI.Enable = config.LoadConfigENV().APIEnanle
	env.EnvAPI.Port = env.Ifstring("API_PORT", "3000") //Puerot por defecto del API
	env.EnvAPI.Timeout = env.Ifint("API_TIMEOUT", 3)   //Tiempo de inactividad del endpoint en segundos

	//Websocket
	env.EnvWS.Enable = config.LoadConfigENV().WSEnable
	env.EnvWS.Port = env.Ifstring("WS_PORT", "8000") //Puerot por defecto del API
	env.EnvWS.Timeout = env.Ifint("WS_TIMEOUT", 3)   //Tiempo de inactividad del endpoint en segundos

	//DOCUMENTACION
	env.EnvDoc.Enable = config.LoadConfigENV().DOCEnable
	env.EnvDoc.URL = env.Ifstring("DOC_URL", "")       //URL de documentacion
	env.EnvDoc.Port = env.Ifstring("DOC_PORT", "3100") //Puerot de SWAGGER
	env.EnvDoc.Timeout = env.Ifint("WS_TIMEOUT", 3)    //Tiempo de inactividad del endpoint en segundos
	env.EnvDoc.Token = env.Ifstring("DOC_TOKEN_MSG", "Token desactivado envie nombre de usuarios")

	//MONGODB
	env.EnvMongo.Enable = env.Ifbool("MONGO_ENABLE")
	env.EnvMongo.StrCon = env.Ifstring("MONGO_URL", config.LoadConfigENV().MONGOURL)

	//MYSQL
	env.EnvMysql.Enable = env.Ifbool("MYSQL_ENABLE")
	env.EnvMysql.Db = env.Ifstring("MYSQL_DB", "daira")
	env.EnvMysql.Host = env.Ifstring("MYSQL_HOST", config.LoadConfigENV().MYSQLHost)
	env.EnvMysql.Pass = env.Ifstring("MYSQL_PASS", config.LoadConfigENV().MYSQLPass)
	env.EnvMysql.User = env.Ifstring("MYSQL_USER", config.LoadConfigENV().MYSQLUser)
	env.EnvMysql.Port = env.Ifint("MYSQL_PORT", config.LoadConfigENV().MYSQLPort)
	env.EnvMysql.SetConnMaxLifetime = env.Ifint("MYSQL_SET_MAX_LIFE_TIMA", 3)
	env.EnvMysql.SetMaxIdleConns = env.Ifint("MYSQL_SET_MAX_IDLE_CONNS", 10)
	env.EnvMysql.SetMaxOpenConns = env.Ifint("MYSQL_SET_MAX_OPEN_CONN", 10)
	env.EnvMysql.SetMaxTimeOutCon = env.Ifint("MYSQL_MAX_TIMEOUT_CON", 60)
	env.EnvMysql.Logger = !env.Ifbool("MYSQL_LOGER")

	config.LoadConfigUser()
	loadConfigHTTP()

}

func loadConfigHTTP() {
	http.DefaultClient.Timeout = time.Minute * time.Duration(env.GetEnvGeneral().TimeoutClient)
}
