package db

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"

	"mortgage-project/customerrors"
	"mortgage-project/globals"
	"mortgage-project/helpers"
	"os"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

func New() DBConfig {
	var dbConfig DBConfig
	dbConfigDTO := getDBConfig()
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
	helpers.CheckErr(errors.New(""), customerrors.ErrDBEngineInvalid, dbConfig.Engine)

	return nil
}

// getDBConfig get db's config from JSON
func getDBConfig() DBConfigDTO {
	filename := fmt.Sprintf("internals/db/db-%s.env.json", globals.Config.GetMode())
	config := DBConfigDTO{}

	file, err := os.Open(filename)
	helpers.CheckErr(err, customerrors.ErrorFileOpening, filename)
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	helpers.CheckErr(err, customerrors.ErrJSONDecoding)

	return config
}

func getPostgresConn(db DBConfigDTO) *sql.DB {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", db.User, db.Password, db.Server, db.Port, db.Database)

	conn, err := sql.Open("postgres", dsn)
	helpers.CheckErr(err, customerrors.ErrDBOpen)

	fmt.Println("Postgres DB's ready")
	fmt.Println()

	return conn
}

func getMysqlConn(db DBConfigDTO) *sql.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?tls=false&autocommit=true", db.User, db.Password, db.Server, db.Port, db.Database)

	conn, err := sql.Open("mysql", dsn)
	helpers.CheckErr(err, customerrors.ErrDBOpen)

	fmt.Println("MySQL DB's ready")
	fmt.Println()

	return conn
}
