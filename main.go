package main

import (
	"fmt"
	"mortgage-project/config"
)

func init() {
	config.Config = config.New("db/db-dev.env.json", "config/config.env.json")
}

func main() {
	fmt.Println(config.Config)
	fmt.Println("ok")
	test1Select()

}

func test1Select() {
	fmt.Println("\n==> test1Select")
	db := config.Config.DB
	mode := config.Config.Mode.String()
	fmt.Println(mode)

	// Execute the query
	results, err := db.Query("SELECT id, price FROM houses")
	checkErr(err)

	for results.Next() {
		var house struct {
			id              int
			price           int
			min_downpayment int
			property_tax    int
			maintenance_fee int
		}

		err = results.Scan(&house.id, &house.price)
		checkErr(err)

		fmt.Printf("%d\t%d \n", house.id, house.price)
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}
