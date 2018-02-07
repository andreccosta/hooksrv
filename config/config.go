package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// Config represents the config file
type Config struct {
	Hooks map[string]Hook `json:"hooks"`
}

// ReadConfig reads configuration file
func ReadConfig(path string) Config {
	raw, err := ioutil.ReadFile(path)

	if err != nil {
		log.Fatal(err)
	}

	var c Config
	err = json.Unmarshal(raw, &c)

	if err != nil {
		log.Fatal(err)
	}

	return c
}