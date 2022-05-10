package mortgage

type mortgageController struct {
	service mortgageService
}

func (mc mortgageController) GetApprovalDecision(id int) *mortgageDecision {
	return mc.service.getApprovalDecision(id)
}

func (mc mortgageController) GetApprovalDecisionAllCustomers() []*mortgageDecision {
	return mc.service.getApprovalDecisionAllCustomers()
}
