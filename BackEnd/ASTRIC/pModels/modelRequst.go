package pmodels

import (
	"context"
	"database/sql"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

// GlobalRequest estructura compartida proveniente del middleware de aotorizacion.
type GlobalRequest struct {
	Mysql         *sql.DB
	Mongo         *mongo.Database
	MysqlORM      *gorm.DB
	MysqlORMClose context.CancelFunc
	Usuario       Usuario
	Sucursales    Sucursales
	Tkn           string
}

// Usuario modelo para usu de endpoint
type Usuario struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nombre   string             `bson:"nombre" json:"nombre,omitempty"  validate:"required"`
	Apellido string             `bson:"apellido" json:"apellido,omitempty" validate:"required"`
	IDMysql  int                `bson:"id" json:"id,omitempty" validate:"required"`
}

// Sucursales Estructura de listado de Sucursales
type Sucursales struct {
	Iva float32 `bson:"iva" json:"iva"`
}
