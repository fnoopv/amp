-- migrate:up
CREATE TABLE application_networks (
    "application_id" char(36) NOT NULL,
    "network_id" char(36) NOT NULL,
    PRIMARY KEY (application_id, network_id),

    CONSTRAINT fk_application_networks_application
    FOREIGN KEY (application_id)
    REFERENCES applications (id)
    ON DELETE CASCADE,

    CONSTRAINT fk_application_networks_network
    FOREIGN KEY (network_id)
    REFERENCES networks (id)
    ON DELETE CASCADE
);

-- migrate:down
DROP TABLE IF EXISTS application_networks;
