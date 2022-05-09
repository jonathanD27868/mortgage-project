package db

import "database/sql"

type DBConfigDTO struct {
	Engine   string
	Server   string
	Port     string
	User     string
	Password string
	Database string
}

type DBConfig struct {
	DB     *sql.DB
	Engine string
}
