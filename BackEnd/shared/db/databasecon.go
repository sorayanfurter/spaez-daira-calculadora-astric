package db

import (
	"ASTRIC/BackEnd/ASTRIC/config/connection/database"
	"ASTRIC/BackEnd/ASTRIC/global"
)

// MongoDb conexion a la bd de mongoDb con la base de datos setada desde el request
var MongoDb = global.GetMongoDB

// MySQL conexion a la bd de mysql con la base de datos setada desde el request
var MySQL = global.GetMysql

// MysqlORM conexion ORM a la bd de mysql con la base de datos setada desde el request
var MysqlORM = global.GetMysqlORM

/*
MongoNewCN Funcion que retorna la coleccion lista para interactuar con la BD
@database: El nombre de la bd
@colection: El nopmbre de la colection
*/
var MongoNewCN = database.MongoColection

/*
MySQLNewCN Debuelve la conexion a Mysql
@database: Nombre de la bd con la que se creara la conexion
*/
var MySQLNewCN = database.SetDbMysql

/*
MysqlORMNewCN Debuelve la conexion a Mysql ORM
@database: Nombre de la bd con la que se creara la conexion
*/
var MysqlORMNewCN = database.SetMysqlCNORM
