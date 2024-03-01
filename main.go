package main

import (
	"fmt"

	"dayo.dev/task-tracker/database"
	"dayo.dev/task-tracker/routes"
	"dayo.dev/task-tracker/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/encryptcookie"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	config := utils.InitConfig()
	err := database.Init(config)
	if err != nil {
		panic(err)
	}
	app := fiber.New()

	// Middlewaare init
	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(encryptcookie.New(encryptcookie.Config{
		Key: config.CookieEncryptionKey,
	}))

	fmt.Println("Connected to db...")
	routes.RegisterRoutes(app)
	err = app.Listen(fmt.Sprintf(":%v", config.PORT))
	if err != nil {
		panic(err)
	}
}
