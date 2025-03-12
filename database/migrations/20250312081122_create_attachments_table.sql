-- migrate:up
CREATE TABLE attachments (
    "id" char(36) NOT NULL,
    "uploader_id" char(36),
    "name" varchar(255) NOT NULL,
    "mime" varchar(255) NOT NULL,
    "size" bigint NOT NULL,
    "storage_path" varchar(255) NOT NULL,
    "sha256_sum" char(64) NOT NULL,
    "business_kind" varchar(255) DEFAULT NULL,
    "business_id" char(36) DEFAULT NULL,
    "upload_at" timestamp NOT NULL,
    "created_at" timestamp NOT NULL,
    "updated_at" timestamp NOT NULL,
    PRIMARY KEY (id),

    CONSTRAINT fk_attachments_users_user
    FOREIGN KEY (uploader_id)
    REFERENCES users (id)
    ON DELETE SET NULL
);

-- migrate:down
DROP TABLE IF EXISTS attachments;
