package house

import (
	"log"
	"mortgage-project/globals"
)

func GetHouseController() HouseController {
	dao := daoFactory(globals.Config.GetDBEngine())
	repo := houseRepo{dao}
	service := houseService{repo}
	controller := HouseController{service}
	return controller
}

func daoFactory(e string) houseDAOInterface {
	var i houseDAOInterface
	switch e {
	case "mysql":
		i = daoMysql{}
	case "postgres":
		i = daoPostgres{}
	default:
		log.Fatalf("the engine %s is not implemented", e)
	}
	return i
}
