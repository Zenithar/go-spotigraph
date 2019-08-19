-- +migrate Up
CREATE TABLE chapters (
    id          VARCHAR(32) NOT NULL PRIMARY KEY,
    label       VARCHAR(50) NOT NULL,
    meta        JSON        NOT NULL,
    leader_id   VARCHAR(32) NOT NULL
);

-- +migrate Down
DROP TABLE chapters;
