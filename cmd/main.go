package main

import (
	"context"
	"log"
	"os/signal"
	"syscall"

	"github.com/binsabit/fasthttp-v1/internal/api"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer cancel()
	if err := api.Run(ctx); err != nil {
		log.Fatal(err)
	}
}
