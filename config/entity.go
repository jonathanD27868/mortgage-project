package config

import (
	"mortgage-project/db"
	"mortgage-project/enums"
)

// global variable
var Config configuration

type configuration struct {
	AppConfig AppConfig
	db.DBConfig
}

type AppConfig struct {
	Mode   enums.Mode
	Author string
}

type AppConfigDTO struct {
	Mode   string
	Author string
}
