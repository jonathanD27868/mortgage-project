package config

import (
	"database/sql"
	"fmt"
	"mortgage-project/enums"
	"mortgage-project/internals/db"
)

type Configuration struct {
	AppConfig
	db.DBConfig
}

func (c Configuration) PrintGlobal() {
	fmt.Printf("mode: %s | author: %s | db engine: %s", c.Mode, c.Author, c.Engine)
}

func (c Configuration) GetMode() string {
	fmt.Println("******************$")
	return c.Mode.String()
}

func (c Configuration) GetAuthor() string {
	return c.Author
}

func (c Configuration) GetDB() *sql.DB {
	return c.DB
}

func (c Configuration) GetDBEngine() string {
	return c.Engine
}

type AppConfig struct {
	Mode   enums.Mode
	Author string
}

type AppConfigDTO struct {
	Mode   string
	Author string
}
