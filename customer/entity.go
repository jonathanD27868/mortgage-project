package customer

import "fmt"

type Customer struct {
	Id          int
	FirstName   string
	LastName    string
	CreditScore int
	Salary      int
	DownPayment int
	HouseID     int
}

func (c Customer) String() string {
	return fmt.Sprintf("%d %s %s %d %d %d %d",
		c.Id, c.FirstName, c.LastName, c.CreditScore, c.Salary, c.DownPayment, c.HouseID)
}
