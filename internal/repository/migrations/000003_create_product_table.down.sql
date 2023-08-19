DROP TRIGGER update_updated_at_on_products_change_trigger on products;
ALTER TABLE products DROP CONSTRAINT IF EXISTS products_price_check;
DROP TABLE IF EXISTS products;
DROP TYPE PRODUCT_STATUS_ENUM;
