CREATE TABLE "subcategories" (
    "id" BIGSERIAL  PRIMARY KEY,
    "category_id" BIGINT NOT NULL,
    "name" VARCHAR NOT NULL,
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT (NOW())
);

CREATE TABLE "categories" (
    "id" BIGSERIAL  PRIMARY KEY,
    "user_id" BIGINT NOT NULL,
    "name" VARCHAR NOT NULL,
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT (NOW())
);


ALTER TABLE "subcategories" ADD FOREIGN KEY ("category_id") REFERENCES "categories" ("id") ON DELETE CASCADE;

