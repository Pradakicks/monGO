package Interfaces

import "go.mongodb.org/mongo-driver/mongo"

type MongoDatabase interface {
	GetVersion() string
	ConnectDB()
	ConnectDataDB()
	GetCollection() *mongo.Collection
	GetDataCollection() *mongo.Collection
}
