package house

type HouseController struct {
	service houseService
}

func (hc HouseController) GetHouse(id int) *House {
	return hc.service.getHouse(id)
}
