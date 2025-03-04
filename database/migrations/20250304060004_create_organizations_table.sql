-- migrate:up
CREATE TABLE "organizations" (
    "id" char(36) NOT NULL,
    "parent_id" char(36) DEFAULT null,
    "name" varchar(255) NOT NULL,
    "kind" varchar(255) NOT NULL,
    "order" integer DEFAULT null,
    "created_at" timestamp NOT NULL,
    "updated_at" timestamp NOT NULL,
    PRIMARY KEY ("id")
);
CREATE INDEX "idx_parent_id" ON "organizations" ("parent_id");
CREATE UNIQUE INDEX "uni_name" ON "organizations" ("name");
CREATE INDEX "idx_kind" ON "organizations" ("kind");
CREATE INDEX "idx_created_at" ON "organizations" ("created_at");

-- migrate:down
DROP TABLE IF EXISTS organizations;
