BEGIN;

CREATE TABLE courses
(
    id              integer         PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    uuid            uuid            DEFAULT uuid_generate_v4() NOT NULL,
    code            varchar         NOT NULL,
    name            varchar         NOT NULL,
    underline       varchar         NOT NULL,
    image           varchar         NULL,
    image_cover     varchar         NULL,
    excerpt         varchar         NOT NULL,
    description     text            NULL,
    created_at      timestamp       DEFAULT now() NOT NULL,
    updated_at      timestamp       DEFAULT now() NOT NULL,
    deleted_at      timestamp
);

COMMENT ON COLUMN courses.deleted_at IS 'Timestamp indicating when a course was softly deleted, allowing for data recovery and historical queries without permanently removing the record from the database. A NULL value means the subscription is active.';

CREATE UNIQUE INDEX courses_uuid_uindex
    ON courses (uuid);
CREATE UNIQUE INDEX courses_code_uindex
    ON courses (code, deleted_at) NULLS NOT DISTINCT;

COMMIT;
