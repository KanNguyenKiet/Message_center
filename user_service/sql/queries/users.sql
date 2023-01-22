-- name: CreateNewUser :execresult
INSERT INTO users (
    first_name, last_name, email, phone, user_name
) values (
    ?, ?, ?, ?, ?
);

-- name: CreateNewUserCredential :execresult
INSERT INTO credential(
    user_id, password_hashed
) values (
    ?, ?
);

-- name: GetUserByUsername :one
SELECT * FROM users
WHERE user_name = ?;

-- name: GetCrendentailByUserId :one
SELECT password_hashed from credential
WHERE user_id = ?;

-- name: UpdateSessionInfo :execresult
UPDATE users
SET session_key = ?, session_expired = ?
WHERE id = ?;

-- name: GetUserInfoBySessionKey :one
SELECT last_name, first_name, email, phone, session_expired
FROM users
WHERE session_key = ?;

-- name: GetUserSession :one
SELECT id, session_key, session_expired
FROM users
WHERE user_name = ?;

-- name: ClearUserSession :execresult
UPDATE users
SET session_key = null, session_expired = null
WHERE id = ?;