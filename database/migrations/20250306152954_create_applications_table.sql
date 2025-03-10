-- migrate:up
CREATE TABLE applications (
    "id" char(36) NOT NULL,
    "organization_id" char(36) DEFAULT NULL,
    "name" varchar(255) NOT NULL,
    "description" varchar(255) DEFAULT NULL,
    "launched_at" timestamp DEFAULT NULL,
    "created_at" timestamp NOT NULL,
    "updated_at" timestamp NOT NULL,
    "deleted_at" timestamp DEFAULT NULL,
    PRIMARY KEY (id),

    CONSTRAINT fk_application_orgsnization
    FOREIGN KEY (organization_id)
    REFERENCES organizations (id)
    ON UPDATE CASCADE
    ON DELETE SET NULL
);
CREATE INDEX applications_organization_id_idx ON applications (organization_id);

-- migrate:down
DROP TABLE IF EXISTS applications;
