CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE subscriptions
(
    id              bigserial       CONSTRAINT subscriptions_pk PRIMARY KEY,
    uuid            uuid            DEFAULT uuid_generate_v4() NOT NULL,
    user_id         uuid            NOT NULL,
    course_id       uuid            NOT NULL,
    matrix_id       uuid            NOT NULL,
    valid_until     timestamp,
    created_at      timestamp       DEFAULT now() NOT NULL,
    updated_at      timestamp       DEFAULT now() NOT NULL,
    deleted_at      timestamp
);

CREATE UNIQUE INDEX subscriptions_uuid_uindex
    ON subscriptions (uuid);