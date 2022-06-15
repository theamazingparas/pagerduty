package db

import "go.mongodb.org/mongo-driver/mongo"

var db *mongo.Database

func SetMongoDatabase(database *mongo.Database) {
	db = database
}

func DB() *mongo.Database {
	return db
}
