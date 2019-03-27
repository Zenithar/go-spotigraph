-- +migrate Up
CREATE TABLE IF NOT EXISTS users (
    id          VARCHAR(64) NOT NULL PRIMARY KEY,
    principal   VARCHAR(50) NOT NULL,
    meta        JSON        NOT NULL
);

-- +migrate Down
DROP TABLE users;
