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