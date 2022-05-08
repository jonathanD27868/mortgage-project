package enums

import (
	"log"
)

type Mode struct {
	slug string
}

var (
	Unknown = Mode{""}
	Dev     = Mode{"dev"}
	Prod    = Mode{"prod"}
)

func (m Mode) String() string {
	return m.slug
}

func FromString(s string) Mode {
	switch s {
	case Dev.slug:
		return Dev
	case Prod.slug:
		return Prod
	}
	log.Fatalf("Unknwon mode in the app's config file: %s", s)
	return Unknown
}
