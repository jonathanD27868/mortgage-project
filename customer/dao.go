package customer

type customerDAOInterface interface {
	getCustomer(id int) Customer
}
