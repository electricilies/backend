-- Create "attributes" table
CREATE TABLE "public"."attributes" (
  "id" serial NOT NULL,
  "code" character varying(100) NOT NULL,
  "name" text NOT NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "attributes_code_key" UNIQUE ("code")
);
-- Create "attribute_values" table
CREATE TABLE "public"."attribute_values" (
  "id" serial NOT NULL,
  "attribute_id" integer NOT NULL,
  "value" text NOT NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "attribute_values_attribute_id_fkey" FOREIGN KEY ("attribute_id") REFERENCES "public"."attributes" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Modify "users" table
ALTER TABLE "public"."users" DROP COLUMN "avatar", DROP COLUMN "first_name", DROP COLUMN "last_name", DROP COLUMN "username", DROP COLUMN "email", DROP COLUMN "birthday", DROP COLUMN "phone_number", DROP COLUMN "created_at", DROP COLUMN "deleted_at";
-- Create "carts" table
CREATE TABLE "public"."carts" (
  "id" serial NOT NULL,
  "user_id" uuid NOT NULL,
  "updated_at" timestamp NOT NULL DEFAULT now(),
  PRIMARY KEY ("id"),
  CONSTRAINT "carts_user_id_key" UNIQUE ("user_id"),
  CONSTRAINT "carts_user_id_fkey" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create "products" table
CREATE TABLE "public"."products" (
  "id" serial NOT NULL,
  "name" text NOT NULL,
  "description" text NOT NULL,
  "views_count" integer NOT NULL DEFAULT 0,
  "total_purchase" integer NOT NULL DEFAULT 0,
  "trending_score" double precision NOT NULL DEFAULT 0,
  "created_at" timestamp NOT NULL DEFAULT now(),
  "updated_at" timestamp NOT NULL DEFAULT now(),
  "deleted_at" timestamp NULL,
  PRIMARY KEY ("id")
);
-- Create "product_variants" table
CREATE TABLE "public"."product_variants" (
  "id" serial NOT NULL,
  "sku" text NOT NULL,
  "price" numeric NOT NULL,
  "quantity" integer NOT NULL,
  "purchase_count" integer NOT NULL DEFAULT 0,
  "product_id" integer NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT now(),
  "deleted_at" timestamp NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "product_variants_sku_key" UNIQUE ("sku"),
  CONSTRAINT "product_variants_product_id_fkey" FOREIGN KEY ("product_id") REFERENCES "public"."products" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create "cart_items" table
CREATE TABLE "public"."cart_items" (
  "id" serial NOT NULL,
  "quantity" integer NOT NULL,
  "cart_id" integer NOT NULL,
  "product_id" integer NOT NULL,
  "product_variant_id" integer NOT NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "cart_items_cart_id_fkey" FOREIGN KEY ("cart_id") REFERENCES "public"."carts" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "cart_items_product_id_fkey" FOREIGN KEY ("product_id") REFERENCES "public"."products" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "cart_items_product_variant_id_fkey" FOREIGN KEY ("product_variant_id") REFERENCES "public"."product_variants" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create "options" table
CREATE TABLE "public"."options" (
  "id" serial NOT NULL,
  "name" text NOT NULL,
  "product_id" integer NOT NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "options_name_key" UNIQUE ("name"),
  CONSTRAINT "options_product_id_fkey" FOREIGN KEY ("product_id") REFERENCES "public"."products" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create "option_values" table
CREATE TABLE "public"."option_values" (
  "id" serial NOT NULL,
  "value" text NOT NULL,
  "option_id" integer NOT NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "option_values_option_id_fkey" FOREIGN KEY ("option_id") REFERENCES "public"."options" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create "option_values_product_variants" table
CREATE TABLE "public"."option_values_product_variants" (
  "product_variant_id" integer NOT NULL,
  "option_value_id" integer NOT NULL,
  PRIMARY KEY ("product_variant_id", "option_value_id"),
  CONSTRAINT "option_values_product_variants_option_value_id_fkey" FOREIGN KEY ("option_value_id") REFERENCES "public"."option_values" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "option_values_product_variants_product_variant_id_fkey" FOREIGN KEY ("product_variant_id") REFERENCES "public"."product_variants" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create "order_statuses" table
CREATE TABLE "public"."order_statuses" (
  "id" serial NOT NULL,
  "name" text NOT NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "order_statuses_name_key" UNIQUE ("name")
);
-- Create "payment_methods" table
CREATE TABLE "public"."payment_methods" (
  "id" serial NOT NULL,
  "name" text NOT NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "payment_methods_name_key" UNIQUE ("name")
);
-- Create "payment_providers" table
CREATE TABLE "public"."payment_providers" (
  "id" serial NOT NULL,
  "name" text NOT NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "payment_providers_name_key" UNIQUE ("name")
);
-- Create "payment_statuses" table
CREATE TABLE "public"."payment_statuses" (
  "id" serial NOT NULL,
  "name" text NOT NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "payment_statuses_name_key" UNIQUE ("name")
);
-- Create "payments" table
CREATE TABLE "public"."payments" (
  "id" serial NOT NULL,
  "amount" numeric NOT NULL,
  "updated_at" timestamp NOT NULL DEFAULT now(),
  "payment_method_id" integer NOT NULL,
  "payment_status_id" integer NOT NULL,
  "payment_provider_id" integer NOT NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "payments_payment_method_id_fkey" FOREIGN KEY ("payment_method_id") REFERENCES "public"."payment_methods" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "payments_payment_provider_id_fkey" FOREIGN KEY ("payment_provider_id") REFERENCES "public"."payment_providers" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "payments_payment_status_id_fkey" FOREIGN KEY ("payment_status_id") REFERENCES "public"."payment_statuses" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create "orders" table
CREATE TABLE "public"."orders" (
  "id" serial NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT now(),
  "updated_at" timestamp NOT NULL DEFAULT now(),
  "user_id" uuid NOT NULL,
  "order_status_id" integer NOT NULL,
  "payment_id" integer NOT NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "orders_order_status_id_fkey" FOREIGN KEY ("order_status_id") REFERENCES "public"."order_statuses" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "orders_payment_id_fkey" FOREIGN KEY ("payment_id") REFERENCES "public"."payments" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "orders_user_id_fkey" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create "order_items" table
CREATE TABLE "public"."order_items" (
  "id" serial NOT NULL,
  "quantity" integer NOT NULL,
  "order_id" integer NOT NULL,
  "product_id" integer NOT NULL,
  "product_variant_id" integer NOT NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "order_items_order_id_fkey" FOREIGN KEY ("order_id") REFERENCES "public"."orders" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "order_items_product_id_fkey" FOREIGN KEY ("product_id") REFERENCES "public"."products" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "order_items_product_variant_id_fkey" FOREIGN KEY ("product_variant_id") REFERENCES "public"."product_variants" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create "product_attributes_values" table
CREATE TABLE "public"."product_attributes_values" (
  "product_id" integer NOT NULL,
  "attribute_value_id" integer NOT NULL,
  PRIMARY KEY ("product_id", "attribute_value_id"),
  CONSTRAINT "product_attributes_values_attribute_value_id_fkey" FOREIGN KEY ("attribute_value_id") REFERENCES "public"."attribute_values" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "product_attributes_values_product_id_fkey" FOREIGN KEY ("product_id") REFERENCES "public"."products" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create "product_images" table
CREATE TABLE "public"."product_images" (
  "id" serial NOT NULL,
  "url" text NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT now(),
  "order" integer NOT NULL,
  "product_id" integer NOT NULL,
  "product_variant_id" integer NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "product_images_product_id_fkey" FOREIGN KEY ("product_id") REFERENCES "public"."products" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "product_images_product_variant_id_fkey" FOREIGN KEY ("product_variant_id") REFERENCES "public"."product_variants" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create "categories" table
CREATE TABLE "public"."categories" (
  "id" serial NOT NULL,
  "name" text NOT NULL,
  "description" text NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT now(),
  "deleted_at" timestamp NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "categories_name_key" UNIQUE ("name")
);
-- Create "products_categories" table
CREATE TABLE "public"."products_categories" (
  "product_id" integer NOT NULL,
  "category_id" integer NOT NULL,
  PRIMARY KEY ("product_id", "category_id"),
  CONSTRAINT "products_categories_category_id_fkey" FOREIGN KEY ("category_id") REFERENCES "public"."categories" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "products_categories_product_id_fkey" FOREIGN KEY ("product_id") REFERENCES "public"."products" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create "return_request_statuses" table
CREATE TABLE "public"."return_request_statuses" (
  "id" serial NOT NULL,
  "name" text NOT NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "return_request_statuses_name_key" UNIQUE ("name")
);
-- Create "return_requests" table
CREATE TABLE "public"."return_requests" (
  "id" serial NOT NULL,
  "reason" character varying(150) NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT now(),
  "updated_at" timestamp NOT NULL DEFAULT now(),
  "status_id" integer NOT NULL,
  "user_id" uuid NOT NULL,
  "order_item_id" integer NOT NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "return_requests_order_item_id_fkey" FOREIGN KEY ("order_item_id") REFERENCES "public"."order_items" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "return_requests_status_id_fkey" FOREIGN KEY ("status_id") REFERENCES "public"."return_request_statuses" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "return_requests_user_id_fkey" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create "refund_statuses" table
CREATE TABLE "public"."refund_statuses" (
  "id" serial NOT NULL,
  "name" text NOT NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "refund_statuses_name_key" UNIQUE ("name")
);
-- Create "refunds" table
CREATE TABLE "public"."refunds" (
  "id" serial NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT now(),
  "updated_at" timestamp NOT NULL DEFAULT now(),
  "status_id" integer NOT NULL,
  "payment_id" integer NOT NULL,
  "return_request_id" integer NOT NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "refunds_payment_id_fkey" FOREIGN KEY ("payment_id") REFERENCES "public"."payments" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "refunds_return_request_id_fkey" FOREIGN KEY ("return_request_id") REFERENCES "public"."return_requests" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "refunds_status_id_fkey" FOREIGN KEY ("status_id") REFERENCES "public"."refund_statuses" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create "reviews" table
CREATE TABLE "public"."reviews" (
  "id" serial NOT NULL,
  "rating" integer NOT NULL,
  "content" text NULL,
  "image_url" text NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT now(),
  "updated_at" timestamp NOT NULL DEFAULT now(),
  "deleted_at" timestamp NULL,
  "user_id" uuid NOT NULL,
  "product_id" integer NOT NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "reviews_product_id_fkey" FOREIGN KEY ("product_id") REFERENCES "public"."products" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "reviews_user_id_fkey" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "reviews_rating_check" CHECK ((rating > 0) AND (rating <= 5))
);
