package main

import (
	"os"

	"github.com/RobertoSuarez/creditos/config/db"
	"github.com/RobertoSuarez/creditos/controllers"
	"github.com/gofiber/fiber/v2"
)

func main() {
	configDB := db.NewConfigDB("postgres://qysutatpjecelb:666b9428a543a4cc9d9c8aef579ad4a57c1f88f9bc9f442df4601c499e7019e7@ec2-34-228-154-153.compute-1.amazonaws.com:5432/d3jatadh6qabag")
	app := fiber.New()

	app.Mount("/credit-assignment", controllers.NewCreditController(configDB))
	app.Mount("/statistics", controllers.NewStatisticsContrller(configDB))

	port := os.Getenv("PORT")
	if os.Getenv("PORT") == "" {
		port = "3000"
	}

	app.Listen(":" + port)
}
