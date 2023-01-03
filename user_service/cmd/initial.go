package main

import (
	"github.com/jmoiron/sqlx"
)

func newDB(dsn string) (*sqlx.DB, error) {
	db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	// ping to make sure database work
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
