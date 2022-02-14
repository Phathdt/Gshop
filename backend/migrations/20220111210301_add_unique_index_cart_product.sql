-- +goose Up
-- +goose StatementBegin
CREATE UNIQUE INDEX "cart_product_unique_index" ON "checkout"."cart_products"
USING BTREE ("cart_id", "product_id");
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX "cart_product_unique_index";
-- +goose StatementEnd
