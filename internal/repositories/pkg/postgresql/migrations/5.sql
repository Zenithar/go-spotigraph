-- +migrate Up
CREATE TABLE tribes (
    id          VARCHAR(32) NOT NULL PRIMARY KEY,
    label       VARCHAR(50) NOT NULL,
    meta        JSON        NOT NULL,
    squad_ids   JSON        NOT NULL,
    leader_id   VARCHAR(32)
);

-- +migrate Down
DROP TABLE tribes;
