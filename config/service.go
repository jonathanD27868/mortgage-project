package config

import (
	"encoding/json"
	"log"
	"mortgage-project/db"
	"mortgage-project/enums"
	"os"
)

func New() configuration {
	config := configuration{}

	// get db connection
	config.DB = db.New()

	// get global configs
	config.AppConfig = NewAppConfig("config/config.env.json")

	return config
}

func NewAppConfig(filename string) AppConfig {
	configDTO := AppConfigDTO{}

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Error while opening the app's config file: %s", err.Error())
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&configDTO)
	if err != nil {
		log.Fatalf("Error while decoding the app's JSON config file : %s", err.Error())
	}

	config := AppConfig{
		Mode:   enums.FromString(configDTO.Mode),
		Author: configDTO.Author,
	}

	return config
}
