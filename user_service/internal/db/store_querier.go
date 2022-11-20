package db

import "database/sql"

type StoreQuerier interface {
	Querier
	WithTx(tx *sql.Tx) StoreQuerier
	Transaction(txFunc func(tx *sql.Tx) error) error
}
