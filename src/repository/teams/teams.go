package teams

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"pagerduty/src/models"
	"pagerduty/src/repository/db"
)

const COLLECTION = "teams"

func CreateTeam(ctx context.Context, teamName string, developers []models.Developer) (string, error) {
	team := models.Team{
		TeamID:     primitive.NewObjectID(),
		Name:       teamName,
		Developers: developers,
	}

	_, err := db.DB().Collection(COLLECTION).InsertOne(ctx, team)
	return team.TeamID.String(), err
}

func GetTeamByID(ctx context.Context, teamID string) (*models.Team, error) {

	id, err := primitive.ObjectIDFromHex(teamID)
	if err != nil {
		return nil, err
	}

	filter := bson.D{{"_id", id}}
	result := db.DB().Collection(COLLECTION).FindOne(ctx, filter)
	if result == nil {
		return nil, errors.New("Unexpected error")
	}
	if result.Err() != nil && result.Err() != mongo.ErrNoDocuments {
		return nil, result.Err()
	}
	var team models.Team
	err = result.Decode(&team)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return &team, nil
}
