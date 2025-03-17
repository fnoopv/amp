-- migrate:up
CREATE TABLE domains (
    "id" char(36) NOT NULL,
    "domain" varchar(255) NOT NULL,
    "organization_id" char(36),
    "description" varchar(255) DEFAULT NULL,
    "created_at" timestamp NOT NULL,
    "updated_at" timestamp NOT NULL,
    "deleted_at" timestamp,
    PRIMARY KEY (id),

    CONSTRAINT fk_domains_organization
    FOREIGN KEY (organization_id)
    REFERENCES organizations (id)
    ON DELETE SET NULL
);
CREATE UNIQUE INDEX "domains_domain_deleted_at_uni"
ON "domains" ("domain", "deleted_at");

-- migrate:down
DROP TABLE IF EXISTS domains;
