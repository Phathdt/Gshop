-- +goose Up
-- +goose StatementBegin
CREATE SCHEMA "checkout";
CREATE TABLE "checkout"."carts" (
    "id" BIGSERIAL PRIMARY KEY,
    "total" bigint DEFAULT 0,
    "user_id" bigint NOT NULL,
    "created_at" timestamp(0) DEFAULT now(),
    "updated_at" timestamp(0) DEFAULT now()
);

ALTER TABLE "checkout"."carts" ADD FOREIGN KEY ("user_id") REFERENCES "auth"."users" ("id");
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "checkout"."carts";
DROP SCHEMA "checkout";
-- +goose StatementEnd
