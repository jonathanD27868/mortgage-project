package house

type houseController struct {
	service houseService
}

func (hc houseController) GetHouse(id int) House {
	return hc.service.getHouse(id)
}
