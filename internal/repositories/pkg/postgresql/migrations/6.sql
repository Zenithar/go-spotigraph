-- +migrate Up
CREATE TABLE memberships (
    id          VARCHAR(32) NOT NULL PRIMARY KEY,
    user_id     VARCHAR(32) NOT NULL,
    group_type  VARCHAR(32) NOT NULL,
    group_id    VARCHAR(32) NOT NULL
);

-- +migrate Down
DROP TABLE memberships;