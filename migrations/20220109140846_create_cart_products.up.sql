CREATE TABLE "cart_products" (
    "cart_id" bigint NOT NULL,
    "product_id" bigint NOT NULL,
    "quantity" int DEFAULT 0,
    "total" bigint DEFAULT 0,
    "created_at" timestamp(0) DEFAULT now(),
    "updated_at" timestamp(0) DEFAULT now(),
    PRIMARY KEY(cart_id, product_id)
);

ALTER TABLE "cart_products" ADD FOREIGN KEY ("cart_id") REFERENCES "carts" ("id") ON DELETE CASCADE;
ALTER TABLE "cart_products" ADD FOREIGN KEY ("product_id") REFERENCES "products" ("id") ON DELETE CASCADE;
