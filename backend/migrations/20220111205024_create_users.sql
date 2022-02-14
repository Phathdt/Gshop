-- +goose Up
-- +goose StatementBegin
CREATE SCHEMA "auth";
CREATE TABLE "auth"."users" (
    "id" BIGSERIAL PRIMARY KEY,
    "username" text,
    "password" text,
    "created_at" timestamp(0) DEFAULT now(),
    "updated_at" timestamp(0) DEFAULT now()
);

CREATE UNIQUE INDEX ON "auth"."users" ("username");
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "auth"."users";
DROP SCHEMA "auth"
-- +goose StatementEnd
