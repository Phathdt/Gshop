CREATE TABLE "products" (
    "id" BIGSERIAL PRIMARY KEY,
    "sku" text,
    "name" text,
    "price" bigint,
    "category" text,
    "created_at" timestamp(0) DEFAULT now(),
    "updated_at" timestamp(0) DEFAULT now()
);

CREATE UNIQUE INDEX ON "products" ("sku");
