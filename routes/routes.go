package routes

import (
	"github.com/JacklO0p/weather_forecast/controllers"
	"github.com/labstack/echo/v4"
)

func RegisteredRoutes(e *echo.Echo) {
	e.GET("/location", controllers.GetWeather)
}
