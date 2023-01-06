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

-- name: UpdateSessionKey :execresult
UPDATE users
SET session_key = ?
WHERE id = ?;