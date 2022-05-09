package house

type houseDAOInterface interface {
	getHouse(id int) *House
}
