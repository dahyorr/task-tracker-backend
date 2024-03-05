package routes

import (
	"github.com/dahyorr/task-tracker-backend/middlewares"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func RegisterRoutes(app *fiber.App) {
	app.Get("/swagger/*", swagger.HandlerDefault) // default

	// Auth routes
	authRouter := app.Group("/auth")
	authRouter.Post("/register", createUser)
	authRouter.Post("/login", loginUser)
	authRouter.Get("/session", middlewares.AuthMiddleware, getSession)
	authRouter.Get("/me", middlewares.AuthMiddleware, getUser)
	authRouter.Post("/refresh", middlewares.AuthMiddleware, refreshSession)
	authRouter.Post("/logout", middlewares.AuthMiddleware, logout)

	workspacesRouter := app.Group("/workspaces", middlewares.AuthMiddleware)
	// workspacesRouter.Post("/", middlewares.AuthMiddleware, createWorkspace)
	workspacesRouter.Get("/", getWorkspaces)
	workspacesRouter.Post("/", createWorkspace)
	workspacesRouter.Get("/:id", getWorkspaceDetails)
	workspacesRouter.Delete("/:id", deleteWorkspace)
	workspacesRouter.Post("/:id/members", addUserToWorkspace)
	workspacesRouter.Post("/:id/members/:userId", removeUserFromWorkspace)
	workspacesRouter.Get("/:id/members", getWorkspaceMembers)
	workspacesRouter.Get("/:id/tasks", getWorkspaceTasks)

	TaskGroup := app.Group("/tasks", middlewares.AuthMiddleware)
	TaskGroup.Get("/:id", getTask)
	TaskGroup.Post("/", createTask)
	TaskGroup.Delete("/:id", deleteTask)
	TaskGroup.Patch("/:id/status", updateStatus)

}
