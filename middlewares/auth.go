package middlewares

import (
	"dayo.dev/task-tracker/models"
	"dayo.dev/task-tracker/utils"
	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(ctx *fiber.Ctx) error {
	authToken := ctx.Cookies(utils.Config.SessionCookieName)
	session, err := models.GetSessionByToken(authToken)

	if err != nil || session.IsExpired() {
		return fiber.ErrUnauthorized
	}
	ctx.Locals("session", session)
	return ctx.Next()
}
