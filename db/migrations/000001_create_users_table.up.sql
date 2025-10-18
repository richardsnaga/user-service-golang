CREATE TABLE users (
  "id" SERIAL PRIMARY KEY,
  "name" VARCHAR(100),
  "email" VARCHAR(100) UNIQUE,
  "password_hash" VARCHAR(225),
  "created_at" timestamptz DEFAULT 'now()',
  "updated_at" timestamptz
);
