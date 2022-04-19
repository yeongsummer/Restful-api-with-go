package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yeongsummer/restful-api-with-go/config"
	"github.com/yeongsummer/restful-api-with-go/database"
	"github.com/yeongsummer/restful-api-with-go/router"
)

func main() {
	app := fiber.New()
	dsn := database.MakeDSN(config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_NAME"))
	database.ConnetDB(dsn)

	router.SetupRoutes(app)
	app.Listen(":8080")
}
