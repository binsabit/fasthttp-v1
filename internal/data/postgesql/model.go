package postgesql

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Model struct {
	someModel *pgxpool.Pool
}

func NewModels(pool *pgxpool.Pool) *Model {
	return &Model{someModel: pool}
}

func NewPGXPool(ctx context.Context, connString string) (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(context.Background(), connString)
	if err != nil {
		return nil, fmt.Errorf("pgx connection error: %w", err)
	}
	return pool, nil
}
