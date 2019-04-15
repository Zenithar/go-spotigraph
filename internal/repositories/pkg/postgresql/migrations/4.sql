-- +migrate Up
CREATE TABLE IF NOT EXISTS squads (
    id                  VARCHAR(32) NOT NULL PRIMARY KEY,
    name                VARCHAR(50) NOT NULL,
    meta                JSON        NOT NULL,
    product_owner_id    VARCHAR(32) NOT NULL,
    member_ids          JSON        NOT NULL
);

-- +migrate Down
DROP TABLE squads;
