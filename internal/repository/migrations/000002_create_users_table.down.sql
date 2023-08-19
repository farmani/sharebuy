DROP TABLE IF EXISTS users;
DROP TRIGGER IF EXISTS update_updated_at_on_users_change_trigger ON users;
DROP TYPE IF EXISTS USER_STATUS_ENUM;