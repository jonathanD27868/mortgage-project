package mortgage

import (
	"mortgage-project/customer"
	"mortgage-project/enums"
	"mortgage-project/house"
	"sync"
)

type mortgageService struct {
	c customer.CustomerController
	h house.HouseController
}

func (ms mortgageService) getApprovalDecision(id int) *mortgageDecision {
	c := ms.c.GetCustomer(id)
	return ms.decisionMaker(c)
}

func (ms mortgageService) decisionMaker(c *customer.Customer) *mortgageDecision {
	house := ms.h.GetHouse(c.HouseID)

	yearlyMortage := float32(house.Price-c.DownPayment) * 0.0612
	yearlyExpense := yearlyMortage + float32(house.PropertyTax) + float32(house.MaintenanceFee)
	income := float32(c.Salary) * 0.32

	isCreditScoreOk := c.CreditScore >= 650
	isDownPaymentOk := c.DownPayment >= house.MinDownpayment
	isIncomeOk := income >= yearlyExpense

	isAllConditions := isCreditScoreOk && isDownPaymentOk && isIncomeOk
	var decision enums.Decision = enums.Rejected

	if isAllConditions {
		decision = enums.Approved
	}

	md := mortgageDecision{
		c.Id,
		c.FirstName,
		c.LastName,
		c.Salary,
		c.DownPayment,
		decision,
	}

	return &md
}

func (ms mortgageService) getApprovalDecisionAllCustomers() []*mortgageDecision {
	customers := ms.c.GetAllCustomers()
	var decisions []*mortgageDecision
	var wg sync.WaitGroup
	var mu sync.Mutex

	for _, c := range customers {
		wg.Add(1)
		go func(c *customer.Customer) {
			d := ms.decisionMaker(c)
			// fmt.Println(d)

			// lock to avoid data racing
			mu.Lock()
			decisions = append(decisions, d)
			mu.Unlock()

			wg.Done()
		}(c)
	}
	wg.Wait()

	return decisions
}
