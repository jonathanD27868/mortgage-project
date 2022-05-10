package customer

type customerRepo struct {
	dao customerDAOInterface
}

func (cr customerRepo) getCustomer(id int) *Customer {
	return cr.dao.getCustomer(id)
}

func (cr customerRepo) getAllIDs() []int {
	return cr.dao.getAllIDs()
}

func (cr customerRepo) getAllCustomers() []*Customer {
	return cr.dao.getAllCustomers()
}
