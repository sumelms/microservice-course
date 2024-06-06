BEGIN;

CREATE TABLE matrix_subjects
(
    id              bigint          PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    uuid            uuid            DEFAULT uuid_generate_v4() NOT NULL,
    subject_id      bigint          NOT NULL REFERENCES subjects (id),
    matrix_id       bigint          NOT NULL REFERENCES matrices (id),
    is_required     boolean         NOT NULL DEFAULT TRUE,
    "group"         varchar         NOT NULL,
    created_at      timestamp       DEFAULT now() NOT NULL,
    updated_at      timestamp       DEFAULT now() NOT NULL,
    deleted_at      timestamp
);

COMMENT ON COLUMN matrix_subjects.deleted_at IS 'Timestamp indicating when a matrix subject was softly deleted, allowing for data recovery. A NULL value means the matrix subject is active.';

CREATE UNIQUE INDEX matrix_subjects_id_uindex
    ON matrix_subjects (id);

CREATE UNIQUE INDEX matrix_subjects_uindex
    ON matrix_subjects (subject_id, matrix_id, deleted_at) NULLS NOT DISTINCT;

COMMIT;
