-- +migrate Up
CREATE TABLE IF NOT EXISTS guilds (
    id          VARCHAR(64) NOT NULL PRIMARY KEY,
    name        VARCHAR(50) NOT NULL,
    meta        JSON        NOT NULL,
    member_ids  JSON        NOT NULL
);

-- +migrate Down
DROP TABLE guilds;
