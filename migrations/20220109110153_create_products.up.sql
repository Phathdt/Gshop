CREATE TABLE "products" (
    "id" BIGSERIAL PRIMARY KEY,
    "sku" text,
    "name" text,
    "price" bigint,
    "category_id" bigint NOT NULL,
    "created_at" timestamp(0) DEFAULT now(),
    "updated_at" timestamp(0) DEFAULT now()
);

ALTER TABLE "products" ADD FOREIGN KEY ("category_id") REFERENCES "categories" ("id");
