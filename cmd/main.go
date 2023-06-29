package main

import (
	"context"
	"log"

	"github.com/binsabit/fasthttp-v1/config"
	"github.com/binsabit/fasthttp-v1/internal/api"
	"github.com/binsabit/fasthttp-v1/internal/data/postgesql"
)

func main() {

	config := config.Configure()

	pool, err := postgesql.NewPGXPool(context.Background(), config.DB_DSN)
	if err != nil {
		log.Fatalf("could not establish db connetion pool: %v", err)
	}
	defer pool.Close()

	pqStorage := postgesql.NewModels(pool)

	app := api.NewAplication(config.ServerPort, pqStorage)

	err = app.ConfigureAndRun()
	if err != nil {
		log.Fatalf("could not run application: %v", err)
	}
}
