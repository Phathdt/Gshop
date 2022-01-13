-- +goose Up
-- +goose StatementBegin
CREATE TABLE "shopping"."products" (
    "id" BIGSERIAL PRIMARY KEY,
    "sku" text,
    "name" text,
    "price" bigint,
    "category_id" bigint NOT NULL,
    "created_at" timestamp(0) DEFAULT now(),
    "updated_at" timestamp(0) DEFAULT now()
);

ALTER TABLE "shopping"."products" ADD FOREIGN KEY ("category_id") REFERENCES "shopping"."categories" ("id");
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "shopping"."products";
-- +goose StatementEnd
