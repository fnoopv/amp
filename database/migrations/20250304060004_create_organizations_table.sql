-- migrate:up
CREATE TABLE "organizations" (
    "id" char(36) NOT NULL,
    "parent_id" char(36),
    "name" varchar(255) NOT NULL,
    "kind" varchar(255) NOT NULL,
    "order" integer DEFAULT null,
    "description" text,
    "created_at" timestamp NOT NULL,
    "updated_at" timestamp NOT NULL,
    "deleted_at" timestamp DEFAULT null,
    PRIMARY KEY ("id"),

    CONSTRAINT fk_organization
    FOREIGN KEY (parent_id)
    REFERENCES organizations (id)
    ON DELETE SET NULL
);
CREATE INDEX "organizations_parent_id_idx" ON "organizations" ("parent_id");
CREATE INDEX "organizations_kind_idx" ON "organizations" ("kind");
CREATE INDEX "organizations_deleted_at_idx" ON "organizations" ("deleted_at");

-- migrate:down
DROP TABLE IF EXISTS organizations;
