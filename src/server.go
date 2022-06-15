package src

import (
	"fmt"
	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"pagerduty/src/handlers"
)

const (
	PORT    = "1234"
	BaseURL = "/api/v1"
)

func StartServer() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST(fmt.Sprintf("%v/teams", BaseURL), handlers.CreateTeamHandler)
	e.POST(fmt.Sprintf("%v/teams/:id/alert", BaseURL), handlers.AlertTeamHandler)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", PORT)))
}
