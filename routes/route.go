package routes

import (
	"log"
	"net/http"
	"warung/config"
	"warung/controllers"

	"github.com/labstack/echo/v4"
)

func ApiRoutes(e *echo.Echo) {
	db, err := config.InitDB()
	if err != nil {
		log.Fatal("Failed Connect to Database", err)
	}

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "This is Warung API")
	})

	// HEALTH CHECK
	e.GET("/health-check", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"status": "API is running",
		})
	})

	// USER
	userController := controllers.NewUserController(db)
	e.POST("/auth/register", userController.Register)
	e.POST("/auth/login", userController.Login)
}
