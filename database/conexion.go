package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoC = conectarDB()
var clientOptions = options.Client().ApplyURI("mongodb+srv://instagram_admin:e3Zrx4YiUCpYfsX1@farmacia.1k3hs.mongodb.net/myFirstDatabase?retryWrites=true&w=majority")

func conectarDB() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
		return client
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
		return client
	}

	log.Println("Conexión exitosa a la base de datos")
	return client
}

func chequeoBase() int {
	err := mongoC.Ping(context.TODO(), nil)

	if err != nil {
		return 0
	}

	return 1
}
