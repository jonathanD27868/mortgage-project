package customer

import (
	"mortgage-project/errors"
	"mortgage-project/globals"
	"mortgage-project/helpers"
)

type daoMysql struct{}

func (d daoMysql) getCustomer(id int) *Customer {
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
	helpers.CheckErr(err, errors.ErrDBQueryExec)

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

func (d daoMysql) getAllIDs() []int {
	var (
		db  = globals.Config.GetDB()
		id  int
		ids []int
	)

	query := `select id from customers order by id`
	rows, err := db.Query(query)
	helpers.CheckErr(err, errors.ErrDBQueryExec)
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&id)
		helpers.CheckErr(err, errors.ErrDBQueryExec)

		ids = append(ids, id)
	}

	return ids
}

func (d daoMysql) getAllCustomers() []*Customer {
	var (
		db          = globals.Config.GetDB()
		c           *Customer
		cs          []*Customer
		id          int
		firstName   string
		lastName    string
		creditScore int
		salary      int
		downPayment int
		houseID     int
	)

	query := `select * from customers order by id`
	rows, err := db.Query(query)
	helpers.CheckErr(err, errors.ErrDBQueryExec)
	defer rows.Close()

	for rows.Next() {
		rows.Scan(&id, &firstName, &lastName, &creditScore, &salary, &downPayment, &houseID)
		c = &Customer{
			id,
			firstName,
			lastName,
			creditScore,
			salary,
			downPayment,
			houseID,
		}
		cs = append(cs, c)
	}

	return cs
}
