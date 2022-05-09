package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

func New(dbConfigFile string) DBConfig {
	var dbConfig DBConfig
	dbConfigDTO := getDBConfig(dbConfigFile)
	dbConfig.DB = getConnection(dbConfigDTO)
	dbConfig.Engine = dbConfigDTO.Engine
	return dbConfig
}

// getConnection get the correct db's connection according to the Engine
// specified in the config file
func getConnection(dbConfig DBConfigDTO) *sql.DB {
	switch dbConfig.Engine {
	case "mysql":
		return getMysqlConn(dbConfig)
	case "postgres":
		return getPostgresConn(dbConfig)
	}
	log.Fatalf("No implementation for the engine: %s", dbConfig.Engine)
	return nil
}

// getDBConfig get db's config from JSON
func getDBConfig(filename string) DBConfigDTO {
	config := DBConfigDTO{}

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Error while opening %s", filename)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		log.Fatalf("Error while decoding %s", err.Error())
	}

	return config
}

func getPostgresConn(db DBConfigDTO) *sql.DB {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", db.User, db.Password, db.Server, db.Port, db.Database)

	conn, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalln(err)
	}
	return conn
}

func getMysqlConn(db DBConfigDTO) *sql.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?tls=false&autocommit=true", db.User, db.Password, db.Server, db.Port, db.Database)

	conn, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalln(err)
	}
	return conn
}
