package customer

type customerService struct {
	repo customerRepo
}

func (cs customerService) getCustomer(id int) *Customer {
	return cs.repo.getCustomer(id)
}

func (cs customerService) getAllIDs() []int {
	return cs.repo.getAllIDs()
}

func (cs customerService) getAllCustomers() []*Customer {
	return cs.repo.getAllCustomers()
}
