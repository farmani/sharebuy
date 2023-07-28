CREATE TYPE user_status AS ENUM ('active', 'inactive', 'banned');

CREATE TABLE IF NOT EXISTS users(
    id serial PRIMARY KEY,
    name VARCHAR (150) NOT NULL,
    email VARCHAR (300) UNIQUE NOT NULL,
    email_verified_at TIMESTAMP,
    username VARCHAR (50) UNIQUE NOT NULL,
    password VARCHAR (50) NOT NULL,
    remember_token VARCHAR (100),
    status USER_STATUS  NOT NULL DEFAULT 'inactive',
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL DEFAULT NULL
);

CREATE  FUNCTION update_updated_at_on_users_change()
    RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = now();
    RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER update_updated_at_on_users_change
    BEFORE UPDATE
    ON users FOR EACH ROW
    EXECUTE PROCEDURE update_updated_at_on_users_change();