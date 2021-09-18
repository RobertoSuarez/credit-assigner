package models

import (
	"context"
	"fmt"

	"github.com/RobertoSuarez/creditos/config/db"
	"github.com/RobertoSuarez/creditos/jsonview"
)

type CreditStore struct {
	*db.ConfigDB
}

// SaveCredit Almacena el credito en la db.
func (cs *CreditStore) SaveCredit(credit Credit) error {
	conn, err := cs.OpenDB(context.Background()) // Abrimos conexión con la config principal.
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close(context.Background())
	_, err = conn.Exec(
		context.Background(),
		`call insert_credit($1, $2, $3, $4, $5)`,
		credit.Investment, credit.Credit_type_300, credit.Credit_type_500, credit.Credit_type_700, credit.Successful)
	if err != nil {
		return err
	}
	return nil
}

// AllCredit trae todo los datos de las estadisticas.
func (cs *CreditStore) AllCredit() jsonview.Statistics {
	data := jsonview.Statistics{}
	conn, err := cs.OpenDB(context.Background())
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close(context.Background())
	rows, err := conn.Query(context.Background(), `select * from get_estaditicas();`)
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		err = rows.Scan(
			&data.Total_assignments,
			&data.Total_assignments_successful,
			&data.Total_assignments_failed,
			&data.Average_investment_successful,
			&data.Average_investment_failed)
		if err != nil {
			fmt.Println(err)
		}
	}
	fmt.Println(data)
	return data
}
