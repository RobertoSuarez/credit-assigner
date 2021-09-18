package models

import (
	"context"
	"fmt"

	"github.com/RobertoSuarez/creditos/config/db"
)

var (
	CreditoDB []Credit
)

func init() {
	CreditoDB = make([]Credit, 0)
}

type CreditStore struct {
	*db.ConfigDB
}

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

// AllCredit trae todos los creditos
func (cs *CreditStore) AllCredit() []Credit {
	conn, err := cs.OpenDB(context.Background())
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close(context.Background())
	rows, err := conn.Query(context.Background(), `select "ID", "Investment", "Credit_type_300", "Credit_type_500", "Credit_type_700", "Successful" from credit;`)
	if err != nil {
		fmt.Println(err)
	}

	for rows.Next() {
		cre := Credit{}
		err = rows.Scan(&cre.ID, &cre.Investment, &cre.Credit_type_300, &cre.Credit_type_500, &cre.Credit_type_700, &cre.Successful)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(cre)
	}

	return CreditoDB
}
