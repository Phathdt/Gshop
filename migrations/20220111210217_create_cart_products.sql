-- +goose Up
-- +goose StatementBegin
CREATE TABLE "checkout"."cart_products" (
    "cart_id" bigint NOT NULL,
    "product_id" bigint NOT NULL,
    "quantity" int DEFAULT 0,
    "total" bigint DEFAULT 0,
    "created_at" timestamp(0) DEFAULT now(),
    "updated_at" timestamp(0) DEFAULT now(),
    PRIMARY KEY(cart_id, product_id)
);

ALTER TABLE "checkout"."cart_products" ADD FOREIGN KEY ("cart_id") REFERENCES "checkout"."carts" ("id") ON DELETE CASCADE;
ALTER TABLE "checkout"."cart_products" ADD FOREIGN KEY ("product_id") REFERENCES "shopping"."products" ("id") ON DELETE CASCADE;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "checkout"."cart_products";
-- +goose StatementEnd
