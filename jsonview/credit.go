package jsonview

import "github.com/RobertoSuarez/creditos/models"

type Credit struct {
	Credit_type_300 int32 `json:"credit_type_300"`
	Credit_type_500 int32 `json:"credit_type_500"`
	Credit_type_700 int32 `json:"credit_type_700"`
}

// Convertir de models.credit a json.credit, para enviar al cliente
func ModelCreditTojson(mCredit *models.Credit) *Credit {
	return &Credit{
		Credit_type_300: mCredit.Credit_type_300,
		Credit_type_500: mCredit.Credit_type_500,
		Credit_type_700: mCredit.Credit_type_700,
	}
}
