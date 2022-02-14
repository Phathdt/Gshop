-- +goose Up
-- +goose StatementBegin
CREATE UNIQUE INDEX "user_cart_index" ON "checkout"."carts" USING BTREE ("user_id");
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX "checkout"."user_cart_index";
-- +goose StatementEnd
