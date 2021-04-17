package database

import (
	"context"
	"log"

	"github.com/Jsagudelo1704/Go/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() (*mongo.Client, *mongo.Database, *mongo.Collection) {
	conf := config.GetConfig()
	connection := options.Client().ApplyURI(conf.Mongo.Server)
	client, err := mongo.Connect(context.TODO(), connection)
	if err != nil {
		log.Fatal(err)
	}
	db := client.Database(conf.Mongo.Database)
	collec := db.Collection(conf.Mongo.Collection)
	return client, db, collec
}
