package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yeongsummer/Restful-api-with-go/config"
	"github.com/yeongsummer/Restful-api-with-go/database"
)

func main() {
	app := fiber.New()
	dsn := database.MakeDSN(config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_NAME"))
	database.ConnetDB(dsn)

	app.Get("/", func(c *fiber.Ctx) error {
		err := c.SendString("And the API is up!")
		return err
	})

	app.Listen(":8080")
}
