package customer

type CustomerController struct {
	service customerService
}

func (cc CustomerController) GetCustomer(id int) *Customer {
	return cc.service.getCustomer(id)
}

func (cc CustomerController) GetAllIDs() []int {
	return cc.service.getAllIDs()
}

func (cc CustomerController) GetAllCustomers() []*Customer {
	return cc.service.getAllCustomers()
}
