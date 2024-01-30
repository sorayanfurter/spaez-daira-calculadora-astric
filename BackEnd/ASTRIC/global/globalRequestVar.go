package global

import (
	ASTRICmodels "ASTRIC/BackEnd/ASTRIC/pModels"
	"context"
	"database/sql"

	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

// Request Estructura global donde se cargan las BD recividad por los request
var request ASTRICmodels.GlobalRequest

// responceWS Estructura global de responce de request
var responceWS ASTRICmodels.ResponseWS

// SetBDResponceWS setea la bd en el responce del websocket
func SetBDResponceWS(bd string) {
	responceWS.BD = bd
}

// SetTagResponceWS seta el tag del responce websocket
func SetTagResponceWS(tag string) {
	responceWS.Tag = tag
}

// SetUsuarioResponceWS Setea el usuario del responce websocket
func SetUsuarioResponceWS(usuraio string) {
	responceWS.Usuario = usuraio
}

// SetTablaResponceWS setea la tabla del responce websoket
func SetTablaResponceWS(tabla string) {
	responceWS.Tabla = tabla
}

// SetAppResponceWS setea la ruta del websoket
func SetAppResponceWS(app string) {
	responceWS.App = app
}

// SetEndPointResponceWS setea la ruta del websoket
func SetEndPointResponceWS(EndPoint string) {
	responceWS.EndPoint = EndPoint
}

// SetMysql Setea la conexion con la bd solitida
func SetMysql(mysql *sql.DB) {
	request.Mysql = mysql
}

// SetTKN Setea el token renovado
func SetTKN(tkn string) {
	request.Tkn = tkn
}

// SetMongo Setea la bd de mongo con los datos obtenidos del request
func SetMongo(mongo *mongo.Database) {
	request.Mongo = mongo
}

// SetMysqlORM Setea la bd en la conexion de mysql con ORM
func SetMysqlORM(mysqlOrm *gorm.DB, close context.CancelFunc) {
	request.MysqlORM = mysqlOrm
	request.MysqlORMClose = close
}

// GetResponceWS Retrona el reaponceWS
func GetResponceWS() ASTRICmodels.ResponseWS {
	return responceWS
}

// GetTKN retorna el token actualizado
func GetTKN() string {
	return request.Tkn
}

// GetRequestVar devuelve el usuario logeado
func GetRequestVar() ASTRICmodels.GlobalRequest {
	return request
}

// GetMongoDB debuelve la conexion a la bd de mongo ya seteada
func GetMongoDB() *mongo.Database {
	return request.Mongo
}

// GetMysql Debuleve la conexion a la bd de mysql ya seteada
func GetMysql() *sql.DB {
	return request.Mysql
}

// GetMysqlORM Debuelve la conexion a  mysql ORM ya setead
func GetMysqlORM() (*gorm.DB, context.CancelFunc) {
	return request.MysqlORM, request.MysqlORMClose
}
