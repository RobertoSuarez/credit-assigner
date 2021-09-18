package controllers

import (
	"net/http"

	"github.com/RobertoSuarez/creditos/config/db"
	"github.com/RobertoSuarez/creditos/models"
	"github.com/gofiber/fiber/v2"
)

func NewStatisticsContrller(confdb *db.ConfigDB) *fiber.App {
	statisticsHandler := &StatisticsController{
		DB: &models.CreditStore{ConfigDB: confdb},
	}

	statisticsAPI := fiber.New()

	statisticsAPI.Post("/", statisticsHandler.Post)
	return statisticsAPI
}

type StatisticsController struct {
	DB *models.CreditStore
}

func (sc *StatisticsController) Post(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).JSON(sc.DB.AllCredit())
}
