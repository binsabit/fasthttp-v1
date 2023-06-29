package postgesql_test

import (
	"context"
	"testing"

	"github.com/binsabit/fasthttp-v1/config"
	"github.com/binsabit/fasthttp-v1/internal/data/postgesql"
	"github.com/stretchr/testify/assert"
)

func TestConnectoin(t *testing.T) {
	config := config.Configure()
	// dsn := "postgres://yerdaulet:pa55word@localhost:5432/prosclad"
	pool, err := postgesql.NewPGXPool(context.Background(), config.DB_DSN)
	assert.Nil(t, err)

	err = pool.Ping(context.Background())
	assert.Nil(t, err)
}
