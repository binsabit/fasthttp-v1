package postgesql_test

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/binsabit/fasthttp-v1/config"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/assert"
)

func TestConnectoin(t *testing.T) {
	config := config.MustLoad()
	storageConf := config.Storage
	log.Println(storageConf)
	// dsn := "postgres://yerdaulet:pa55word@localhost:5432/prosclad"
	dns := fmt.Sprintf(`user=%s password=%s host=%s port=%s 
	dbname=%s sllmode=%s pool_max_conns=%d pool_min_conns=%d pool_max_conn_idle_time=%s`,
		storageConf.User,
		storageConf.Password,
		storageConf.Host,
		storageConf.Port,
		storageConf.DBName,
		storageConf.SSLMode,
		storageConf.PoolMaxConns,
		storageConf.PoolMinConns,
		storageConf.PoolMaxConnIdleTime,
	)
	poolConfig, err := pgxpool.ParseConfig(dns)
	assert.Nil(t, err)

	db, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	assert.NotNil(t, db)
	assert.Nil(t, err)
	err = db.Ping(context.Background())
	assert.Nil(t, err)
}
