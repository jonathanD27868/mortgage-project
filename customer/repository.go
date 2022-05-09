package customer

type customerRepo struct {
	dao customerDAOInterface
}

func (cr customerRepo) getCustomer(id int) *Customer {
	return cr.dao.getCustomer(id)
}
