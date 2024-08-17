ALTER TABLE if exists "accounts" DROP CONSTRAINT if exists "owner_currency_key";
ALTER TABLE if exists "accounts" DROP CONSTRAINT if exists "accounts_owner_key";

DROP TABLE "users" CASCADE;
