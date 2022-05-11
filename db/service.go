package db

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
<<<<<<< Updated upstream:db/service.go
	"log"
=======
	"mortgage-project/customerrors"
	"mortgage-project/globals"
	"mortgage-project/helpers"
>>>>>>> Stashed changes:internals/db/helpers.go
	"os"

	"mortgage-project/errors"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

func New(dbConfigFile string) *sql.DB {
	dbConfig := getDBConfig(dbConfigFile)
	return getConnection(dbConfig)
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
<<<<<<< Updated upstream:db/service.go
	log.Fatalf(errors.ErrDBEngineInvalid.Error(), dbConfig.Engine)
=======

	helpers.CheckErr(errors.New(""), customerrors.ErrDBEngineInvalid, dbConfig.Engine)
>>>>>>> Stashed changes:internals/db/helpers.go
	return nil
}

// getDBConfig get db's config from JSON
func getDBConfig(filename string) DBConfigDTO {
	config := DBConfigDTO{}

	file, err := os.Open(filename)
<<<<<<< Updated upstream:db/service.go
	if err != nil {
		log.Fatalf(errors.ErrOpenFile.Error(), filename, err.Error())
	}
=======
	helpers.CheckErr(err, customerrors.ErrorFileOpening, filename)
>>>>>>> Stashed changes:internals/db/helpers.go
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
<<<<<<< Updated upstream:db/service.go
	if err != nil {
		log.Fatalf(errors.ErrJSONDecoding.Error(), err.Error())
	}
=======
	helpers.CheckErr(err, customerrors.ErrJSONDecoding)
>>>>>>> Stashed changes:internals/db/helpers.go

	return config
}

func getPostgresConn(db DBConfigDTO) *sql.DB {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", db.User, db.Password, db.Server, db.Port, db.Database)

	conn, err := sql.Open("postgres", dsn)
<<<<<<< Updated upstream:db/service.go
	if err != nil {
		log.Fatalln(err)
	}
=======
	helpers.CheckErr(err, customerrors.ErrDBOpen)

	fmt.Println("Postgres DB's ready")
	fmt.Println()
>>>>>>> Stashed changes:internals/db/helpers.go
	return conn
}

func getMysqlConn(db DBConfigDTO) *sql.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?tls=false&autocommit=true", db.User, db.Password, db.Server, db.Port, db.Database)

	conn, err := sql.Open("mysql", dsn)
<<<<<<< Updated upstream:db/service.go
	if err != nil {
		log.Fatalln(err)
	}
=======
	helpers.CheckErr(err, customerrors.ErrDBOpen)

	fmt.Println("MySQL DB's ready")
	fmt.Println()
>>>>>>> Stashed changes:internals/db/helpers.go
	return conn
}
