package postgesql_test

import (
	"context"
	"testing"

	"github.com/binsabit/fasthttp-v1/internal/config"
	"github.com/binsabit/fasthttp-v1/internal/data/postgesql"
	"github.com/stretchr/testify/assert"
)

func TestConnectoin(t *testing.T) {
	config := config.MustLoad()
	// dsn := "postgres://yerdaulet:pa55word@localhost:5432/prosclad"
	pool, err := postgesql.NewPGXPool(context.Background(), config.Storage)
	assert.Nil(t, err)

	err = pool.Ping(context.Background())
	assert.Nil(t, err)
}
