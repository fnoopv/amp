-- migrate:up
CREATE TABLE evaluations (
    "id" char(36) NOT NULL,
    "filling_id" char(36) NOT NULL,
    "serial_number" varchar(50) NOT NULL,
    "completed_at" timestamp NOT NULL,
    "created_at" timestamp NOT NULL,
    "updated_at" timestamp NOT NULL,
    "deleted_at" timestamp,
    PRIMARY KEY (id),

    CONSTRAINT fk_evaluations_fillings_evaluation
    FOREIGN KEY (filling_id)
    REFERENCES fillings (id)
);
CREATE INDEX "fillings_deleted_at_idx" ON "fillings" ("deleted_at");

-- migrate:down
DROP TABLE IF EXISTS evaluations;
