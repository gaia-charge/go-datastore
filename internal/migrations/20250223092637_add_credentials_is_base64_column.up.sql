-- Credentials
ALTER TABLE credentials 
    ADD COLUMN is_base64 BOOLEAN NOT NULL DEFAULT false;
