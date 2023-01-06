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

const createNewUserCredential = `-- name: CreateNewUserCredential :execresult
INSERT INTO credential(
    user_id, password_hashed
) values (
    ?, ?
)
`

type CreateNewUserCredentialParams struct {
	UserID         int64          `json:"user_id"`
	PasswordHashed sql.NullString `json:"password_hashed"`
}

func (q *Queries) CreateNewUserCredential(ctx context.Context, arg CreateNewUserCredentialParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createNewUserCredential, arg.UserID, arg.PasswordHashed)
}

const getCrendentailByUserId = `-- name: GetCrendentailByUserId :one
SELECT password_hashed from credential
WHERE user_id = ?
`

func (q *Queries) GetCrendentailByUserId(ctx context.Context, userID int64) (sql.NullString, error) {
	row := q.db.QueryRowContext(ctx, getCrendentailByUserId, userID)
	var password_hashed sql.NullString
	err := row.Scan(&password_hashed)
	return password_hashed, err
}

const getUserByUsername = `-- name: GetUserByUsername :one
SELECT id, last_name, first_name, phone, email, user_name, created_at, updated_at, last_login, session_key FROM users
WHERE user_name = ?
`

func (q *Queries) GetUserByUsername(ctx context.Context, userName sql.NullString) (Users, error) {
	row := q.db.QueryRowContext(ctx, getUserByUsername, userName)
	var i Users
	err := row.Scan(
		&i.ID,
		&i.LastName,
		&i.FirstName,
		&i.Phone,
		&i.Email,
		&i.UserName,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.LastLogin,
		&i.SessionKey,
	)
	return i, err
}

const updateSessionKey = `-- name: UpdateSessionKey :execresult
UPDATE users
SET session_key = ?
WHERE id = ?
`

type UpdateSessionKeyParams struct {
	SessionKey sql.NullString `json:"session_key"`
	ID         int64          `json:"id"`
}

func (q *Queries) UpdateSessionKey(ctx context.Context, arg UpdateSessionKeyParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, updateSessionKey, arg.SessionKey, arg.ID)
}
