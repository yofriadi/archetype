-- Add new schema named "merchandise"
CREATE SCHEMA "merchandise";
-- Add new schema named "public"
CREATE SCHEMA IF NOT EXISTS "public";
-- Set comment to schema: "public"
COMMENT ON SCHEMA "public" IS 'standard public schema';
-- Create "design_requests" table
CREATE TABLE "merchandise"."design_requests" ("id" bigint NOT NULL GENERATED ALWAYS AS IDENTITY, "number" character varying(255) NOT NULL, "sketch_due_date" date NULL, "brj_due_date" date NULL, "brand" smallint NOT NULL, "price_class" smallint NOT NULL, "product_category" smallint NOT NULL, "market_segment" smallint NOT NULL, "note" character varying(1000) NULL, "stage" smallint NULL DEFAULT 0, "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP, "updated_at" timestamp NULL, "deleted_at" timestamp NULL, PRIMARY KEY ("id"));
-- Create index "design_requests_number_idx" to table: "design_requests"
CREATE UNIQUE INDEX "design_requests_number_idx" ON "merchandise"."design_requests" ("number");
-- Create "design_request_customers" table
CREATE TABLE "merchandise"."design_request_customers" ("id" bigint NOT NULL GENERATED ALWAYS AS IDENTITY, "design_request_id" bigint NOT NULL, "store_id" smallint NULL, "customer_id" integer NOT NULL, "designer_id" smallint NULL, "po_number" character varying(256) NOT NULL, "product_group" smallint NOT NULL, "diamond_brand" smallint NOT NULL, "ring_size" smallint NULL, "width" smallint NULL, "clarity" smallint NOT NULL, "center_stone_size" smallint NOT NULL, "item" smallint NOT NULL, "item_color" smallint NULL, "color" smallint NULL, "size" smallint NULL, "height" smallint NULL, "weight" smallint NULL, PRIMARY KEY ("id"), CONSTRAINT "design_request_customers_design_request_id_fkey" FOREIGN KEY ("design_request_id") REFERENCES "merchandise"."design_requests" ("id") ON UPDATE NO ACTION ON DELETE CASCADE);
-- Create index "design_request_customers_design_request_id_idx" to table: "design_request_customers"
CREATE UNIQUE INDEX "design_request_customers_design_request_id_idx" ON "merchandise"."design_request_customers" ("design_request_id");
-- Create "design_request_internals" table
CREATE TABLE "merchandise"."design_request_internals" ("id" bigint NOT NULL GENERATED ALWAYS AS IDENTITY, "design_request_id" bigint NOT NULL, "launch_date" date NOT NULL, "collection" character varying(255) NOT NULL, "project_type" smallint NOT NULL, "design_classification" smallint NOT NULL, "total_item" smallint NOT NULL, "total" smallint NOT NULL, "target_gender" smallint NOT NULL, "target_age" smallint NOT NULL, PRIMARY KEY ("id"), CONSTRAINT "design_request_internals_design_request_id_fkey" FOREIGN KEY ("design_request_id") REFERENCES "merchandise"."design_requests" ("id") ON UPDATE NO ACTION ON DELETE CASCADE);
-- Create index "design_request_internals_design_request_id_idx" to table: "design_request_internals"
CREATE UNIQUE INDEX "design_request_internals_design_request_id_idx" ON "merchandise"."design_request_internals" ("design_request_id");
