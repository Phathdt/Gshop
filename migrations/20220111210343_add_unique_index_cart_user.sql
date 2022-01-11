-- +goose Up
-- +goose StatementBegin
CREATE UNIQUE INDEX "user_cart_index" ON "carts" USING BTREE ("user_id");
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX "user_cart_index";
-- +goose StatementEnd
