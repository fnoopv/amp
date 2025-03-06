-- migrate:up
CREATE TABLE role_users (
    "user_id" char(36) NOT NULL,
    "role_id" char(36) NOT NULL,
    PRIMARY KEY (user_id, role_id),

    CONSTRAINT fk_role_users_user
    FOREIGN KEY (user_id)
    REFERENCES users (id)
    ON DELETE CASCADE,

    CONSTRAINT fk_role_users_role
    FOREIGN KEY (role_id)
    REFERENCES roles (id)
    ON DELETE CASCADE
);

-- migrate:down
DROP TABLE IF EXISTS role_users;
