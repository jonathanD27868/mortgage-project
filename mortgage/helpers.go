package mortgage

import (
	"mortgage-project/customer"
	"mortgage-project/house"
)

func GetMortgageController(c customer.CustomerController, h house.HouseController) mortgageController {
	service := mortgageService{c, h}
	controller := mortgageController{service}
	return controller
}
