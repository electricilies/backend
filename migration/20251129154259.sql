-- Create "attributes" table
CREATE TABLE "public"."attributes" (
  "id" uuid NOT NULL,
  "code" character varying(100) NOT NULL,
  "name" text NOT NULL,
  "deleted_at" timestamptz NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "attributes_code_key" UNIQUE ("code"),
  CONSTRAINT "attributes_name_key" UNIQUE ("name")
);
-- Create "attribute_values" table
CREATE TABLE "public"."attribute_values" (
  "id" uuid NOT NULL,
  "attribute_id" uuid NOT NULL,
  "value" text NOT NULL,
  "deleted_at" timestamptz NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "attribute_values_attribute_id_fkey" FOREIGN KEY ("attribute_id") REFERENCES "public"."attributes" ("id") ON UPDATE CASCADE ON DELETE CASCADE
);
-- Create "carts" table
CREATE TABLE "public"."carts" (
  "id" uuid NOT NULL,
  "user_id" uuid NOT NULL,
  "updated_at" timestamptz NOT NULL DEFAULT now(),
  PRIMARY KEY ("id"),
  CONSTRAINT "carts_user_id_key" UNIQUE ("user_id")
);
-- Create "categories" table
CREATE TABLE "public"."categories" (
  "id" uuid NOT NULL,
  "name" text NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT now(),
  "updated_at" timestamptz NOT NULL DEFAULT now(),
  "deleted_at" timestamptz NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "categories_name_key" UNIQUE ("name")
);
-- Create "products" table
CREATE TABLE "public"."products" (
  "id" uuid NOT NULL,
  "name" text NOT NULL,
  "description" text NOT NULL,
  "price" numeric(12) NOT NULL,
  "views_count" integer NOT NULL DEFAULT 0,
  "total_purchase" integer NOT NULL DEFAULT 0,
  "rating" real NOT NULL DEFAULT 0,
  "trending_score" real NOT NULL DEFAULT 0,
  "category_id" uuid NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT now(),
  "updated_at" timestamptz NOT NULL DEFAULT now(),
  "deleted_at" timestamptz NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "products_category_id_fkey" FOREIGN KEY ("category_id") REFERENCES "public"."categories" ("id") ON UPDATE CASCADE ON DELETE NO ACTION
);
-- Create "product_variants" table
CREATE TABLE "public"."product_variants" (
  "id" uuid NOT NULL,
  "sku" text NOT NULL,
  "price" numeric(12) NOT NULL,
  "quantity" integer NOT NULL,
  "purchase_count" integer NOT NULL DEFAULT 0,
  "product_id" uuid NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT now(),
  "updated_at" timestamptz NOT NULL DEFAULT now(),
  "deleted_at" timestamptz NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "product_variants_sku_key" UNIQUE ("sku"),
  CONSTRAINT "product_variants_product_id_fkey" FOREIGN KEY ("product_id") REFERENCES "public"."products" ("id") ON UPDATE CASCADE ON DELETE NO ACTION
);
-- Create "cart_items" table
CREATE TABLE "public"."cart_items" (
  "id" uuid NOT NULL,
  "quantity" integer NOT NULL,
  "cart_id" uuid NOT NULL,
  "product_variant_id" uuid NOT NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "cart_items_cart_id_fkey" FOREIGN KEY ("cart_id") REFERENCES "public"."carts" ("id") ON UPDATE CASCADE ON DELETE NO ACTION,
  CONSTRAINT "cart_items_product_variant_id_fkey" FOREIGN KEY ("product_variant_id") REFERENCES "public"."product_variants" ("id") ON UPDATE CASCADE ON DELETE NO ACTION
);
-- Create "options" table
CREATE TABLE "public"."options" (
  "id" uuid NOT NULL,
  "name" text NOT NULL,
  "product_id" uuid NOT NULL,
  "deleted_at" timestamptz NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "options_product_id_fkey" FOREIGN KEY ("product_id") REFERENCES "public"."products" ("id") ON UPDATE CASCADE ON DELETE NO ACTION
);
-- Create "option_values" table
CREATE TABLE "public"."option_values" (
  "id" uuid NOT NULL,
  "value" text NOT NULL,
  "deleted_at" timestamptz NULL,
  "option_id" uuid NOT NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "option_values_option_id_fkey" FOREIGN KEY ("option_id") REFERENCES "public"."options" ("id") ON UPDATE CASCADE ON DELETE CASCADE
);
-- Create "option_values_product_variants" table
CREATE TABLE "public"."option_values_product_variants" (
  "product_variant_id" uuid NOT NULL,
  "option_value_id" uuid NOT NULL,
  PRIMARY KEY ("product_variant_id", "option_value_id"),
  CONSTRAINT "option_values_product_variants_option_value_id_fkey" FOREIGN KEY ("option_value_id") REFERENCES "public"."option_values" ("id") ON UPDATE CASCADE ON DELETE NO ACTION,
  CONSTRAINT "option_values_product_variants_product_variant_id_fkey" FOREIGN KEY ("product_variant_id") REFERENCES "public"."product_variants" ("id") ON UPDATE CASCADE ON DELETE NO ACTION
);
-- Create "order_providers" table
CREATE TABLE "public"."order_providers" (
  "id" uuid NOT NULL,
  "name" text NOT NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "order_providers_name_key" UNIQUE ("name")
);
-- Create "order_statuses" table
CREATE TABLE "public"."order_statuses" (
  "id" uuid NOT NULL,
  "name" text NOT NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "order_statuses_name_key" UNIQUE ("name")
);
-- Create "orders" table
CREATE TABLE "public"."orders" (
  "id" uuid NOT NULL,
  "address" text NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT now(),
  "updated_at" timestamptz NOT NULL DEFAULT now(),
  "total_amount" numeric(12) NOT NULL,
  "is_paid" boolean NOT NULL DEFAULT false,
  "user_id" uuid NOT NULL,
  "status_id" uuid NOT NULL,
  "provider_id" uuid NOT NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "orders_provider_id_fkey" FOREIGN KEY ("provider_id") REFERENCES "public"."order_providers" ("id") ON UPDATE CASCADE ON DELETE NO ACTION,
  CONSTRAINT "orders_status_id_fkey" FOREIGN KEY ("status_id") REFERENCES "public"."order_statuses" ("id") ON UPDATE CASCADE ON DELETE NO ACTION
);
-- Create "order_items" table
CREATE TABLE "public"."order_items" (
  "id" uuid NOT NULL,
  "quantity" integer NOT NULL,
  "order_id" uuid NOT NULL,
  "price" numeric(12) NOT NULL,
  "product_variant_id" uuid NOT NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "order_items_order_id_fkey" FOREIGN KEY ("order_id") REFERENCES "public"."orders" ("id") ON UPDATE CASCADE ON DELETE NO ACTION,
  CONSTRAINT "order_items_product_variant_id_fkey" FOREIGN KEY ("product_variant_id") REFERENCES "public"."product_variants" ("id") ON UPDATE CASCADE ON DELETE NO ACTION
);
-- Create "product_images" table
CREATE TABLE "public"."product_images" (
  "id" uuid NOT NULL,
  "url" text NOT NULL,
  "order" integer NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT now(),
  "deleted_at" timestamptz NULL,
  "product_id" uuid NOT NULL,
  "product_variant_id" uuid NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "product_images_product_id_fkey" FOREIGN KEY ("product_id") REFERENCES "public"."products" ("id") ON UPDATE CASCADE ON DELETE NO ACTION,
  CONSTRAINT "product_images_product_variant_id_fkey" FOREIGN KEY ("product_variant_id") REFERENCES "public"."product_variants" ("id") ON UPDATE CASCADE ON DELETE NO ACTION
);
-- Create "products_attribute_values" table
CREATE TABLE "public"."products_attribute_values" (
  "product_id" uuid NOT NULL,
  "attribute_value_id" uuid NOT NULL,
  PRIMARY KEY ("product_id", "attribute_value_id"),
  CONSTRAINT "products_attribute_values_attribute_value_id_fkey" FOREIGN KEY ("attribute_value_id") REFERENCES "public"."attribute_values" ("id") ON UPDATE CASCADE ON DELETE NO ACTION,
  CONSTRAINT "products_attribute_values_product_id_fkey" FOREIGN KEY ("product_id") REFERENCES "public"."products" ("id") ON UPDATE CASCADE ON DELETE NO ACTION
);
-- Create "return_request_statuses" table
CREATE TABLE "public"."return_request_statuses" (
  "id" uuid NOT NULL,
  "name" text NOT NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "return_request_statuses_name_key" UNIQUE ("name")
);
-- Create "return_requests" table
CREATE TABLE "public"."return_requests" (
  "id" uuid NOT NULL,
  "reason" character varying(150) NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT now(),
  "updated_at" timestamptz NOT NULL DEFAULT now(),
  "status_id" uuid NOT NULL,
  "user_id" uuid NOT NULL,
  "order_item_id" uuid NOT NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "return_requests_order_item_id_fkey" FOREIGN KEY ("order_item_id") REFERENCES "public"."order_items" ("id") ON UPDATE CASCADE ON DELETE NO ACTION,
  CONSTRAINT "return_requests_status_id_fkey" FOREIGN KEY ("status_id") REFERENCES "public"."return_request_statuses" ("id") ON UPDATE CASCADE ON DELETE NO ACTION
);
-- Create "refund_statuses" table
CREATE TABLE "public"."refund_statuses" (
  "id" uuid NOT NULL,
  "name" text NOT NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "refund_statuses_name_key" UNIQUE ("name")
);
-- Create "refunds" table
CREATE TABLE "public"."refunds" (
  "id" uuid NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT now(),
  "updated_at" timestamptz NOT NULL DEFAULT now(),
  "status_id" uuid NOT NULL,
  "order_item_id" uuid NOT NULL,
  "return_request_id" uuid NOT NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "refunds_order_item_id_fkey" FOREIGN KEY ("order_item_id") REFERENCES "public"."order_items" ("id") ON UPDATE CASCADE ON DELETE NO ACTION,
  CONSTRAINT "refunds_return_request_id_fkey" FOREIGN KEY ("return_request_id") REFERENCES "public"."return_requests" ("id") ON UPDATE CASCADE ON DELETE NO ACTION,
  CONSTRAINT "refunds_status_id_fkey" FOREIGN KEY ("status_id") REFERENCES "public"."refund_statuses" ("id") ON UPDATE CASCADE ON DELETE NO ACTION
);
-- Create "reviews" table
CREATE TABLE "public"."reviews" (
  "id" uuid NOT NULL,
  "rating" smallint NOT NULL,
  "content" text NULL,
  "image_url" text NULL,
  "created_at" timestamptz NOT NULL DEFAULT now(),
  "updated_at" timestamptz NOT NULL DEFAULT now(),
  "deleted_at" timestamptz NULL,
  "user_id" uuid NOT NULL,
  "order_item_id" uuid NOT NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "reviews_order_item_id_fkey" FOREIGN KEY ("order_item_id") REFERENCES "public"."order_items" ("id") ON UPDATE CASCADE ON DELETE NO ACTION,
  CONSTRAINT "reviews_rating_check" CHECK ((rating >= 1) AND (rating <= 5))
);
