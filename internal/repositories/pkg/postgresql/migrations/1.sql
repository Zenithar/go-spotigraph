-- +migrate Up
CREATE TABLE IF NOT EXISTS chapters (
    id          VARCHAR(32) NOT NULL PRIMARY KEY,
    label       VARCHAR(50) NOT NULL,
    meta        JSON        NOT NULL,
    leader_id   VARCHAR(32) NOT NULL,
    member_ids  JSON        NOT NULL
);

-- +migrate Down
DROP TABLE chapters;
