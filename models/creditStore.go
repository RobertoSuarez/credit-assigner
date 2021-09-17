package models

import "fmt"

var (
	CreditoDB []Credit
)

func init() {
	CreditoDB = make([]Credit, 0)
}

type CreditStore struct {
}

func (cs *CreditStore) SaveCredit(credit Credit) {
	CreditoDB = append(CreditoDB, credit)
	fmt.Println(CreditoDB)
}
