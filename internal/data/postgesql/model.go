package postgesql

import (
	"context"
	"fmt"
	"log"

	"github.com/binsabit/fasthttp-v1/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPGXPool(ctx context.Context, storageConf config.Storage) (*pgxpool.Pool, error) {
	log.Println(storageConf)

	dns := fmt.Sprintf(`user=%s password=%s host=%s port=%s 
						dbname=%s sslmode=%s pool_max_conns=%d pool_min_conns=%d pool_max_conn_idle_time=%s`,
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
	log.Println(dns)
	poolConfig, err := pgxpool.ParseConfig(dns)
	if err != nil {
		return nil, err
	}

	db, err := pgxpool.NewWithConfig(context.Background(), poolConfig)

	if err != nil {
		return nil, err
	}
	err = db.Ping(context.Background())
	if err != nil {
		return nil, err
	}
	return db, nil
}
