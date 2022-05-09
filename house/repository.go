package house

type houseRepo struct {
	dao houseDAOInterface
}

func (hr houseRepo) getHouse(id int) *House {
	return hr.dao.getHouse(id)
}
