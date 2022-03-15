CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE "persons" (
    tenant_id VARCHAR(50) NOT NULL,
    person_id VARCHAR(50) NOT NULL,
    principal varchar(255) NOT NULL,
    locked boolean NOT NULL DEFAULT false,
    created_at TIMESTAMP NOT NULL DEFAULT (now() AT TIME ZONE 'utc'),
    PRIMARY KEY(tenant_id, person_id),
    UNIQUE(tenant_id, principal)
);
