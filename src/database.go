package src

import (
	"log"
	"pagerduty/src/repository/db"
	"pagerduty/src/utils/mongo"
)

const DATABASE = "pagerduty"

func InitDatabase() {
	mongoClient, err := mongo.NewMongoClient()
	if err != nil {
		log.Fatal("Error initialising mongo client: %v", err)
	}

	db.SetMongoDatabase(mongoClient.Database(DATABASE))
}
