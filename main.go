package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ronaregen/auth/handlers"
	"github.com/ronaregen/auth/initializers"
	"github.com/ronaregen/auth/middleware"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDatabase()
}

func main() {
	app := fiber.New()
	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("pong")
	})
	app.Post("/signin", handlers.Signin)
	app.Get("/me", middleware.ReqAuth, handlers.Validate)

	app.Listen(":5000")
}
