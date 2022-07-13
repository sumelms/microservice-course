BEGIN;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE courses
(
    id              bigserial       CONSTRAINT courses_pk PRIMARY KEY,
    uuid            uuid            DEFAULT uuid_generate_v4() NOT NULL,
    code            varchar         NOT NULL UNIQUE,
    "name"          varchar         NOT NULL,
    underline       varchar         NOT NULL,
    image           varchar         NULL,
    image_cover     varchar         NULL,
    excerpt         varchar         NOT NULL,
    description     text            NULL,
    created_at      timestamp       DEFAULT now() NOT NULL,
    updated_at      timestamp       DEFAULT now() NOT NULL,
    deleted_at      timestamp
);

CREATE UNIQUE INDEX courses_id_uindex
    ON courses (id);
CREATE UNIQUE INDEX courses_uuid_uindex
    ON courses (uuid);

COMMIT;
