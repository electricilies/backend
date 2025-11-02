-- sync products.total_purchase

CREATE OR REPLACE FUNCTION ele_sync_product_total_purchase_on_insert()
RETURNS TRIGGER AS $$
BEGIN
  UPDATE products SET total_purchase = total_purchase + NEW.purchase_count WHERE id = NEW.product_id;
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER ele_sync_product_total_purchase_on_insert
AFTER INSERT ON product_variants FOR EACH ROW
EXECUTE FUNCTION ele_sync_product_total_purchase_on_insert();

CREATE OR REPLACE FUNCTION ele_sync_product_total_purchase_on_update()
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

CREATE TRIGGER ele_sync_product_total_purchase_on_update
AFTER UPDATE OF purchase_count ON product_variants FOR EACH ROW
WHEN (old.purchase_count IS DISTINCT FROM new.purchase_count)
EXECUTE FUNCTION ele_sync_product_total_purchase_on_update();

CREATE OR REPLACE FUNCTION ele_sync_product_total_purchase_on_delete()
RETURNS TRIGGER AS $$
BEGIN
  UPDATE products SET total_purchase = GREATEST(0, total_purchase - OLD.purchase_count) WHERE id = OLD.product_id;
  RETURN OLD;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER ele_sync_product_total_purchase_on_delete
AFTER DELETE ON product_variants FOR EACH ROW
EXECUTE FUNCTION ele_sync_product_total_purchase_on_delete();
