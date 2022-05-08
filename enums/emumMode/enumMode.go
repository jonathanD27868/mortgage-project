package enumMode

import "errors"

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

func FromString(s string) (Mode, error) {
	switch s {
	case Dev.slug:
		return Dev, nil
	case Prod.slug:
		return Prod, nil
	}
	return Unknown, errors.New("unknown mode: " + s)
}
