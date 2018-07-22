package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// Config represents the configuration file
type Config struct {
	Hooks map[string]Hook `json:"hooks"`
}

// Load sets up the configuration
func Load(path string) Config {
	return read(path)
}

func read(path string) Config {
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
