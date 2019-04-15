-- +migrate Up
CREATE TABLE IF NOT EXISTS tribes (
    id          VARCHAR(32) NOT NULL PRIMARY KEY,
    name        VARCHAR(50) NOT NULL,
    meta        JSON        NOT NULL,
    squad_ids   JSON        NOT NULL
);

-- +migrate Down
DROP TABLE tribes;
