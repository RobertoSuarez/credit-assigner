package main

import (
	"github.com/RobertoSuarez/creditos/controllers"
	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	app.Mount("/credit-assignment", controllers.NewCreditController())

	app.Listen(":3000")
}
