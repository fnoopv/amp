-- migrate:up
CREATE TABLE apis (
    "id" char(36) NOT NULL,
    "feature_id" char(36) NOT NULL,
    "method" char(7) NOT NULL,
    "path" varchar(255) NOT NULL,
    "description" varchar(255) DEFAULT NULL,
    "created_at" timestamp NOT NULL,
    "updated_at" timestamp NOT NULL,
    PRIMARY KEY (id),

    CONSTRAINT fk_api_features_feature
    FOREIGN KEY (feature_id)
    REFERENCES features (id)
    ON UPDATE CASCADE
);

-- migrate:down
DROP TABLE IF EXISTS apis;
