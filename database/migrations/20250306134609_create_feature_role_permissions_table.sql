-- migrate:up
CREATE TABLE feature_role_permissions (
    "role_id" char(36) NOT NULL,
    "feature_id" char(36) NOT NULL,
    "read" boolean NOT NULL DEFAULT FALSE,
    "write" boolean NOT NULL DEFAULT FALSE,
    PRIMARY KEY (role_id, feature_id),

    CONSTRAINT fk_feature_roles_role
    FOREIGN KEY (role_id)
    REFERENCES roles (id)
    ON DELETE CASCADE,

    CONSTRAINT fk_feature_roles_feature
    FOREIGN KEY (feature_id)
    REFERENCES features (id)
    ON UPDATE CASCADE
);

-- migrate:down
DROP TABLE IF EXISTS featute_role_permissions;
