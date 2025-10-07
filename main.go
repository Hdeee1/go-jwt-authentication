package main

import (
	"go-jwt-authenticaion/database"
	"go-jwt-authenticaion/routes"

	"github.com/gofiber/fiber"
)

func main() {
	database.Connect()

	app := fiber.New()
	routes.Setup(app)

	app.Listen(":2020")
}