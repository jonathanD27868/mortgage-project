package config

import (
	"encoding/json"
	"log"
	"mortgage-project/enums"
	"os"
)

func New() Configuration {
	appConfigFile := "internals/config/config.env.json"
	config := Configuration{}

	// get global configs
	config.AppConfig = NewAppConfig(appConfigFile)

	// get db connection
	// config.DBConfig = db.New(dbConfigFile)

	return config
}

func NewAppConfig(filename string) AppConfig {
	configDTO := AppConfigDTO{}

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln("Error while opening the app's config file")
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&configDTO)
	if err != nil {
		log.Fatalln("Error while decoding the app's JSON config file")
	}

	config := AppConfig{
		Mode:   enums.FromString(configDTO.Mode),
		Author: configDTO.Author,
	}

	return config
}
