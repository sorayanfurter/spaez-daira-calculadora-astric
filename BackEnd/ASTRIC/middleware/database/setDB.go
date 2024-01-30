package database

import (
	"ASTRIC/BackEnd/ASTRIC/config/connection/database"
	"ASTRIC/BackEnd/ASTRIC/global"
)

// SetDb Setea las bd de los usuaraios
func SetDb(mysql string) {

	//col := database.MongoDatabaseUnprotected(mongo)
	db := database.SetDbMysql(mysql)
	orm, cancel := database.SetMysqlCNORM(mysql)

	//global.SetMongo(col)
	global.SetBDResponceWS(mysql)
	global.SetMysql(db)
	global.SetMysqlORM(orm, cancel)

}
