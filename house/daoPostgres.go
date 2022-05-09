package house

import (
	"mortgage-project/errors"
	"mortgage-project/globals"
)

type daoPostgres struct{}

func (d daoPostgres) getHouse(id int) House {
	db := globals.Config.GetDB()
	var (
		price           int
		min_downpayment int
		property_tax    int
		maintenance_fee int
	)

	query := `SELECT price, min_downpayment, property_tax, maintenance_fee FROM houses WHERE id=$1`

	err := db.QueryRow(query, id).Scan(&price, &min_downpayment, &property_tax, &maintenance_fee)
	checkErr(err, errors.ErrDBQueryExec)

	h := House{
		id,
		price,
		min_downpayment,
		property_tax,
		maintenance_fee,
	}

	return h
}
