package customer

type customerController struct {
	service customerService
}

func (cc customerController) GetCustomer(id int) Customer {
	return cc.service.getCustomer(id)
}
