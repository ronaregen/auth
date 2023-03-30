package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
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
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))
	app.Post("/login", handlers.Signin)
	app.Get("/me", middleware.ReqAuth, handlers.Validate)

	app.Listen(":5000")
}
