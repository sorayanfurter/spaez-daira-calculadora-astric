package environment

import (
	"ASTRIC/BackEnd/ASTRIC/helper/env"
	"context"
	"time"
)

// DatabaseContexto Retorna un contexto general para la conexion a la bd toma los valores de .env
func DatabaseContexto() (context.Context, context.CancelFunc) {

	timer := env.GetEnvGeneral().ContextTime
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timer)*time.Second)
	return ctx, cancel

}
