CREATE TABLE "diary_entries" (
  "id" bigserial PRIMARY KEY,
  "title" varchar NOT NULL,
  "text" varchar NOT NULL,
  "entry_date" timestamptz NOT NULL DEFAULT (now()),
  "user_id" bigserial NOT NULL REFERENCES users(id),
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
)