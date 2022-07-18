BEGIN;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE subjects
(
    id              bigserial       CONSTRAINT subjects_pk PRIMARY KEY,
    uuid            uuid            DEFAULT uuid_generate_v4() NOT NULL,
    code            varchar         NOT NULL UNIQUE,
    name            varchar         NOT NULL,
    objective       text            NULL,
    credit          decimal         NULL,
    workload        decimal         NULL,
    created_at      timestamp       DEFAULT now() NOT NULL,
    updated_at      timestamp       DEFAULT now() NOT NULL,
    deleted_at      timestamp
);

CREATE UNIQUE INDEX subjects_id_uindex
    ON subjects (id);
CREATE UNIQUE INDEX subjects_uuid_uindex
    ON subjects (uuid);

COMMIT;