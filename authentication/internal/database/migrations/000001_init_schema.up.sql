CREATE TABLE "users" (
    "id" BIGSERIAL  PRIMARY KEY,
    "name" VARCHAR NOT NULL,
    "surname" VARCHAR UNIQUE NOT NULL,
    "email" VARCHAR NOT NULL,
    "country" VARCHAR NOT NULL,
    "hashed_password" VARCHAR NOT NULL,
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT (NOW())
);