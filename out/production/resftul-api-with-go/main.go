package main

import (
	"github.com/yeongsummer/restful-api-with-go/database"
)

func main() {
	app := fiber.New()

	database.ConnetDB()

	app.Get("/", func(c *fiber.Ctx) error {
		err := c.SendString("And the API is up!")
		return err
	})

	app.Listen(":8080")
}
