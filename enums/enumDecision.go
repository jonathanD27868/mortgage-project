package enums

import "log"

type Decision struct {
	slug string
}

var (
	UnknownDecision = Decision{""}
	Approved        = Decision{"approved"}
	Rejected        = Decision{"rejected"}
)

func (d Decision) String() string {
	return d.slug
}

func FromStringDecision(s string) Decision {
	switch s {
	case Approved.slug:
		return Approved
	case Rejected.slug:
		return Rejected
	}
	log.Fatalf("Unknown decision for the mortgage")
	return UnknownDecision
}
