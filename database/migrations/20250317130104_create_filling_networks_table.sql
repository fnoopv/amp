-- migrate:up
CREATE TABLE filling_networks (
    "network_id" char(36) NOT NULL,
    "filling_id" char(36) NOT NULL,
    PRIMARY KEY (network_id, filling_id),

    CONSTRAINT fk_filling_networks_network
    FOREIGN KEY (network_id)
    REFERENCES networks (id)
    ON DELETE CASCADE,

    CONSTRAINT fk_filling_networks_filling
    FOREIGN KEY (filling_id)
    REFERENCES fillings (id)
    ON DELETE CASCADE
);

-- migrate:down
DROP TABLE IF EXISTS filling_networks;
