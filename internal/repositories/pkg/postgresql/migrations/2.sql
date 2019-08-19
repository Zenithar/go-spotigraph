-- +migrate Up
CREATE TABLE users (
    id          VARCHAR(32) NOT NULL PRIMARY KEY,
    principal   VARCHAR(150) NOT NULL,
    meta        JSON        NOT NULL
);

-- +migrate Down
DROP TABLE users;
