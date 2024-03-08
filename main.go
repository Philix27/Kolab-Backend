package main

import (
	"fmt"
	_ "harmony/docs"
	appConfig "harmony/libs/config"
	"harmony/libs/helper"
	"harmony/services/app"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

// @title Harmony
// @version 1.0
// @description Harmony Backend Server
// @termsOfService http://swagger.io/terms/

// @contact.name Felix Eligbue
// @contact.email philixbob@gmail.com
// @contact.url    http://www.swagger.io/support

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:3111
// @BasePath /api/v1/

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/

func main() {
	err := godotenv.Load(".env")

	helper.ErrorPanic(err, "Could not load env")

	server := fiber.New(fiber.Config{
		Prefork:           false,
		CaseSensitive:     true,
		StrictRouting:     true,
		ServerHeader:      "Harmony Server",
		AppName:           "Harmony",
		EnablePrintRoutes: true,
		IdleTimeout:       5,
		ReduceMemoryUsage: true,
	})

	server.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:1400",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	config := &appConfig.DbConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASSWORD"),
		User:     os.Getenv("DB_USER"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
		DbName:   os.Getenv("DB_NAME"),
	}

	db, err := appConfig.DbNewConnection(config)
	helper.ErrorPanic(err, "Cannot connect to db")

	appState := app.AppState{
		DB: db,
	}

	appState.SetupRoutes(server)
	
	serverErr := server.Listen(":3434")
	// helper.ErrorPanic(serverErr, "Error loading server")
	if serverErr != nil {
		fmt.Println("Error Error")
		panic("Panic error")
	}
	fmt.Println("After listen")
	// log.Fatal(server.Listen(os.Getenv("PORT")))
}
