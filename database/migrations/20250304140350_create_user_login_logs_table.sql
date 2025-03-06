-- migrate:up
CREATE TABLE "user_login_logs" (
    "id" char(36) NOT NULL,
    "user_id" char(36) NOT NULL,
    "login_at" timestamp NOT NULL,
    "is_success" boolean NOT NULL,
    "ip_address" varchar(255) NOT NULL,
    "user_agent" varchar(255) DEFAULT NULL,
    "failure_reason" varchar(255) DEFAULT NULL,
    "created_at" timestamp NOT NULL,
    "updated_at" timestamp NOT NULL,
    PRIMARY KEY ("id"),

    CONSTRAINT fk_user_user_login_log
    FOREIGN KEY (user_id)
    REFERENCES users (id)
    ON UPDATE CASCADE
);
CREATE INDEX "user_login_logs_user_id_idx" ON "user_login_logs" ("user_id");
CREATE INDEX "user_login_logs_login_at_idx" ON "user_login_logs" ("login_at");
CREATE INDEX "user_login_logs_is_success_idx"
ON "user_login_logs" ("is_success");
CREATE INDEX "user_login_logs_ip_address_idx"
ON "user_login_logs" ("ip_address");
CREATE INDEX "user_login_logs_created_at_idx"
ON "user_login_logs" ("created_at");

-- migrate:down
DROP TABLE IF EXISTS user_login_logs;
