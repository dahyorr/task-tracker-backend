package middlewares

import (
	"github.com/dahyorr/task-tracker-backend/models"
	"github.com/dahyorr/task-tracker-backend/utils"
	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(ctx *fiber.Ctx) error {
	authToken := ctx.Cookies(utils.Config.SessionCookieName)
	session, err := models.GetSessionByToken(authToken)

	if err != nil || session.IsExpired() {
		return fiber.ErrUnauthorized
	}
	ctx.Locals("session", session)
	ctx.Locals("user_Id", session.UserId)
	return ctx.Next()
}
