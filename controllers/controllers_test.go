package controllers_test

import (
	"testing"

	"github.com/RobertoSuarez/creditos/controllers"
)

func TestAssign(t *testing.T) {
	var credit controllers.CreditAssigner = &controllers.Credit{}
	var monto int32 = 3000
	c300, c500, c700, err := credit.Assign(monto)
	if err != nil {
		if c300 != 0 || c500 != 0 || c700 != 0 {
			t.Error("Algun credito es diferente a 0")
			t.Fail()
		}
	} else {
		if c300*300+c500*500+c700*700 != monto {
			t.Fail()
		}
	}

	t.Log("Test Assign correcto.")
}
