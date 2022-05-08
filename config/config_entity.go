package config

import (
	"mortgage-project/db"
	enumMode "mortgage-project/enums/emumMode"
)

type configuration struct {
	mode enumMode.Mode
	db   db.DBConfig
}
