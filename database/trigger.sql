-- products

CREATE OR REPLACE FUNCTION ele_sync_product_price()
RETURNS TRIGGER AS $$
DECLARE min_price DECIMAL(12, 0);
BEGIN
  SELECT MIN(price) INTO min_price FROM product_variants WHERE product_id = NEW.product_id AND deleted_at IS NULL;
  IF min_price IS NOT NULL THEN
    UPDATE products SET price = min_price WHERE id = NEW.product_id;
  END IF;
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE TRIGGER ele_product_price_after_insert
AFTER INSERT ON product_variants FOR EACH ROW
EXECUTE FUNCTION ele_sync_product_price();

CREATE OR REPLACE TRIGGER ele_product_price_after_update
AFTER UPDATE OF price, deleted_at ON product_variants FOR EACH ROW
WHEN (old.price IS DISTINCT FROM new.price OR old.deleted_at IS DISTINCT FROM new.deleted_at)
EXECUTE FUNCTION ele_sync_product_price();

CREATE OR REPLACE TRIGGER ele_product_price_after_delete
AFTER DELETE ON product_variants FOR EACH ROW
EXECUTE FUNCTION ele_sync_product_price();

-- sync products.total_purchase

CREATE OR REPLACE FUNCTION ele_update_product_total_purchase_on_insert()
RETURNS TRIGGER AS $$
BEGIN
  UPDATE products SET total_purchase = total_purchase + NEW.purchase_count WHERE id = NEW.product_id;
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE TRIGGER ele_update_product_total_purchase_on_insert
AFTER INSERT ON product_variants FOR EACH ROW
EXECUTE FUNCTION ele_update_product_total_purchase_on_insert();

CREATE OR REPLACE FUNCTION ele_update_product_total_purchase_on_update()
RETURNS TRIGGER AS $$
DECLARE delta INTEGER;
BEGIN
  delta := NEW.purchase_count - OLD.purchase_count;
  IF delta != 0 THEN
    UPDATE products SET total_purchase = total_purchase + delta WHERE id = NEW.product_id;
  END IF;
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE TRIGGER ele_update_product_total_purchase_on_update
AFTER UPDATE OF purchase_count ON product_variants FOR EACH ROW
WHEN (old.purchase_count IS DISTINCT FROM new.purchase_count)
EXECUTE FUNCTION ele_update_product_total_purchase_on_update();

CREATE OR REPLACE FUNCTION ele_update_product_total_purchase_on_delete()
RETURNS TRIGGER AS $$
BEGIN
  UPDATE products SET total_purchase = GREATEST(0, total_purchase - OLD.purchase_count) WHERE id = OLD.product_id;
  RETURN OLD;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE TRIGGER ele_update_product_total_purchase_on_delete
AFTER DELETE ON product_variants FOR EACH ROW
EXECUTE FUNCTION ele_update_product_total_purchase_on_delete();

-- cart.quantity

CREATE OR REPLACE FUNCTION ele_update_cart_quantity_after_insert_cart_items()
RETURNS TRIGGER AS $$
BEGIN
  UPDATE carts SET quantity = quantity + NEW.quantity WHERE id = NEW.cart_id;
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE TRIGGER ele_update_cart_quantity_after_insert_cart_items
AFTER INSERT ON cart_items FOR EACH ROW
EXECUTE FUNCTION ele_update_cart_quantity_after_insert_cart_items();

CREATE OR REPLACE FUNCTION ele_update_cart_quantity_after_update_cart_items()
RETURNS TRIGGER AS $$
DECLARE delta INTEGER;
BEGIN
  delta := NEW.quantity - OLD.quantity;
  IF delta != 0 THEN
    UPDATE carts SET quantity = quantity + delta WHERE id = NEW.cart_id;
  END IF;
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE TRIGGER ele_update_cart_quantity_after_update_cart_items
AFTER UPDATE OF quantity ON cart_items FOR EACH ROW
WHEN (old.quantity IS DISTINCT FROM new.quantity)
EXECUTE FUNCTION ele_update_cart_quantity_after_update_cart_items();

CREATE OR REPLACE FUNCTION ele_update_cart_quantity_after_delete_cart_items()
RETURNS TRIGGER AS $$
BEGIN
  UPDATE carts SET quantity = GREATEST(0, quantity - OLD.quantity) WHERE id
  = OLD.cart_id;
  RETURN OLD;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE TRIGGER ele_update_cart_quantity_after_delete_cart_items
AFTER DELETE ON cart_items FOR EACH ROW
EXECUTE FUNCTION ele_update_cart_quantity_after_delete_cart_items();
