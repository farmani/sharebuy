CREATE TABLE IF NOT EXISTS tokens
(
    hash       BYTEA PRIMARY KEY,
    user_id    BIGSERIAL                   NOT NULL REFERENCES users ON DELETE CASCADE ON UPDATE CASCADE,
    expired_at TIMESTAMP(0) WITH TIME ZONE NOT NULL,
    scope      TEXT                        NOT NULL
);