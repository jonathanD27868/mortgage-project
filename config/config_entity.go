package config

import (
	"mortgage-project/db"
	enumMode "mortgage-project/enums/emumMode"
)

type config struct {
	mode enumMode.Mode
	db   db.DBConfig
}
