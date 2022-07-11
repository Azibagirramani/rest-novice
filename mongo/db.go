package mongo

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DbConnect() *mongo.Client {

	print := fmt.Println
	connectionString := "mongodb://localhost:27017" // edit to connect to your own remote mongo instance
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(connectionString))

	if err != nil {
		print(err, "--- from mongo connection")
	}

	print(" Mongo conncted ")

	return client
}
