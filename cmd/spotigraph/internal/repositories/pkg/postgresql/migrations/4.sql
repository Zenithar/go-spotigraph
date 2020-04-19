-- +migrate Up
CREATE TABLE squads (
    id                  VARCHAR(32) NOT NULL PRIMARY KEY,
    label               VARCHAR(50) NOT NULL,
    meta                JSON        NOT NULL,
    product_owner_id    VARCHAR(32) NOT NULL,
    UNIQUE (label)
);

-- +migrate Down
DROP TABLE squads;
