CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE courses
(
    id              bigserial       CONSTRAINT courses_pk PRIMARY KEY,
    uuid            uuid            DEFAULT uuid_generate_v4() NOT NULL,
    title           varchar         NOT NULL,
    subtitle        varchar         NOT NULL,
    excerpt         varchar         NOT NULL,
    description     text,
    created_at      timestamp       DEFAULT now() NOT NULL,
    updated_at      timestamp       DEFAULT now() NOT NULL,
    deleted_at      timestamp
);

CREATE UNIQUE INDEX courses_uuid_uindex
    ON courses (uuid);