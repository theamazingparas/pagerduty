package handlers

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"pagerduty/src/handlers/utils"
	"pagerduty/src/repository/teams"
	"pagerduty/src/utils/notifications"
)

func AlertTeamHandler(c echo.Context) error {

	teamId := c.Param("id")

	if teamId == "" {
		return utils.HandleError(c, http.StatusBadRequest, fmt.Errorf("Team id not provided"))
	}

	ctx := context.Background()

	//Find team by id
	team, err := teams.GetTeamByID(ctx, teamId)
	if err != nil {
		return utils.HandleError(c, http.StatusInternalServerError, err)
	}

	//Send alerts
	status := notifications.SendNotifications(team.Developers)

	return c.JSON(http.StatusOK, status)
}
