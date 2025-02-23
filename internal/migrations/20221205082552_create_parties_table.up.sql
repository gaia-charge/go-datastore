
-- Parties
CREATE TABLE IF NOT EXISTS parties (
    id                          BIGSERIAL PRIMARY KEY,
    credential_id               BIGINT NOT NULL,
    country_code                TEXT NOT NULL,
    party_id                    TEXT NOT NULL,
    is_intermediate_cdr_capable BOOLEAN NOT NULL,
    publish_location            BOOLEAN NOT NULL,
    publish_null_tariff         BOOLEAN NOT NULL
);

ALTER TABLE parties 
    ADD CONSTRAINT uq_parties_country_code_party_id 
    UNIQUE (credential_id, country_code, party_id);

ALTER TABLE parties 
    ADD CONSTRAINT fk_parties_credential_id
    FOREIGN KEY (credential_id) 
    REFERENCES credentials(id) 
    ON DELETE CASCADE;

-- Locations
ALTER TABLE locations
    ADD COLUMN is_intermediate_cdr_capable BOOLEAN NOT NULL DEFAULT true;

-- Tariffs
ALTER TABLE tariffs
    DROP IF EXISTS is_intermediate_cdr_capable;
