BEGIN;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE matrix_subjects
(
    id              bigserial       CONSTRAINT matrix_subjects_pk PRIMARY KEY,
    subject_id      bigserial       NOT NULL,
    matrix_id       bigserial       NOT NULL,
    group           varchar,
    is_required     boolean         NULL DEFAULT TRUE,
    created_at      timestamp       DEFAULT now() NOT NULL,
    updated_at      timestamp       DEFAULT now() NOT NULL,
    deleted_at      timestamp
);

CREATE UNIQUE INDEX matrix_subjects_id_uindex
    ON matrix_subjects (id);

COMMIT;