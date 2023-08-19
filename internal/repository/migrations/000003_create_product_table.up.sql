CREATE TYPE PRODUCT_STATUS_ENUM AS ENUM ('active', 'inactive');

CREATE TABLE IF NOT EXISTS products
(
    id         BIGSERIAL PRIMARY KEY,
    uuid       UUID                        NOT NULL DEFAULT gen_random_uuid(),
    title      TEXT                        NOT NULL,
    price      INTEGER                     NOT NULL DEFAULT 1,
    currency   TEXT                        NOT NULL DEFAULT 'USD',
    url        TEXT                        NOT NULL DEFAULT '',
    images     INTEGER[]                   NOT NULL DEFAULT '{}',
    status     PRODUCT_STATUS_ENUM         NOT NULL DEFAULT 'inactive',
    created_at TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP(0) WITH TIME ZONE NULL     DEFAULT NULL,
    version    INTEGER                     NOT NULL DEFAULT 1
);

CREATE TRIGGER update_updated_at_on_products_change_trigger
    BEFORE UPDATE
    ON products
    FOR EACH ROW
EXECUTE PROCEDURE update_updated_at_on_change();

ALTER TABLE products
    ADD CONSTRAINT products_price_check CHECK (price > 0);

CREATE INDEX products_uuid_idx ON products (uuid);

ALTER TABLE products ADD CONSTRAINT products_unique_uuid UNIQUE (uuid);
