CREATE TABLE "notes" (
  "id" bigserial PRIMARY KEY,
  "title" varchar NOT NULL,
  "note" varchar,
  "tags" varchar[],
  "user_id" bigint NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
);