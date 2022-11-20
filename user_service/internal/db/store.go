package db

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
)

type Store struct {
	*Queries
	db *sqlx.DB
}

func NewStore(db *sqlx.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

func (s Store) Transaction(txFunc func(tx *sql.Tx) error) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback()
			panic(p) // re-throw panic after Rollback
		} else if err != nil {
			_ = tx.Rollback() // err is non-nil; don't change it
		} else {
			err = tx.Commit() // err is nil; if Commit returns error update err
		}
	}()
	err = txFunc(tx)
	return err
}
