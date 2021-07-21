package database

import (
	"context"
	"log"

	"github.com/Jsagudelo1704/Go/structs"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Crear conexion a la BD en mongo. Se retorna el cliente para poderlo finalizar cuando termine la peticion
//Se reotrna la coleccion para poder operar sobre ella
func ConnectDB(ctx context.Context, conf structs.MongoConfiguration) (*mongo.Client, *mongo.Collection) {

	//conf := config.GetConfig()
	//connection := options.Client().ApplyURI(conf.Mongo.Server)
	connection := options.Client().ApplyURI(conf.Server)

	//connection := options.Client().ApplyURI("mongodb+srv://mongoadmin:Juanito_1704@cluster0.cxpm1.mongodb.net/test")
	client, err := mongo.Connect(ctx, connection)
	if err != nil {
		log.Fatal(err)
	}
	db := client.Database(conf.Database)
	collec := db.Collection(conf.Collection)

	return client, collec
}
