CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE "persons" (
    person_id VARCHAR(50) NOT NULL,
    principal varchar(255) NOT NULL,
    locked boolean NOT NULL DEFAULT false,
    created_at TIMESTAMP NOT NULL DEFAULT (now() AT TIME ZONE 'utc'),
    PRIMARY KEY(person_id),
    UNIQUE(principal)
);
