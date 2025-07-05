package main

import (
	"fmt"
	"log"
	"os"
	"warung/routes"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	// Load .env File
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var secretKey = []byte(os.Getenv("JWT_SECRET"))
	fmt.Println("Secret Length:", len(secretKey))

	//Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())
	e.Use(middleware.Recover())

	// Routes
	routes.ApiRoutes(e)

	localhost := os.Getenv("LOCALHOST")
	port := os.Getenv("APP_PORT")

	e.Logger.Fatal(e.Start(localhost + ":" + port))
}
