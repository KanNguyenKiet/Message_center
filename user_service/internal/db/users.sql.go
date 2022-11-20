// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: users.sql

package db

import (
	"context"
	"database/sql"
)

const createNewUser = `-- name: CreateNewUser :execresult
INSERT INTO users (
    first_name, last_name, email, phone, user_name
) values (
    ?, ?, ?, ?, ?
)
`

type CreateNewUserParams struct {
	FirstName sql.NullString `json:"first_name"`
	LastName  sql.NullString `json:"last_name"`
	Email     string         `json:"email"`
	Phone     sql.NullString `json:"phone"`
	UserName  sql.NullString `json:"user_name"`
}

func (q *Queries) CreateNewUser(ctx context.Context, arg CreateNewUserParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createNewUser,
		arg.FirstName,
		arg.LastName,
		arg.Email,
		arg.Phone,
		arg.UserName,
	)
}
