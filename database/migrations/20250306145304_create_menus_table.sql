-- migrate:up
CREATE TABLE menus (
    "id" char(36) NOT NULL,
    "parent_id" char(36),
    "feature_id" char(36) NOT NULL,
    "method" char(7) NOT NULL,
    "icon" varchar(255) NOT NULL,
    "path" varchar(255) NOT NULL,
    "order" integer NOT NULL DEFAULT 0,
    "is_hidden" boolean NOT NULL DEFAULT FALSE,
    "description" varchar(255) DEFAULT NULL,
    "created_at" timestamp NOT NULL,
    "updated_at" timestamp NOT NULL,
    PRIMARY KEY (id),

    CONSTRAINT fk_feature_menus_feature
    FOREIGN KEY (feature_id)
    REFERENCES features (id)
    ON UPDATE CASCADE,

    CONSTRAINT fk_menus_menu
    FOREIGN KEY (parent_id)
    REFERENCES menus (id)
    ON UPDATE CASCADE
);

-- migrate:down
DROP TABLE IF EXISTS menus;
