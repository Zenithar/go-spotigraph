-- +migrate Up
CREATE TABLE persons (
    id          VARCHAR(32) NOT NULL PRIMARY KEY,
    principal   VARCHAR(150) NOT NULL,
    meta        JSON        NOT NULL,
    UNIQUE (principal)
);

-- +migrate Down
DROP TABLE persons;
