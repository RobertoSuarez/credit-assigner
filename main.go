package main

import (
	"os"

	"github.com/RobertoSuarez/creditos/controllers"
	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	app.Mount("/credit-assignment", controllers.NewCreditController())
	app.Mount("/statistics", controllers.NewStatisticsContrller())

	port := os.Getenv("PORT")
	if os.Getenv("PORT") == "" {
		port = "3000"
	}

	app.Listen(":" + port)
}
