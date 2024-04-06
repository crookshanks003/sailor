package main

import (
	"log"

	"github.com/pritesh-mantri/sailor/config"
	"github.com/pritesh-mantri/sailor/internal/data"
	"github.com/pritesh-mantri/sailor/cmd/server"
)

func main() {
	cfg := config.ReadConfig()
	app := server.Application{
		Config: cfg,
		Models: data.New(cfg),
	}

	r := app.Routes()

	err := r.Run(cfg.Port)
	if err != nil {
		log.Fatalf("failed to run server: %e", err)
	}
}
