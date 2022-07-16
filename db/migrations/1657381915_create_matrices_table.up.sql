BEGIN;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE matrices
(
    id              bigserial       CONSTRAINT matrices_pk PRIMARY KEY,
    uuid            uuid            DEFAULT uuid_generate_v4() NOT NULL,
    code            varchar         NOT NULL UNIQUE,
    name            varchar         NOT NULL,
    description     text            NULL,
    course_id       uuid            NOT NULL,
    created_at      timestamp       DEFAULT now() NOT NULL,
    updated_at      timestamp       DEFAULT now() NOT NULL,
    deleted_at      timestamp
);

CREATE UNIQUE INDEX matrices_id_uindex
    ON matrices (id);
CREATE UNIQUE INDEX matrices_uuid_uindex
    ON matrices (uuid);

COMMIT;
