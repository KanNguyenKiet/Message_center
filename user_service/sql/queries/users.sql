-- name: CreateNewUser :execresult
INSERT INTO users (
    first_name, last_name, email, phone
) values (
    ?, ?, ?, ?
);