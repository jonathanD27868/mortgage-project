package mortgage

import (
	"fmt"
	"mortgage-project/enums"
)

type mortgageDecision struct {
	id          int
	firstName   string
	lastName    string
	salary      int
	downPayment int
	decision    enums.Decision
}

func (md mortgageDecision) String() string {
	return fmt.Sprintf("%d %s %s %d %d %v",
		md.id, md.firstName, md.lastName, md.salary, md.downPayment, md.decision)
}
