package routes

import (
	"dayo.dev/task-tracker/middlewares"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func RegisterRoutes(app *fiber.App) {
	app.Get("/swagger/*", swagger.HandlerDefault) // default

	// Auth routes
	authGroup := app.Group("/auth")
	authGroup.Post("/register", createUser)
	authGroup.Post("/login", loginUser)
	authGroup.Get("/session", middlewares.AuthMiddleware, getSession)
	authGroup.Get("/me", middlewares.AuthMiddleware, getUser)

}
