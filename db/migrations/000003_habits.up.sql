CREATE TABLE "habits" (
  "id" bigserial PRIMARY KEY,
  "title" varchar NOT NULL,
  "frequency" varchar NOT NULL DEFAULT 'daily',
  "target_count" SMALLINT NOT NULL DEFAULT 1,
  "user_id" bigint NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "deleted_at" timestamptz
);

CREATE TABLE "habit_logs" (
  "id" bigserial PRIMARY KEY,
  "habit_id" bigint NOT NULL REFERENCES habits(id) ON DELETE CASCADE,
  "log_date" date NOT NULL DEFAULT CURRENT_DATE,
  "actual_count" SMALLINT NOT NULL DEFAULT 0,
  "note" varchar,
  "created_at" timestamptz NOT NULL DEFAULT (now()),

  CONSTRAINT unique_habit_log_per_day UNIQUE (habit_id, log_date)
)