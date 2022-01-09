CREATE UNIQUE INDEX "cart_product_unique_index" ON "cart_products"
USING BTREE ("cart_id", "product_id");
