-- Modify "categories" table
ALTER TABLE "public"."categories" ADD COLUMN "updated_at" timestamp NOT NULL DEFAULT now();
-- Modify "product_variants" table
ALTER TABLE "public"."product_variants" ALTER COLUMN "price" TYPE numeric(12);
-- Modify "products" table
ALTER TABLE "public"."products" ALTER COLUMN "trending_score" TYPE real, ADD COLUMN "price" numeric(12) NOT NULL, ADD COLUMN "rating" real NOT NULL DEFAULT 0;
-- Modify "reviews" table
ALTER TABLE "public"."reviews" ALTER COLUMN "rating" TYPE smallint;
-- Modify "attribute_values" table
ALTER TABLE "public"."attribute_values" DROP CONSTRAINT "attribute_values_attribute_id_fkey", ADD CONSTRAINT "attribute_values_attribute_id_fkey" FOREIGN KEY ("attribute_id") REFERENCES "public"."attributes" ("id") ON UPDATE CASCADE ON DELETE CASCADE;
-- Modify "orders" table
ALTER TABLE "public"."orders" DROP COLUMN "payment_id";
-- Modify "payments" table
ALTER TABLE "public"."payments" DROP COLUMN "method_id", ADD COLUMN "order_id" integer NOT NULL, ADD CONSTRAINT "payments_order_id_fkey" FOREIGN KEY ("order_id") REFERENCES "public"."orders" ("id") ON UPDATE CASCADE ON DELETE NO ACTION;
-- Drop "payment_methods" table
DROP TABLE "public"."payment_methods";
