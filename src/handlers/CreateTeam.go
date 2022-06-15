package handlers

import (
	"context"
	"github.com/labstack/echo/v4"
	"net/http"
	"pagerduty/src/handlers/utils"
	"pagerduty/src/models"
	"pagerduty/src/repository/teams"
)

type CreateTeamRequest struct {
	Team       Team               `json:"team"`
	Developers []models.Developer `json:"developers"`
}

type Team struct {
	Name string `json:"name"`
}

type CreateTeamResponse struct {
	TeamID string `json:"team_id"`
}

func CreateTeamHandler(c echo.Context) error {

	var request CreateTeamRequest

	err := c.Bind(&request)
	if err != nil {
		return utils.HandleError(c, http.StatusBadRequest, err)
	}

	ctx := context.Background()

	//Cream team
	id, err := teams.CreateTeam(ctx, request.Team.Name, request.Developers)
	if err != nil {
		return utils.HandleError(c, http.StatusInternalServerError, err)
	}

	response := CreateTeamResponse{
		TeamID: id,
	}

	return c.JSON(http.StatusOK, response)
}
