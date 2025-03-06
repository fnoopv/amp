-- migrate:up
CREATE TABLE "users" (
    "id" char(36) NOT NULL,
    "email" varchar(255) DEFAULT null,
    "nick_name" varchar(255) NOT NULL,
    "username" varchar(255) NOT NULL,
    "password" char(60) NOT NULL,
    "status" varchar(255) NOT NULL,
    "mfa_key" varchar(255) DEFAULT null,
    "is_mfa_active" bool NOT NULL DEFAULT false,
    "password_updated_at" timestamp DEFAULT null,
    "created_at" timestamp NOT NULL,
    "updated_at" timestamp NOT NULL,
    PRIMARY KEY ("id")
);
CREATE INDEX "users_nick_name_idx" ON "users" ("nick_name");
CREATE UNIQUE INDEX "users_username_uni" ON "users" ("username");
CREATE INDEX "users_status_idx" ON "users" ("status");
CREATE INDEX "users_created_at_idx" ON "users" ("created_at");

-- migrate:down
DROP TABLE IF EXISTS users;
