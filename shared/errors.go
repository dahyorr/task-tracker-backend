package shared

import (
	"github.com/gofiber/fiber/v2"
)

var (
	ErrNoPasswordProvided     = fiber.NewError(400, "No password provided")
	ErrInvalidEmailOrPassword = fiber.NewError(401, "Invalid email/password")
	// ErrCreatingSession        = fiber.NewError("error creating session")
	// ErrCreatingUser           = fiber.NewError("error creating user")
	ErrInvalidAuthSession     = fiber.NewError(401, "Invalid Authentication Credential")
	ErrUserEmailAlreadyExists = fiber.NewError(400, "User with the email already exists")
	ErrUserNotFound           = fiber.NewError(400, "User not found")
	ErrSomethingWentWrong     = fiber.NewError(500, "something went wrong")
)
