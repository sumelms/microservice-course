BEGIN;

CREATE TABLE subscriptions
(
    id              bigint          PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    uuid            uuid            DEFAULT uuid_generate_v4() NOT NULL,
    user_uuid       uuid            NOT NULL,
    course_id       bigint          NOT NULL REFERENCES courses (id),
    matrix_id       bigint          NULL REFERENCES matrices (id),
    role            varchar         NOT NULL,
    expires_at      timestamp       NULL,
    created_at      timestamp       DEFAULT NOW() NOT NULL,
    updated_at      timestamp       DEFAULT NOW() NOT NULL,
    deleted_at      timestamp       NULL,
    reason          varchar         NULL,
    CONSTRAINT reason_required_if_deleted CHECK ((deleted_at IS NULL) OR (reason IS NOT NULL))
);

COMMENT ON COLUMN subscriptions.deleted_at IS 'Timestamp indicating when a subscription was softly deleted, allowing for data recovery. A NULL value means the subscription is active.';
COMMENT ON COLUMN subscriptions.reason IS 'Provides a reason for why the subscription was deleted.';
COMMENT ON COLUMN subscriptions.user_uuid IS 'External Universal Unique Identifier (UUID) for the User, used to cross-reference Users across different microservices.';
COMMENT ON COLUMN subscriptions.course_id IS 'Internal identifier used exclusively within the microservice-course to reference specific courses, indicating a direct relationship with the Course entity.';
COMMENT ON COLUMN subscriptions.matrix_id IS 'Internal identifier used exclusively within the microservice-course to reference specific matrices, indicating a direct relationship with the Matrix entity.';

CREATE UNIQUE INDEX subscriptions_uuid_uindex
    ON subscriptions (uuid);
CREATE UNIQUE INDEX subscriptions_uindex
    ON subscriptions (user_uuid, course_id, matrix_id, deleted_at) NULLS NOT DISTINCT;

COMMIT;
