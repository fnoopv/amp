-- migrate:up
CREATE TABLE networks (
    "id" char(36) NOT NULL,
    "name" varchar(255) NOT NULL,
    "organization_id" char(36),
    "filling_id" char(36),
    "description" varchar(255) DEFAULT NULL,
    "created_at" timestamp NOT NULL,
    "updated_at" timestamp NOT NULL,
    "deleted_at" timestamp,
    PRIMARY KEY (id),

    CONSTRAINT fk_networks_organization
    FOREIGN KEY (organization_id)
    REFERENCES organizations (id)
    ON DELETE SET NULL,

    CONSTRAINT fk_networks_filling
    FOREIGN KEY (filling_id)
    REFERENCES fillings (id)
    ON DELETE SET NULL
);
CREATE INDEX "networks_deleted_at_idx" ON "networks" ("deleted_at");

-- migrate:down
DROP TABLE IF EXISTS networks;
