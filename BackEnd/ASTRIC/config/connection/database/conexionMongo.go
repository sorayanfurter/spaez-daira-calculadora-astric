package database

import (
	"ASTRIC/BackEnd/ASTRIC/helper/env"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*
MongoColection Funcion que retorna la coleccion lista para interactuar con la BD
@database: El nombre de la bd
@colection: El nopmbre de la colection
*/
func MongoColection(database string, colecction string) *mongo.Collection {

	if !checkDB(database) {
		panic("MongoColection - La BASE de datos " + database + " no se encuentra")
	}
	db := GetMongoCN().Database(database)
	if !checkCollection(db, colecction) {
		panic("MongoColection - La COLECCION de datos " + colecction + " no se encuentra")
	}

	col := db.Collection(colecction)
	return col
}

/*
MongoDatabase Funcion que retorna la bd para interactuar con las colecciones
@database: la base de datos de la cual queremos obtener las colecciones
*/
func MongoDatabase(database string) *mongo.Database {

	if !checkDB(database) {
		panic("MongoDatabase - La BASE de datos " + database + " no se encuentra")
	}
	db := GetMongoCN().Database(database)
	return db
}

/*
MongoDatabaseUnprotected Funcion que retorna la bd para interactuar con las colecciones
"SIN CHEQUEO DE EXISTENCIA DE LA MISMA"
@database: la base de datos de la cual queremos obtener las colecciones
*/
func MongoDatabaseUnprotected(database string) *mongo.Database {
	db := GetMongoCN().Database(database)
	return db
}

func checkDB(db string) bool {
	con := GetMongoCN()
	databases, _ := con.ListDatabaseNames(context.TODO(), bson.M{"name": db})
	return len(databases) > 0
}

func checkCollection(db *mongo.Database, Collection string) bool {
	collection, _ := db.ListCollectionNames(context.TODO(), bson.M{"name": Collection})
	return len(collection) > 0
}

// GetMongoCN Debuelve la conexion a mongoDB
func GetMongoCN() *mongo.Client {

	var clientOptions = options.Client().ApplyURI(env.GetEnvMongo().StrCon)
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatalln("Error al obtener el cliente de MongoDB")
		return client
	}
	return client
}
