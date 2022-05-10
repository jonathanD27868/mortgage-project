package main

import (
	"fmt"
	"mortgage-project/config"
	"mortgage-project/customer"
	"mortgage-project/globals"
	"mortgage-project/house"
	"mortgage-project/mortgage"
	"runtime"
)

func init() {
	maxCPU := runtime.NumCPU()
	cpuUsed := 4
	runtime.GOMAXPROCS(cpuUsed)
	fmt.Printf("Number of CPUs (Total=%d - Used=%d) \n", maxCPU, cpuUsed)

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
