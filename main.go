package main

import (
	"github.com/JacklO0p/weather_forecast/routes"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	routes.RegisteredRoutes(e)

	e.Logger.Fatal(e.Start(":1323"))
}
