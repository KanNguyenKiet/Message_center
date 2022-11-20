package db

import "database/sql"

func (s *Store) WithTx(tx *sql.Tx) StoreQuerier {
	return &Store{
		Queries: s.Queries.WithTx(tx),
		db:      s.db,
	}
}
