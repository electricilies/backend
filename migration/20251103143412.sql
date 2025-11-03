-- Modify "attributes" table
ALTER TABLE "public"."attributes" ADD COLUMN "deleted_at" timestamp NULL;
-- Modify "attribute_values" table
ALTER TABLE "public"."attribute_values" DROP CONSTRAINT "attribute_values_attribute_id_fkey", ADD CONSTRAINT "attribute_values_attribute_id_fkey" FOREIGN KEY ("attribute_id") REFERENCES "public"."attributes" ("id") ON UPDATE CASCADE ON DELETE NO ACTION;
-- Modify "carts" table
ALTER TABLE "public"."carts" DROP CONSTRAINT "carts_user_id_fkey", ADD CONSTRAINT "carts_user_id_fkey" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("id") ON UPDATE CASCADE ON DELETE NO ACTION;
-- Create index "idx_user_id" to table: "carts"
CREATE INDEX "idx_user_id" ON "public"."carts" ("user_id");
-- Modify "product_variants" table
ALTER TABLE "public"."product_variants" DROP CONSTRAINT "product_variants_product_id_fkey", ADD COLUMN "updated_at" timestamp NOT NULL DEFAULT now(), ADD CONSTRAINT "product_variants_product_id_fkey" FOREIGN KEY ("product_id") REFERENCES "public"."products" ("id") ON UPDATE CASCADE ON DELETE NO ACTION;
-- Modify "cart_items" table
ALTER TABLE "public"."cart_items" DROP CONSTRAINT "cart_items_cart_id_fkey", DROP CONSTRAINT "cart_items_product_variant_id_fkey", DROP COLUMN "product_id", ADD CONSTRAINT "cart_items_cart_id_fkey" FOREIGN KEY ("cart_id") REFERENCES "public"."carts" ("id") ON UPDATE CASCADE ON DELETE NO ACTION, ADD CONSTRAINT "cart_items_product_variant_id_fkey" FOREIGN KEY ("product_variant_id") REFERENCES "public"."product_variants" ("id") ON UPDATE CASCADE ON DELETE NO ACTION;
-- Modify "options" table
ALTER TABLE "public"."options" DROP CONSTRAINT "options_name_key", DROP CONSTRAINT "options_product_id_fkey", ADD COLUMN "deleted_at" timestamp NULL, ADD CONSTRAINT "options_product_id_fkey" FOREIGN KEY ("product_id") REFERENCES "public"."products" ("id") ON UPDATE CASCADE ON DELETE NO ACTION;
-- Modify "option_values" table
ALTER TABLE "public"."option_values" DROP CONSTRAINT "option_values_option_id_fkey", ADD CONSTRAINT "option_values_option_id_fkey" FOREIGN KEY ("option_id") REFERENCES "public"."options" ("id") ON UPDATE CASCADE ON DELETE NO ACTION;
-- Modify "option_values_product_variants" table
ALTER TABLE "public"."option_values_product_variants" DROP CONSTRAINT "option_values_product_variants_option_value_id_fkey", DROP CONSTRAINT "option_values_product_variants_product_variant_id_fkey", ADD CONSTRAINT "option_values_product_variants_option_value_id_fkey" FOREIGN KEY ("option_value_id") REFERENCES "public"."option_values" ("id") ON UPDATE CASCADE ON DELETE NO ACTION, ADD CONSTRAINT "option_values_product_variants_product_variant_id_fkey" FOREIGN KEY ("product_variant_id") REFERENCES "public"."product_variants" ("id") ON UPDATE CASCADE ON DELETE NO ACTION;
-- Modify "payments" table
ALTER TABLE "public"."payments" DROP COLUMN "payment_method_id", DROP COLUMN "payment_status_id", DROP COLUMN "payment_provider_id", ADD COLUMN "method_id" integer NOT NULL, ADD COLUMN "status_id" integer NOT NULL DEFAULT 1, ADD COLUMN "provider_id" integer NOT NULL, ADD CONSTRAINT "payments_method_id_fkey" FOREIGN KEY ("method_id") REFERENCES "public"."payment_methods" ("id") ON UPDATE CASCADE ON DELETE NO ACTION, ADD CONSTRAINT "payments_provider_id_fkey" FOREIGN KEY ("provider_id") REFERENCES "public"."payment_providers" ("id") ON UPDATE CASCADE ON DELETE NO ACTION, ADD CONSTRAINT "payments_status_id_fkey" FOREIGN KEY ("status_id") REFERENCES "public"."payment_statuses" ("id") ON UPDATE CASCADE ON DELETE NO ACTION;
-- Modify "orders" table
ALTER TABLE "public"."orders" DROP CONSTRAINT "orders_payment_id_fkey", DROP CONSTRAINT "orders_user_id_fkey", DROP COLUMN "order_status_id", ADD COLUMN "status_id" integer NOT NULL DEFAULT 1, ADD CONSTRAINT "orders_payment_id_fkey" FOREIGN KEY ("payment_id") REFERENCES "public"."payments" ("id") ON UPDATE CASCADE ON DELETE NO ACTION, ADD CONSTRAINT "orders_user_id_fkey" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("id") ON UPDATE CASCADE ON DELETE NO ACTION, ADD CONSTRAINT "orders_status_id_fkey" FOREIGN KEY ("status_id") REFERENCES "public"."order_statuses" ("id") ON UPDATE CASCADE ON DELETE NO ACTION;
-- Modify "order_items" table
ALTER TABLE "public"."order_items" DROP CONSTRAINT "order_items_order_id_fkey", DROP CONSTRAINT "order_items_product_variant_id_fkey", DROP COLUMN "product_id", ADD CONSTRAINT "order_items_order_id_fkey" FOREIGN KEY ("order_id") REFERENCES "public"."orders" ("id") ON UPDATE CASCADE ON DELETE NO ACTION, ADD CONSTRAINT "order_items_product_variant_id_fkey" FOREIGN KEY ("product_variant_id") REFERENCES "public"."product_variants" ("id") ON UPDATE CASCADE ON DELETE NO ACTION;
-- Modify "product_images" table
ALTER TABLE "public"."product_images" DROP CONSTRAINT "product_images_product_variant_id_fkey", DROP COLUMN "product_id", ADD CONSTRAINT "product_images_product_variant_id_fkey" FOREIGN KEY ("product_variant_id") REFERENCES "public"."product_variants" ("id") ON UPDATE CASCADE ON DELETE NO ACTION;
-- Create "products_attribute_values" table
CREATE TABLE "public"."products_attribute_values" (
  "product_id" integer NOT NULL,
  "attribute_value_id" integer NOT NULL,
  PRIMARY KEY ("product_id", "attribute_value_id"),
  CONSTRAINT "products_attribute_values_attribute_value_id_fkey" FOREIGN KEY ("attribute_value_id") REFERENCES "public"."attribute_values" ("id") ON UPDATE CASCADE ON DELETE NO ACTION,
  CONSTRAINT "products_attribute_values_product_id_fkey" FOREIGN KEY ("product_id") REFERENCES "public"."products" ("id") ON UPDATE CASCADE ON DELETE NO ACTION
);
-- Modify "categories" table
ALTER TABLE "public"."categories" DROP COLUMN "description";
-- Modify "products_categories" table
ALTER TABLE "public"."products_categories" DROP CONSTRAINT "products_categories_category_id_fkey", DROP CONSTRAINT "products_categories_product_id_fkey", ADD CONSTRAINT "products_categories_category_id_fkey" FOREIGN KEY ("category_id") REFERENCES "public"."categories" ("id") ON UPDATE CASCADE ON DELETE NO ACTION, ADD CONSTRAINT "products_categories_product_id_fkey" FOREIGN KEY ("product_id") REFERENCES "public"."products" ("id") ON UPDATE CASCADE ON DELETE NO ACTION;
-- Modify "return_requests" table
ALTER TABLE "public"."return_requests" DROP CONSTRAINT "return_requests_order_item_id_fkey", DROP CONSTRAINT "return_requests_status_id_fkey", DROP CONSTRAINT "return_requests_user_id_fkey", ALTER COLUMN "status_id" SET DEFAULT 1, ADD CONSTRAINT "return_requests_order_item_id_fkey" FOREIGN KEY ("order_item_id") REFERENCES "public"."order_items" ("id") ON UPDATE CASCADE ON DELETE NO ACTION, ADD CONSTRAINT "return_requests_status_id_fkey" FOREIGN KEY ("status_id") REFERENCES "public"."return_request_statuses" ("id") ON UPDATE CASCADE ON DELETE NO ACTION, ADD CONSTRAINT "return_requests_user_id_fkey" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("id") ON UPDATE CASCADE ON DELETE NO ACTION;
-- Modify "refunds" table
ALTER TABLE "public"."refunds" DROP CONSTRAINT "refunds_payment_id_fkey", DROP CONSTRAINT "refunds_return_request_id_fkey", DROP CONSTRAINT "refunds_status_id_fkey", ALTER COLUMN "status_id" SET DEFAULT 1, ADD CONSTRAINT "refunds_payment_id_fkey" FOREIGN KEY ("payment_id") REFERENCES "public"."payments" ("id") ON UPDATE CASCADE ON DELETE NO ACTION, ADD CONSTRAINT "refunds_return_request_id_fkey" FOREIGN KEY ("return_request_id") REFERENCES "public"."return_requests" ("id") ON UPDATE CASCADE ON DELETE NO ACTION, ADD CONSTRAINT "refunds_status_id_fkey" FOREIGN KEY ("status_id") REFERENCES "public"."refund_statuses" ("id") ON UPDATE CASCADE ON DELETE NO ACTION;
-- Modify "reviews" table
ALTER TABLE "public"."reviews" DROP CONSTRAINT "reviews_product_id_fkey", DROP CONSTRAINT "reviews_user_id_fkey", DROP CONSTRAINT "reviews_rating_check", ADD CONSTRAINT "reviews_rating_check" CHECK ((rating >= 1) AND (rating <= 5)), ADD CONSTRAINT "reviews_product_id_fkey" FOREIGN KEY ("product_id") REFERENCES "public"."products" ("id") ON UPDATE CASCADE ON DELETE NO ACTION, ADD CONSTRAINT "reviews_user_id_fkey" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("id") ON UPDATE CASCADE ON DELETE NO ACTION;
-- Drop "product_attributes_values" table
DROP TABLE "public"."product_attributes_values";
