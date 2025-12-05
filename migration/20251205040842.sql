-- Modify "orders" table
ALTER TABLE "public"."orders" ADD COLUMN "recipient_name" text NOT NULL, ADD COLUMN "phone_number" text NOT NULL;
