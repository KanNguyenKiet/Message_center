package main

import "database/sql"

func newDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
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
