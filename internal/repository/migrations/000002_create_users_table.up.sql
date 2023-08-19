CREATE TYPE USER_STATUS_ENUM AS ENUM ('active', 'inactive', 'banned', 'deleted');

CREATE TABLE IF NOT EXISTS users
(
    id                BIGSERIAL PRIMARY KEY,
    uuid              UUID                        NOT NULL DEFAULT gen_random_uuid(),
    name              TEXT                        NOT NULL,
    email             TEXT UNIQUE                 NOT NULL,
    email_verified_at TIMESTAMP,
    username          TEXT UNIQUE                 NOT NULL,
    password          TEXT                        NOT NULL,
    strip_id          TEXT                        NOT NULL,
    remember_token    TEXT,
    status            USER_STATUS_ENUM            NOT NULL DEFAULT 'inactive',
    pm_type           TEXT                        NOT NULL DEFAULT 'free',
    pm_last_4         TEXT                        NOT NULL DEFAULT 'free',
    created_at        TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at        TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT NOW(),
    deleted_at        TIMESTAMP(0) WITH TIME ZONE NULL     DEFAULT NULL,
    version           INTEGER                     NOT NULL DEFAULT 1
);

CREATE TRIGGER update_updated_at_on_users_change_trigger
    BEFORE UPDATE
    ON users
    FOR EACH ROW
EXECUTE PROCEDURE update_updated_at_on_change();

CREATE INDEX users_email_status_index ON users (email, status);
CREATE INDEX users_username_status_index ON users (username, status);