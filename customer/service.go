package customer

type customerService struct {
	repo customerRepo
}

func (sr customerService) getCustomer(id int) *Customer {
	return sr.repo.getCustomer(id)
}
