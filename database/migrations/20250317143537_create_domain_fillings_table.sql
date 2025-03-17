-- migrate:up
CREATE TABLE domain_fillings (
    "domain_id" char(36) NOT NULL,
    "filling_id" char(36) NOT NULL,
    PRIMARY KEY (domain_id, filling_id),

    CONSTRAINT fk_domain_fillings_domain
    FOREIGN KEY (domain_id)
    REFERENCES domains (id)
    ON DELETE CASCADE,

    CONSTRAINT fk_domain_fillings_filling
    FOREIGN KEY (filling_id)
    REFERENCES fillings (id)
    ON DELETE CASCADE
);

-- migrate:down
DROP TABLE IF EXISTS domain_fillings;
