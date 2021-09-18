package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math"
	"net/http"

	"github.com/RobertoSuarez/creditos/config/db"
	"github.com/RobertoSuarez/creditos/jsonview"
	"github.com/RobertoSuarez/creditos/models"
	"github.com/gofiber/fiber/v2"
)

// var (
// 	CreditAPI *fiber.App
// )

// func init() {

// }

func NewCreditController(confdb *db.ConfigDB) *fiber.App {
	credit := &CreditController{
		DB: models.CreditStore{ConfigDB: confdb},
	}

	// create api
	CreditAPI := fiber.New()
	// Registro de las rutas
	CreditAPI.Get("/", credit.Get)
	CreditAPI.Post("/", credit.Post)
	return CreditAPI
}

type CreditController struct {
	DB models.CreditStore
}

func (credit *CreditController) Get(c *fiber.Ctx) error {
	return c.SendString("Get credit")
}

// Post toma los datos {"investment": 3000}
func (credit *CreditController) Post(c *fiber.Ctx) error {
	investment := make(map[string]int)
	err := json.Unmarshal(c.Body(), &investment)
	if err != nil {
		log.Println(err)
		return c.Status(http.StatusBadRequest).SendString("error, estructura de json")
	}

	var creditHandler CreditAssigner = &Credit{}
	creditrow := &models.Credit{}
	creditrow.Investment = int32(investment["investment"])
	creditrow.Credit_type_300, creditrow.Credit_type_500, creditrow.Credit_type_700, err = creditHandler.Assign(creditrow.Investment)
	if err != nil {
		// Almacenar en la db
		creditrow.Successful = false
		fmt.Println(creditrow)
		credit.DB.SaveCredit(*creditrow)
		return c.SendStatus(http.StatusBadRequest)
	}
	creditrow.Successful = true
	fmt.Println(creditrow)
	credit.DB.SaveCredit(*creditrow)
	return c.Status(http.StatusOK).JSON(jsonview.ModelCreditTojson(creditrow))
}

type CreditAssigner interface {
	Assign(investment int32) (int32, int32, int32, error)
}

type Credit struct{}

// Assign Calcula las posibles cantidades de créditos de $300, $500 y $700
func (c *Credit) Assign(investment int32) (int32, int32, int32, error) {
	var (
		x float64 = 0
		y float64 = 0
		z float64 = 0
	)

	for {
		// Algrebraicamente calcula el valor de <y> y <z>
		y = (float64(investment) - x*300 - z*700) / 500
		z = (float64(investment) - x*300 - y*500) / 700

		//fmt.Println(x*300 + y*500 + z*700)
		//fmt.Println(x, y, z)
		// Si se llega a calcular exactamente los mostos, se sale del for.
		if math.Mod(y, 2) == 0 && (y != 0 || x != 0) {
			break
		}

		// Si no es posible el calculo, retorna el error.
		if x < 0 || y < 0 || z < 0 {
			return 0, 0, 0, errors.New("error, no cuadra")
		}

		x++
	}

	//fmt.Println(x, y, z)
	return int32(x), int32(y), int32(z), nil
}
