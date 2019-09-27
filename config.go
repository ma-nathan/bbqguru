package main

import (
	"github.com/BurntSushi/toml"
	"log"
	"os"
)

// $ cat settings.ini
//
// BBQHost="bbq.fumanchu.com"
// BBQUser="admin"
// BBQPassword="bbq"
// DatabaseURL="http://metrics:8086"
// DatabaseUser="admin"
// DatabasePassword="J500icu"
// DatabaseDatabase="bbq"
//

type Config struct {
	BBQHost	string
	BBQUser	string
	BBQPassword string
	DatabaseURL      string
	DatabaseUser     string
	DatabasePassword string
	DatabaseDatabase string
}

var configfile = "settings.ini"

func ReadConfig() Config {
	_, err := os.Stat(configfile)
	if err != nil {
		log.Fatal("Config file is missing: ", configfile)
	}

	var config Config
	if _, err := toml.DecodeFile(configfile, &config); err != nil {
		log.Fatal(err)
	}
	//log.Print(config.Index)
	return config
}
