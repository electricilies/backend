-- Modify "refund_statuses" table
ALTER TABLE "public"."refund_statuses" ALTER COLUMN "id" DROP DEFAULT, ALTER COLUMN "id" TYPE uuid;
-- Drop sequence used by serial column "id"
DROP SEQUENCE IF EXISTS "public"."refund_statuses_id_seq";
-- Modify "cart_items" table
ALTER TABLE "public"."cart_items" ALTER COLUMN "id" DROP DEFAULT, ALTER COLUMN "id" TYPE uuid, ALTER COLUMN "cart_id" TYPE uuid, ALTER COLUMN "product_variant_id" TYPE uuid;
-- Drop sequence used by serial column "id"
DROP SEQUENCE IF EXISTS "public"."cart_items_id_seq";
-- Modify "return_request_statuses" table
ALTER TABLE "public"."return_request_statuses" ALTER COLUMN "id" DROP DEFAULT, ALTER COLUMN "id" TYPE uuid;
-- Drop sequence used by serial column "id"
DROP SEQUENCE IF EXISTS "public"."return_request_statuses_id_seq";
-- Modify "products_attribute_values" table
ALTER TABLE "public"."products_attribute_values" ALTER COLUMN "product_id" TYPE uuid, ALTER COLUMN "attribute_value_id" TYPE uuid;
-- Modify "product_variants" table
ALTER TABLE "public"."product_variants" ALTER COLUMN "id" DROP DEFAULT, ALTER COLUMN "id" TYPE uuid, ALTER COLUMN "product_id" TYPE uuid, ALTER COLUMN "created_at" TYPE timestamptz, ALTER COLUMN "deleted_at" TYPE timestamptz, ALTER COLUMN "updated_at" TYPE timestamptz;
-- Drop sequence used by serial column "id"
DROP SEQUENCE IF EXISTS "public"."product_variants_id_seq";
-- Modify "option_values_product_variants" table
ALTER TABLE "public"."option_values_product_variants" ALTER COLUMN "product_variant_id" TYPE uuid, ALTER COLUMN "option_value_id" TYPE uuid;
-- Modify "options" table
ALTER TABLE "public"."options" ALTER COLUMN "id" DROP DEFAULT, ALTER COLUMN "id" TYPE uuid, ALTER COLUMN "product_id" TYPE uuid, ALTER COLUMN "deleted_at" TYPE timestamptz;
-- Drop sequence used by serial column "id"
DROP SEQUENCE IF EXISTS "public"."options_id_seq";
-- Modify "attribute_values" table
ALTER TABLE "public"."attribute_values" ALTER COLUMN "id" DROP DEFAULT, ALTER COLUMN "id" TYPE uuid, ALTER COLUMN "attribute_id" TYPE uuid, ADD COLUMN "deleted_at" timestamptz NULL;
-- Drop sequence used by serial column "id"
DROP SEQUENCE IF EXISTS "public"."attribute_values_id_seq";
-- Modify "order_statuses" table
ALTER TABLE "public"."order_statuses" ALTER COLUMN "id" DROP DEFAULT, ALTER COLUMN "id" TYPE uuid;
-- Drop sequence used by serial column "id"
DROP SEQUENCE IF EXISTS "public"."order_statuses_id_seq";
-- Modify "attributes" table
ALTER TABLE "public"."attributes" ALTER COLUMN "id" DROP DEFAULT, ALTER COLUMN "id" TYPE uuid, ALTER COLUMN "deleted_at" TYPE timestamptz, ADD CONSTRAINT "attributes_name_key" UNIQUE ("name");
-- Drop sequence used by serial column "id"
DROP SEQUENCE IF EXISTS "public"."attributes_id_seq";
-- Modify "option_values" table
ALTER TABLE "public"."option_values" DROP CONSTRAINT "option_values_option_id_fkey", ALTER COLUMN "id" DROP DEFAULT, ALTER COLUMN "id" TYPE uuid, ALTER COLUMN "option_id" TYPE uuid, ADD COLUMN "deleted_at" timestamptz NULL, ADD CONSTRAINT "option_values_option_id_fkey" FOREIGN KEY ("option_id") REFERENCES "public"."options" ("id") ON UPDATE CASCADE ON DELETE CASCADE;
-- Drop sequence used by serial column "id"
DROP SEQUENCE IF EXISTS "public"."option_values_id_seq";
-- Create "order_providers" table
CREATE TABLE "public"."order_providers" (
  "id" uuid NOT NULL,
  "name" text NOT NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "order_providers_name_key" UNIQUE ("name")
);
-- Modify "orders" table
ALTER TABLE "public"."orders" DROP CONSTRAINT "orders_user_id_fkey", ALTER COLUMN "id" DROP DEFAULT, ALTER COLUMN "id" TYPE uuid, ALTER COLUMN "created_at" TYPE timestamptz, ALTER COLUMN "updated_at" TYPE timestamptz, ALTER COLUMN "status_id" TYPE uuid, ALTER COLUMN "status_id" DROP DEFAULT, ADD COLUMN "address" text NOT NULL, ADD COLUMN "total_amount" numeric(12) NOT NULL, ADD COLUMN "is_paid" boolean NOT NULL DEFAULT false, ADD COLUMN "provider_id" uuid NOT NULL, ADD CONSTRAINT "orders_provider_id_fkey" FOREIGN KEY ("provider_id") REFERENCES "public"."order_providers" ("id") ON UPDATE CASCADE ON DELETE NO ACTION;
-- Drop sequence used by serial column "id"
DROP SEQUENCE IF EXISTS "public"."orders_id_seq";
-- Modify "order_items" table
ALTER TABLE "public"."order_items" ALTER COLUMN "id" DROP DEFAULT, ALTER COLUMN "id" TYPE uuid, ALTER COLUMN "order_id" TYPE uuid, ALTER COLUMN "product_variant_id" TYPE uuid, ADD COLUMN "price" numeric(12) NOT NULL;
-- Drop sequence used by serial column "id"
DROP SEQUENCE IF EXISTS "public"."order_items_id_seq";
-- Modify "refunds" table
ALTER TABLE "public"."refunds" ALTER COLUMN "id" DROP DEFAULT, ALTER COLUMN "id" TYPE uuid, ALTER COLUMN "created_at" TYPE timestamptz, ALTER COLUMN "updated_at" TYPE timestamptz, ALTER COLUMN "status_id" TYPE uuid, ALTER COLUMN "status_id" DROP DEFAULT, DROP COLUMN "payment_id", ALTER COLUMN "return_request_id" TYPE uuid, ADD COLUMN "order_item_id" uuid NOT NULL, ADD CONSTRAINT "refunds_order_item_id_fkey" FOREIGN KEY ("order_item_id") REFERENCES "public"."order_items" ("id") ON UPDATE CASCADE ON DELETE NO ACTION;
-- Drop sequence used by serial column "id"
DROP SEQUENCE IF EXISTS "public"."refunds_id_seq";
-- Modify "categories" table
ALTER TABLE "public"."categories" ALTER COLUMN "id" DROP DEFAULT, ALTER COLUMN "id" TYPE uuid, ALTER COLUMN "created_at" TYPE timestamptz, ALTER COLUMN "deleted_at" TYPE timestamptz, ALTER COLUMN "updated_at" TYPE timestamptz;
-- Drop sequence used by serial column "id"
DROP SEQUENCE IF EXISTS "public"."categories_id_seq";
-- Modify "products" table
ALTER TABLE "public"."products" ALTER COLUMN "id" DROP DEFAULT, ALTER COLUMN "id" TYPE uuid, ALTER COLUMN "created_at" TYPE timestamptz, ALTER COLUMN "updated_at" TYPE timestamptz, ALTER COLUMN "deleted_at" TYPE timestamptz, ADD COLUMN "category_id" uuid NOT NULL, ADD CONSTRAINT "products_category_id_fkey" FOREIGN KEY ("category_id") REFERENCES "public"."categories" ("id") ON UPDATE CASCADE ON DELETE NO ACTION;
-- Drop sequence used by serial column "id"
DROP SEQUENCE IF EXISTS "public"."products_id_seq";
-- Modify "product_images" table
ALTER TABLE "public"."product_images" ALTER COLUMN "id" DROP DEFAULT, ALTER COLUMN "id" TYPE uuid, ALTER COLUMN "created_at" TYPE timestamptz, ALTER COLUMN "product_variant_id" TYPE uuid, ADD COLUMN "deleted_at" timestamptz NULL, ADD COLUMN "product_id" uuid NOT NULL, ADD CONSTRAINT "product_images_product_id_fkey" FOREIGN KEY ("product_id") REFERENCES "public"."products" ("id") ON UPDATE CASCADE ON DELETE NO ACTION;
-- Drop sequence used by serial column "id"
DROP SEQUENCE IF EXISTS "public"."product_images_id_seq";
-- Modify "reviews" table
ALTER TABLE "public"."reviews" DROP CONSTRAINT "reviews_user_id_fkey", ALTER COLUMN "id" DROP DEFAULT, ALTER COLUMN "id" TYPE uuid, ALTER COLUMN "image_url" DROP NOT NULL, ALTER COLUMN "created_at" TYPE timestamptz, ALTER COLUMN "updated_at" TYPE timestamptz, ALTER COLUMN "deleted_at" TYPE timestamptz, DROP COLUMN "product_id", ADD COLUMN "order_item_id" uuid NOT NULL, ADD CONSTRAINT "reviews_order_item_id_fkey" FOREIGN KEY ("order_item_id") REFERENCES "public"."order_items" ("id") ON UPDATE CASCADE ON DELETE NO ACTION;
-- Drop sequence used by serial column "id"
DROP SEQUENCE IF EXISTS "public"."reviews_id_seq";
-- Drop index "idx_user_id" from table: "carts"
DROP INDEX "public"."idx_user_id";
-- Modify "carts" table
ALTER TABLE "public"."carts" DROP CONSTRAINT "carts_user_id_fkey", ALTER COLUMN "id" DROP DEFAULT, ALTER COLUMN "id" TYPE uuid, ALTER COLUMN "updated_at" TYPE timestamptz;
-- Drop sequence used by serial column "id"
DROP SEQUENCE IF EXISTS "public"."carts_id_seq";
-- Modify "return_requests" table
ALTER TABLE "public"."return_requests" DROP CONSTRAINT "return_requests_user_id_fkey", ALTER COLUMN "id" DROP DEFAULT, ALTER COLUMN "id" TYPE uuid, ALTER COLUMN "created_at" TYPE timestamptz, ALTER COLUMN "updated_at" TYPE timestamptz, ALTER COLUMN "status_id" TYPE uuid, ALTER COLUMN "status_id" DROP DEFAULT, ALTER COLUMN "order_item_id" TYPE uuid;
-- Drop sequence used by serial column "id"
DROP SEQUENCE IF EXISTS "public"."return_requests_id_seq";
-- Drop "products_categories" table
DROP TABLE "public"."products_categories";
-- Drop "payments" table
DROP TABLE "public"."payments";
-- Drop "payment_providers" table
DROP TABLE "public"."payment_providers";
-- Drop "payment_statuses" table
DROP TABLE "public"."payment_statuses";
-- Drop "users" table
DROP TABLE "public"."users";
