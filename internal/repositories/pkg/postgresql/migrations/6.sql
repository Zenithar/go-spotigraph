-- +migrate Up
CREATE TABLE memberships (
    id          VARCHAR(32) NOT NULL PRIMARY KEY,
    person_id   VARCHAR(32) NOT NULL,
    group_type  VARCHAR(32) NOT NULL,
    group_id    VARCHAR(32) NOT NULL,
    UNIQUE (person_id,group_type,group_id)
);

-- +migrate Down
DROP TABLE memberships;