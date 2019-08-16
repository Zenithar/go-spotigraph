-- +migrate Up
CREATE TABLE IF NOT EXISTS guilds (
    id              VARCHAR(32) NOT NULL PRIMARY KEY,
    label           VARCHAR(50) NOT NULL,
    meta            JSON        NOT NULL,
    member_ids      JSON        NOT NULL,
    leader_id       VARCHAR(32)
);

-- +migrate Down
DROP TABLE guilds;
