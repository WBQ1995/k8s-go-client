package config

import (
	"github.com/BurntSushi/toml"
	"log"
	"strings"
)

type Configuration struct {
	Port string
	Dev  bool
}

var config *Configuration

func initConfig(configFile string) {
	if strings.Trim(configFile, " ") == "" {
		configFile = "./configs/config.toml"
	}
	if metaData, err := toml.DecodeFile(configFile, &config); err != nil {
		log.Fatal("error:", err)
	} else {
		if !requiredFieldsAreGiven(metaData) {
			log.Fatal("Required fields not given")
		}
	}
}

func GetConfig() Configuration {
	if config == nil {
		initConfig("")
	}
	return *config
}

func requiredFieldsAreGiven(metaData toml.MetaData) bool {
	requiredFields := [][]string{
		{"port"},
	}

	for _, v := range requiredFields {
		if !metaData.IsDefined(v...) {
			log.Fatal("required fields ", v)
		}
	}

	return true
}
