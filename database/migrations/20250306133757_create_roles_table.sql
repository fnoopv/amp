-- migrate:up
CREATE TABLE roles (
    "id" char(36) NOT NULL,
    "name" varchar(255) NOT NULL,
    "description" varchar(255) DEFAULT NULL,
    "is_builtin" boolean NOT NULL DEFAULT FALSE,
    "created_at" timestamp NOT NULL,
    "updated_at" timestamp NOT NULL,
    "deleted_at" timestamp DEFAULT NULL,
    PRIMARY KEY (id)
);
CREATE INDEX "roles_deleted_at_idx" ON "roles" ("deleted_at");

-- migrate:down
DROP TABLE IF EXISTS roles;
