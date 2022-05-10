package main

import (
	"fmt"
	"mortgage-project/config"
	"mortgage-project/customer"
	"mortgage-project/globals"
	"mortgage-project/house"
	"mortgage-project/mortgage"
)

func init() {
	globals.Config = config.New("config/config.env.json", "db/db-dev.env.json")
}

func main() {
	hc := house.GetHouseController()
	// fmt.Println(hc.GetHouse(4))

	cc := customer.GetCustomerController()
	// fmt.Println(cc.GetCustomer(3))

	mc := mortgage.GetMortgageController(cc, hc)
	// fmt.Println(mc.GetApprovalDecision(1))
	fmt.Println(mc.GetApprovalDecisionAllCustomers())

}
