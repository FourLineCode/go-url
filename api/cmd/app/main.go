package main

import (
	"log"

	"github.com/FourLineCode/url-shortener/api/internel/config"
	"github.com/FourLineCode/url-shortener/api/pkg/server"
)

func main() {
	s := server.New()
	cfg := config.GetConfig()

	if err := s.Listen(cfg.Port); err != nil {
		log.Fatalln("Error starting server", err.Error())
	}
}
