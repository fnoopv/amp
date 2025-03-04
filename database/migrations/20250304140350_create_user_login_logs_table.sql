-- migrate:up
CREATE TABLE "user_login_logs" (
    "id" char(36) NOT NULL,
    "user_id" char(36) NOT null,
    "login_at" timestamp NOT NULL,
    "is_success" boolean NOT NULL,
    "ip_address" varchar(255) NOT NULL,
    "user_agent" varchar(255) DEFAULT NULL,
    "failure_reason" varchar(255) DEFAULT NULL,
    "created_at" timestamp NOT NULL,
    "updated_at" timestamp NOT NULL,
    PRIMARY KEY ("id")
);
CREATE INDEX "idx_user_id" ON "user_login_logs" ("user_id");
CREATE INDEX "idx_login_at" ON "user_login_logs" ("login_at");
CREATE INDEX "idx_is_success" ON "user_login_logs" ("is_success")
CREATE INDEX "idx_ip_address" ON "user_login_logs" ("ip_address")
CREATE INDEX "idx_created_at" ON "user_login_logs" ("created_at")

-- migrate:down
DROP TABLE IF EXISTS user_login_log;
