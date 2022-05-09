package globals

import "database/sql"

// global variable
var Config globalInterface

type globalInterface interface {
	PrintGlobal()
	GetMode() string
	GetAuthor() string
	GetDB() *sql.DB
	GetDBEngine() string
}
