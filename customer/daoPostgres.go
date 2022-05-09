package customer

import (
	"mortgage-project/errors"
	"mortgage-project/globals"
)

type daoPostgres struct{}

func (d daoPostgres) getCustomer(id int) *Customer {
	db := globals.Config.GetDB()
	var (
		firstName   string
		lastName    string
		creditScore int
		salary      int
		downPayment int
		houseID     int
	)

	query := `
	SELECT first_name, last_name, credit_score, salary, downpayment, house_id FROM customers
	WHERE id=?
	`

	err := db.QueryRow(query, id).
		Scan(&firstName, &lastName, &creditScore, &salary, &downPayment, &houseID)
	checkErr(err, errors.ErrDBQueryExec)

	c := Customer{
		id,
		firstName,
		lastName,
		creditScore,
		salary,
		downPayment,
		houseID,
	}

	return &c
}
