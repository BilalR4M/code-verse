package main

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"codetype-api/handlers"
	"codetype-api/services"
)

func main() {
	app := fiber.New(fiber.Config{
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	})

	// Middleware
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000", // Frontend URL
		AllowHeaders: "Origin, Content-Type, Accept",
		AllowMethods: "GET, POST, HEAD, PUT, DELETE, PATCH",
	}))

	// Initialize services
	sessionService := services.NewSessionService()

	// Initialize handlers
	languageHandler := handlers.NewLanguageHandler()
	snippetHandler := handlers.NewSnippetHandler()
	sessionHandler := handlers.NewSessionHandler(sessionService)

	// Routes
	api := app.Group("/api")

	// Language routes
	api.Get("/languages", languageHandler.GetLanguages)

	// Snippet routes
	api.Get("/snippets", snippetHandler.GetSnippets)

	// Session routes
	api.Post("/sessions/start", sessionHandler.StartSession)
	api.Post("/sessions/finish", sessionHandler.FinishSession)

	// Health check
	api.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "ok",
			"time":   time.Now().Unix(),
		})
	})

	log.Printf("Server starting on port 8080")
	log.Fatal(app.Listen(":8080"))
}