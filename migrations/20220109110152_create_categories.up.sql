CREATE TABLE "categories" (
    "id" BIGSERIAL PRIMARY KEY,
    "name" text,
    "created_at" timestamp(0) DEFAULT now(),
    "updated_at" timestamp(0) DEFAULT now()
);
