BEGIN;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE subscriptions
(
    id              bigserial       CONSTRAINT subscriptions_pk PRIMARY KEY,
    user_id         uuid            NOT NULL,
    course_id       uuid            NOT NULL,
    matrix_id       uuid            NULL,
    role            varchar         NULL,
    expires_at      timestamp       NULL,
    created_at      timestamp       DEFAULT now() NOT NULL,
    updated_at      timestamp       DEFAULT now() NOT NULL,
    deleted_at      timestamp
);

CREATE UNIQUE INDEX subscriptions_id_uindex
    ON subscriptions (id);

COMMIT;
