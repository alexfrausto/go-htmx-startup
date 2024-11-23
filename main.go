package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"log"
	"os"
	"startup/config"
	"startup/controllers"
)

func init() {
	config.Init()
	config.InitializeDB()
}

func main() {
	engine := html.New("./views", ".html")
	engine.Debug(os.Getenv("ENV") == "dev")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/users", controllers.GetUsers)

	cfg := config.GetConfig()
	log.Fatal(app.Listen(fmt.Sprintf(":%d", cfg.Port)))
}
