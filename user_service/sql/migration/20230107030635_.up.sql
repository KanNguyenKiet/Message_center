ALTER TABLE users
ADD COLUMN session_key varchar(64),
ADD COLUMN session_expired timestamp;


ALTER TABLE credential
MODIFY COLUMN password_hashed varchar(64);