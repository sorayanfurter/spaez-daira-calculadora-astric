package database

import (
	"ASTRIC/BackEnd/ASTRIC/config/connection/database"
	"ASTRIC/BackEnd/ASTRIC/helper/env"
	"ASTRIC/BackEnd/ASTRIC/helper/errors"
	"context"
	"net/http"
	"time"
)

// ChequeBD Middleware para chequear las conexiones de bd hates de realizar alguna transaccion
func ChequeBD(next http.Handler) http.Handler {

	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {

		e := false
		defer errors.ErrorControlResponse("Conexiones en BD /shared/middleware/debug/ChequeoBD", nil, r)
		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(env.GetEnvGeneral().TimeoutPingDB)*time.Second)
		defer cancel()

		if env.GetEnvMysql().Enable {
			err := database.GetMysqlCN().PingContext(ctx)
			if err != nil {
				e = true
				http.Error(rw, " --- Se perdio la conexion con la bd", 500)
				panic("Se perdio la conexion con MySQL " + err.Error())
			}
		}

		if env.GetEnvMongo().Enable {
			err := database.GetMongoCN().Ping(ctx, nil)
			if err != nil {
				e = true
				http.Error(rw, " --- Se perdio la conexion con la bd", 500)
				panic("Se perdio la conexion con MongoDB " + err.Error())
			}
		}

		if !e {
			next.ServeHTTP(rw, r)
		}

	})
}
