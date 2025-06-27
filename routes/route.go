package routes

import (
	"log"
	"net/http"
	"warung/config"

	"github.com/labstack/echo/v4"
)

func ApiRoutes(e *echo.Echo) {
	_, err := config.InitDB()
	if err != nil {
		log.Fatal("Failed Connect to Database", err)
	}

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "This is MC Member API")
	})

	// HEALTH CHECK
	e.GET("/health-check", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"status": "API is running",
		})
	})
}
