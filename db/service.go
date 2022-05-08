package db

import (
	"encoding/json"
	"log"
	"os"
)

func New(filename string) DBConfig {
	config := DBConfig{}

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Error while opening the db's config file: %s", err.Error())
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		log.Fatalf("Error while decoding the db's JSON config file : %s", err.Error())
	}

	return config
}
