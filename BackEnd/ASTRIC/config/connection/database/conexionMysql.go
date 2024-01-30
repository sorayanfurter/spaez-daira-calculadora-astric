package database

import (
	"ASTRIC/BackEnd/ASTRIC/helper/env"
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	// Register solo el driver de mysql para que lo utilice database/sql
	_ "github.com/go-sql-driver/mysql"
)

// SetDbMysql Debuelve la conexion a Mysql con la bd setada
func SetDbMysql(database string) *sql.DB {
	con := getCN(database)
	return con
}

// GetMysqlCN Debuelve la conexion a Mysql
func GetMysqlCN() *sql.DB {
	con := getCN(env.GetEnvMysql().Db)
	return con
}

// SetMysqlCNORM Debuelve una conexion ORM a la bd solicitada
func SetMysqlCNORM(database string) (*gorm.DB, context.CancelFunc) {
	return getCNORM(database)
}

// GetMysqlCNORM Debuelve una conexion ORM
func GetMysqlCNORM() (*gorm.DB, context.CancelFunc) {
	return getCNORM(env.GetEnvMysql().Db)
}

func getCNORM(database string) (*gorm.DB, context.CancelFunc) {

	env := env.GetEnvMysql()

	usuario := env.User
	pass := env.Pass
	host := "tcp(" + env.Host + ":" + strconv.Itoa(env.Port) + ")"
	db := database

	//Ejemplo de conexion user:password@/dbname?charset=utf8&parseTime=True&loc=Local
	config := mysql.Config{
		DSN:                       fmt.Sprintf("%s:%s@%s/%s?charset=utf8&parseTime=True&loc=Local", usuario, pass, host, db), // data source name
		DefaultStringSize:         256,                                                                                       // default size for string fields
		DisableDatetimePrecision:  true,                                                                                      // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,                                                                                      // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,                                                                                      // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false,                                                                                     // auto configure based on currently MySQL version
	}

	log := logger.Default.LogMode(logger.Silent)
	if env.Logger {
		log = logger.Default.LogMode(logger.Info)
	}

	configGorm := &gorm.Config{
		Logger: log,
	}

	con, err := gorm.Open(mysql.New(config), configGorm)
	if err != nil {
		panic("Error al conectarse con Mysql :" + err.Error())
	}

	set, err := con.DB()
	if err != nil {
		panic("Error al conectarse con Mysql :" + err.Error())
	}

	set.SetMaxIdleConns(env.SetMaxIdleConns)
	set.SetMaxOpenConns(env.SetMaxOpenConns)
	set.SetConnMaxLifetime(time.Hour * time.Duration(env.SetConnMaxLifetime))

	/*
		Hay que investigar como manejar las conexiones a la bd y los contextos para que no colapse la bd

		https://gorm.io/

	*/

	_, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(env.SetMaxTimeOutCon))
	ctx, _ := context.WithTimeout(context.Background(), time.Second*time.Duration(env.SetMaxTimeOutCon))

	return con.WithContext(ctx), cancel
}

func getCN(database string) *sql.DB {

	env := env.GetEnvMysql()

	usuario := env.User
	pass := env.Pass
	host := "tcp(" + env.Host + ":" + strconv.Itoa(env.Port) + ")"
	db := database

	//Ejemplo de string de conexion = "user:password@protocolo(host:puerto)/databaseName?charset=utf8"
	stringConextion := fmt.Sprintf("%s:%s@%s/%s?charset=utf8", usuario, pass, host, db)

	con, err := sql.Open("mysql", stringConextion)
	if err != nil {
		panic("Error al conectarse con Mysql :" + err.Error())
	}

	con.SetConnMaxLifetime(time.Hour * time.Duration(env.SetConnMaxLifetime))
	con.SetMaxOpenConns(env.SetMaxOpenConns)
	con.SetMaxIdleConns(env.SetMaxIdleConns)
	return con

}
