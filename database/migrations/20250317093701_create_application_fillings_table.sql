-- migrate:up
CREATE TABLE application_fillings (
    "application_id" char(36) NOT NULL,
    "filling_id" char(36) NOT NULL,
    PRIMARY KEY (application_id, filling_id),

    CONSTRAINT fk_application_fillings_application
    FOREIGN KEY (application_id)
    REFERENCES applications (id)
    ON DELETE CASCADE,

    CONSTRAINT fk_application_fillings_filling
    FOREIGN KEY (filling_id)
    REFERENCES fillings (id)
    ON DELETE CASCADE
);

-- migrate:down
DROP TABLE IF EXISTS application_fillings;
