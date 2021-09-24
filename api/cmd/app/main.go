package main

import (
	"log"

	"github.com/FourLineCode/url-shortener/api/internal/config"
	"github.com/FourLineCode/url-shortener/api/pkg/server"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	cfg := config.GetConfig()
	s := server.New(cfg)

	if err := s.Listen(cfg.Port); err != nil {
		log.Fatalln("Error starting server", err.Error())
	}
}
