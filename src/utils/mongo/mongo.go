package mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

const (
	MONGO_HOST = "localhost"
	MONGO_PORT = "27017"
)

func NewMongoClient() (*mongo.Client, error) {
	host := MONGO_HOST
	port := MONGO_PORT

	mongoUrl := fmt.Sprintf("mongodb://%v:%v", host, port)

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoUrl))
	return client, err
}
