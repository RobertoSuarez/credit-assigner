package main

import (
	"fmt"
	"log"
	"os"

	"github.com/RobertoSuarez/creditos/config/db"
	"github.com/RobertoSuarez/creditos/controllers"
	"github.com/gofiber/fiber/v2"
)

func main() {
	urldb := os.Getenv("DATABASE_URL")
	if urldb == "" {
		log.Println("No existe url a la db")
		return
	}
	fmt.Println(urldb)

	configDB := db.NewConfigDB(urldb)

	app := fiber.New(fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		StrictRouting: false,
		ServerHeader:  "Server fiber Credit Assignment",
		AppName:       "Credit Assignment v0.0.0",
	})

	app.Mount("/credit-assignment", controllers.NewCreditController(configDB))
	app.Mount("/statistics", controllers.NewStatisticsContrller(configDB))

	port := os.Getenv("PORT")
	if os.Getenv("PORT") == "" {
		port = "3000"
	}

	app.Listen(":" + port)
}
