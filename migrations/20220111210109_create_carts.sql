-- +goose Up
-- +goose StatementBegin
CREATE TABLE "carts" (
    "id" BIGSERIAL PRIMARY KEY,
    "total" bigint DEFAULT 0,
    "user_id" bigint NOT NULL,
    "created_at" timestamp(0) DEFAULT now(),
    "updated_at" timestamp(0) DEFAULT now()
);

ALTER TABLE "carts" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "carts";
-- +goose StatementEnd
