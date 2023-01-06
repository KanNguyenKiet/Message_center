ALTER TABLE credential
MODIFY password_hashed varchar(1024);

ALTER TABLE users
DROP COLUMN session_key;