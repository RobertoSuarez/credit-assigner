package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {

	var credit CreditAssigner = &Credit{}

	fmt.Println(credit.Assign(6700))
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

// func (c *Credit) Assign(investment int32) (int32, int32, int32, error) {

// 	montos := make([]Monto, 0)
// 	//montos = append(montos, Monto{700, 0}, Monto{500, 0}, Monto{300, 0})
// 	montos = append(montos, Monto{300, 0}, Monto{500, 0}, Monto{700, 0})

// 	index := 0
// 	ultima := false
// 	for {

// 		if investment >= montos[index].Valor {
// 			montos[index].Cantidad++
// 			investment -= montos[index].Valor
// 		}

// 		if investment < montos[0].Valor && len(montos)-1 == index {
// 			ultima = true
// 		}

// 		// Control de index.
// 		if index == len(montos)-1 {
// 			index = 0
// 			if ultima {
// 				fmt.Println(investment)
// 				break
// 			}
// 		} else {
// 			index++
// 		}
// 	}

// 	if investment > 0 {
// 		return montos[0].Cantidad, montos[1].Cantidad, montos[2].Cantidad, fmt.Errorf("No cuadra")
// 	}

// 	return montos[0].Cantidad, montos[1].Cantidad, montos[2].Cantidad, nil
// }

// type Monto struct {
// 	Valor    int32
// 	Cantidad int32
// }
