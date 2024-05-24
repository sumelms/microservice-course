BEGIN;

CREATE TABLE subjects
(
    id              bigint          PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    uuid            uuid            DEFAULT uuid_generate_v4() NOT NULL,
    code            varchar         NOT NULL UNIQUE,
    name            varchar         NOT NULL,
    objective       text            NULL,
    credit          float           NULL,
    workload        float           NULL,
    created_at      timestamp       DEFAULT now() NOT NULL,
    updated_at      timestamp       DEFAULT now() NOT NULL,
    published_at    timestamp       DEFAULT now() NOT NULL,
    deleted_at      timestamp
);

COMMENT ON COLUMN subjects.deleted_at IS 'Timestamp indicating when a subject was softly deleted, allowing for data recovery. A NULL value means the subject is active.';

CREATE UNIQUE INDEX subjects_uuid_uindex
    ON subjects (uuid);

COMMIT;