package main

import (
	"fmt"
	"mortgage-project/config"
	"mortgage-project/customer"
	"mortgage-project/globals"
	"mortgage-project/house"
)

func init() {
	globals.Config = config.New("config/config.env.json", "db/db-prod.env.json")
}

func main() {
	houseController := house.GetHouseController()
	fmt.Println(houseController.GetHouse(4))

	customerController := customer.GetCustomerController()
	fmt.Println(customerController.GetCustomer(3))
}
