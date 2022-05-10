package customer

import (
	"log"
	"mortgage-project/enums"
	"mortgage-project/errors"
	"mortgage-project/globals"
)

func GetCustomerController() CustomerController {
	dao := daoFactory(globals.Config.GetDBEngine())
	repo := customerRepo{dao}
	service := customerService{repo}
	controller := CustomerController{service}
	return controller
}

func daoFactory(e string) customerDAOInterface {
	var i customerDAOInterface
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

func checkErr(err error, msg errors.ModelError) {
	if err == nil {
		return
	}
	if globals.Config.GetMode() != enums.Dev.String() {
		log.Fatalln(msg.Error())
	} else {
		log.Fatalln(msg.Error(), "=> ", err)
	}
}
