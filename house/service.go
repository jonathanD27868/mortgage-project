package house

type houseService struct {
	repo houseRepo
}

func (hs houseService) getHouse(id int) House {
	return hs.repo.getHouse(id)
}
