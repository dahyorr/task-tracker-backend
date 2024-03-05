package shared

import (
	"github.com/gofiber/fiber/v2"
)

var (
	ErrFailedToParseId     = fiber.NewError(400, "Invalid id id")

	ErrNoPasswordProvided     = fiber.NewError(400, "No password provided")
	ErrInvalidEmailOrPassword = fiber.NewError(401, "Invalid email/password")
	// ErrCreatingSession        = fiber.NewError("error creating session")
	// ErrCreatingUser           = fiber.NewError("error creating user")
	ErrInvalidAuthSession     = fiber.NewError(401, "Invalid Authentication Credential")
	ErrUserEmailAlreadyExists = fiber.NewError(400, "User with the email already exists")
	ErrUserNotFound           = fiber.NewError(404, "User not found")
	ErrSomethingWentWrong     = fiber.NewError(500, "something went wrong")

	ErrInvalidWorkspaceId     = fiber.NewError(400, "Invalid workspace id")
	ErrWorkspaceNotFound      = fiber.NewError(404, "Workspace not found")
	ErrUserNotInWorkspace     = fiber.NewError(400, "User not in workspace")
	ErrUserAlreadyInWorkspace = fiber.NewError(400, "User already in workspace")
	ErrUserNotOwner           = fiber.NewError(400, "User not owner of workspace")

	ErrTaskNotFound = fiber.NewError(404, "Task not found")
)
