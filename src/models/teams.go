package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Team struct {
	TeamID     primitive.ObjectID `json:"_id" bson:"_id"`
	Name       string             `json:"name" bson:"name"`
	Developers []Developer        `json:"developers" bson:"developers"`
}

type Developer struct {
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
}
