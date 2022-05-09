package house

import "fmt"

type House struct {
	Id             int
	Price          int
	MinDownpayment int
	PropertyTax    int
	MaintenanceFee int
}

func (h House) String() string {
	return fmt.Sprintf("%d %d %d %d %d",
		h.Id, h.Price, h.MinDownpayment, h.PropertyTax, h.MaintenanceFee)
}
