CREATE TABLE "accounts" (
  "id" bigserial PRIMARY KEY,
  "onwer" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "currency" varchar NOT NULL,
  "balance" bigint NOT NULL
);

CREATE TABLE "entries" (
  "id" bigserial PRIMARY KEY,
  "account_id" bigserial NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "amount" bigint NOT NULL
);

CREATE TABLE "transfers" (
  "id" bigserial PRIMARY KEY,
  "from_account_id" bigserial NOT NULL,
  "to_account_id" bigserial NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "amount" bigint NOT NULL
);

CREATE INDEX ON "accounts" ("onwer");

CREATE INDEX ON "entries" ("account_id");

CREATE INDEX ON "transfers" ("from_account_id");

CREATE INDEX ON "transfers" ("to_account_id");

CREATE INDEX ON "transfers" ("from_account_id", "to_account_id");

ALTER TABLE "entries" ADD FOREIGN KEY ("account_id") REFERENCES "accounts" ("id");

ALTER TABLE "transfers" ADD FOREIGN KEY ("from_account_id") REFERENCES "accounts" ("id");

ALTER TABLE "transfers" ADD FOREIGN KEY ("to_account_id") REFERENCES "accounts" ("id");