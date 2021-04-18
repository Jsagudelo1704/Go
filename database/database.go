package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Crear conexion a la BD en mongo. Se retorna el cliente para poderlo finalizar cuando termine la peticion
//Se reotrna la coleccion para poder operar sobre ella
func ConnectDB(ctx context.Context) (*mongo.Client, *mongo.Collection) {

	//conf := config.GetConfig()
	//connection := options.Client().ApplyURI(conf.Mongo.Server)
	connection := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(ctx, connection)
	if err != nil {
		log.Fatal(err)
	}

	//db := client.Database(conf.Mongo.Database)
	//collec := db.Collection(conf.Mongo.Collection)
	db := client.Database("Mutant")
	collec := db.Collection("dnaverified")

	return client, collec
}
