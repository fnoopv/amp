-- migrate:up
CREATE TABLE business_attachments (
    "id" serial PRIMARY KEY,
    "business_type" varchar(50),
    "business_id" char(36) NOT NULL,
    "attachment_type" varchar(50) NOT NULL,
    "attachment_id" char(36) NOT NULL,

    CONSTRAINT fk_business_attachments_attachment
    FOREIGN KEY (attachment_id)
    REFERENCES attachments (id)
    ON DELETE SET NULL
);
CREATE INDEX business_attachmens_idx
ON business_attachments
(business_type, business_id, attachment_type);

-- migrate:down
DROP TABLE IF EXISTS business_attachments;
