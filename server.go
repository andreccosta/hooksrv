package main

import (
	"log"
	"os"

	"gitlab.com/andreccosta/hooksrv/config"
)

func main() {
	addr, ok := os.LookupEnv("PORT")
	if !ok {
		addr = "8080"
	}

	c := config.ReadConfig("config.json")

	for url, _ := range c.Hooks {
		log.Print("Loaded hook with url\t" + url)
	}
}
