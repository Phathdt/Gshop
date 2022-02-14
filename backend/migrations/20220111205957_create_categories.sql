-- +goose Up
-- +goose StatementBegin
CREATE SCHEMA "shopping";
CREATE TABLE "shopping"."categories" (
    "id" BIGSERIAL PRIMARY KEY,
    "name" text,
    "created_at" timestamp(0) DEFAULT now(),
    "updated_at" timestamp(0) DEFAULT now()
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "shopping"."categories";
DROP SCHEMA "shopping";
-- +goose StatementEnd
