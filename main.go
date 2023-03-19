package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ronaregen/auth/handlers"
	"github.com/ronaregen/auth/middleware"
)

func main() {
	app := fiber.New()
	app.Post("/signin", handlers.Signin)
	app.Get("/me", middleware.ReqAuth, handlers.Validate)

	app.Listen(":5000")
}
