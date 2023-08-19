CREATE TABLE IF NOT EXISTS password_reset_tokens
(
    id         BIGSERIAL PRIMARY KEY,
    email      TEXT UNIQUE                 NOT NULL,
    token      TEXT                        NOT NULL DEFAULT '',
    created_at TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT NOW(),
    version    INTEGER                     NOT NULL DEFAULT 1
);