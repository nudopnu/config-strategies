package main

import (
	"log"

	"github.com/nudopnu/config-loading/golang/internal/config"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%+v", config)
}
