package customer

type customerDAOInterface interface {
	getCustomer(id int) *Customer
	getAllIDs() []int
	getAllCustomers() []*Customer
}
