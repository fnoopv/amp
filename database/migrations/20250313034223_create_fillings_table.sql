-- migrate:up
CREATE TABLE fillings (
    "id" char(36) NOT NULL,
    "name" varchar(255) NOT NULL,
    "organization_id" char(36),
    "kind_primary" varchar(50) NOT NULL,
    "kind_secondary" varchar(50) DEFAULT NULL,
    "serial_number" varchar(50) NOT NULL,
    "completed_at" timestamp NOT NULL,
    "grade_level" varchar(10) DEFAULT NULL,
    "description" varchar(255) DEFAULT NULL,
    "created_at" timestamp NOT NULL,
    "updated_at" timestamp NOT NULL,
    "deleted_at" timestamp,
    PRIMARY KEY (id),

    CONSTRAINT fk_fillings_organization
    FOREIGN KEY (organization_id)
    REFERENCES organizations (id)
    ON DELETE SET NULL
);

-- migrate:down
DROP TABLE IF EXISTS fillings;
