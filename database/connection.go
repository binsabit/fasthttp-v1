package database

import (
	"database/sql"
	"log"

	"github.com/binsabit/fasthttp-v1/config"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func ConnectToDB(storageConf config.Storage) *sql.DB {
	db, err := sql.Open(storageConf.DBDriver, storageConf.DSN)
	if err != nil {
		log.Fatalf("couldnot eesatblish database connection: %v", err)
	}

	db.SetMaxIdleConns(storageConf.MaxIdleConns)

	db.SetMaxOpenConns(storageConf.MaxOpenConns)

	db.SetConnMaxIdleTime(storageConf.MaxIdleTime)

	db.SetConnMaxLifetime(storageConf.MaxConnLife)

	if err := db.Ping(); err != nil {
		log.Fatalf("unable to reach the databse %v", err)
	}

	return db
}
