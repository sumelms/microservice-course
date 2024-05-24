BEGIN;

CREATE TABLE matrices
(
    id              bigint          PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    uuid            uuid            DEFAULT uuid_generate_v4() NOT NULL,
    code            varchar         NOT NULL UNIQUE,
    name            varchar         NOT NULL,
    description     text            NULL,
    course_id       bigint          NOT NULL REFERENCES courses (id),
    created_at      timestamp       DEFAULT NOW() NOT NULL,
    updated_at      timestamp       DEFAULT NOW() NOT NULL,
    deleted_at      timestamp       NULL
);

COMMENT ON COLUMN matrices.deleted_at IS 'Timestamp indicating when a matrix was softly deleted, allowing for data recovery. A NULL value means the matrix is active.';
COMMENT ON COLUMN matrices.course_id IS 'Internal identifier used exclusively within the microservice-course to reference specific courses, indicating a direct relationship with the Course entity.';

CREATE UNIQUE INDEX matrices_uuid_uindex
    ON matrices (uuid);

COMMIT;
