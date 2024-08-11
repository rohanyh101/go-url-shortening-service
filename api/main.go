package main

import (
	// "context"
	"fmt"
	"log"
	"os"

	"github.com/roh4nyh/shorten_url/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func setupRoutes(app *fiber.App) {
	app.Get("/:url", routes.ResolveURL)
	app.Post("/api/v1", routes.ShortenURL)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}

	if err != nil {
		log.Fatalf("Error creating Redis client: %v", err)
	}

	app := fiber.New()

	app.Use(logger.New())

	setupRoutes(app)

	log.Fatal(app.Listen(os.Getenv("PORT")))
}

// 18.6
