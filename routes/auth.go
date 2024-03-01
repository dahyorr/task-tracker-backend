package routes

import (
	"dayo.dev/task-tracker/models"
	"dayo.dev/task-tracker/shared"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func createUser(c *fiber.Ctx) error {
	//TODO:  validate input
	var userForm models.UserFormData
	err := c.BodyParser(&userForm)
	if err != nil {
		return err
	}
	user, err := userForm.CreateUser()

	if err != nil {
		log.Error(err)
		return shared.ErrSomethingWentWrong
	}
	return c.JSON(user)
}

func loginUser(c *fiber.Ctx) error {
	//TODO:  validate input
	var userForm models.UserLoginFormData
	if err := c.BodyParser(&userForm); err != nil {
		return err
	}
	user, err := models.GetUserByEmail(userForm.Email)
	if err != nil {
		log.Error(err)
		return shared.ErrInvalidEmailOrPassword
	}
	err = user.ValidatePassword(userForm.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"details": shared.ErrInvalidEmailOrPassword.Error()})
	}
	session, err := models.NewSession(user.Id)
	if err != nil {
		log.Error(err)
		return shared.ErrSomethingWentWrong
	}
	cookie := session.ToCookie()
	c.Cookie(cookie)

	return c.JSON(user)
}

func getSession(c *fiber.Ctx) error {
	session := c.Locals("session").(*models.Session)
	return c.JSON(session)
}

func getUser(c *fiber.Ctx) error {
	session := c.Locals("session").(*models.Session)
	user, err := models.GetUserById(session.UserId)
	if err != nil {
		log.Error(err)
		return shared.ErrSomethingWentWrong
	}
	return c.JSON(user)
}
