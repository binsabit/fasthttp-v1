package main

import (
	"context"

	"github.com/binsabit/fasthttp-v1/config"
	"github.com/binsabit/fasthttp-v1/internal/api"
	"github.com/binsabit/fasthttp-v1/internal/data/postgesql"
	"github.com/binsabit/fasthttp-v1/pkg"
)

func main() {

	config := config.Configure()
	logger := pkg.NewLogger(config.LogFile)

	pool, err := postgesql.NewPGXPool(context.Background(), config.DB_DSN)
	if err != nil {
		logger.ErrorLog.Fatalf("could not establish db connetion pool: %v", err)
	}
	defer pool.Close()

	pqStorage := postgesql.NewModels(pool)

	app := api.NewAplication(config.ServerPort, pqStorage, logger)

	err = app.Run()
	if err != nil {
		logger.ErrorLog.Fatalf("could not run application: %v", err)
	}
}
