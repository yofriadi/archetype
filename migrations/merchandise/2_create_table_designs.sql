CREATE TABLE IF NOT EXISTS "merchandise"."designs" (
    "id" bigint NOT NULL PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    "parent_id" bigint REFERENCES merchandise.designs (id) ON DELETE CASCADE,
    "designer_id" bigint NOT NULL,
    "item" smallint NOT NULL,
    "color" smallint NOT NULL,
    "product_category" smallint NOT NULL,
    "stones" json[] NOT NULL,
    "metal" smallint NOT NULL,
    "metal_rate" float NOT NULL,
    "image_url" varchar(255) NOT NULL,
    "created_at" timestamp DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamp,
    "deleted_at" timestamp
);

CREATE INDEX CONCURRENTLY IF NOT EXISTS "designs_designer_id_idx" ON "merchandise"."designs" (designer_id);

CREATE OR REPLACE TRIGGER on_update
    BEFORE UPDATE ON "merchandise"."designs"
    FOR EACH ROW
    EXECUTE PROCEDURE "merchandise".fill_updated_at();
