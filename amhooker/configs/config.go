package configs

import (
	"amhooker/amhooker/models"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

func ReadConfig(configPath string) *models.AlertConfig {
	var cfg = models.AlertConfig{}
	content, err := os.ReadFile(configPath)
	if err != nil {
		log.Fatalf("Problem reading configuration file: %v", err)
	}
	err = yaml.Unmarshal(content, &cfg)
	if err != nil {
		log.Fatalf("Error parsing configuration file: %v", err)
	}

	return &cfg
}
