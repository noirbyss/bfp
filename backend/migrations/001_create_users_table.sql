CREATE TABLE users (
    id BIGINT FOREIGN KEY,
    emain VARCHAR(128) NOT NULL,
    username varchar(64) NOT NULL,
    password_hash text NOT NULL
)