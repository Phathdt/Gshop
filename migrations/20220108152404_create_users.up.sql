CREATE TABLE "users" (
    "id" BIGSERIAL PRIMARY KEY,
    "username" text,
    "password" text,
    "created_at" timestamp(0) DEFAULT now(),
    "updated_at" timestamp(0) DEFAULT now()
);

CREATE UNIQUE INDEX ON "users" ("username");
