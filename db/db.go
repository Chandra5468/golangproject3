package db

import (
	"database/sql"

	"github.com/go-sql-driver/mysql"
)

func NewMySqlStorage(cfg mysql.Config) (*sql.DB, error) {
	db, err := sql.Open("mysql", cfg.FormatDSN()) // we are establishing a connection here

	if err != nil {
		// log.Fatal("this is opening a new connection", err)
		return nil, err
	}

	return db, nil
}
