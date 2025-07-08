package models

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const ConnectionString = "mongodb://localhost:27017/products"
const DB = "products"
const CollName = "products"

var MongoClient *mongo.Client

func Init() {
	clientOption := options.Client().ApplyURI(ConnectionString)

	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		panic(err)
	}

	MongoClient = client
}
