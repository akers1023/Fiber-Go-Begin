-- +migrate Up
CREATE TABLE "users" (
  "id" text PRIMARY KEY,
  "full_name" text,
  "email" text UNIQUE,
  "password" text,
  "role" text,
  "token" text UNIQUE,
  "refresh_token" text UNIQUE,
  "created_at" TIMESTAMPTZ NOT NULL,
  "updated_at" TIMESTAMPTZ NOT NULL
);

--   "created_at" TIMESTAMPTZ NOT NULL,
--   "updated_at" TIMESTAMPTZ NOT NULL

CREATE TABLE "books" (
  "id" text PRIMARY KEY,
  "author" text,
  "title" text,
  "publisher" text
);

ALTER TABLE "books" ADD FOREIGN KEY ("author") REFERENCES "users" ("id");

-- +migrate Down
DROP TABLE books;
DROP TABLE users;
