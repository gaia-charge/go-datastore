-- Routing events
CREATE TYPE routing_event_type AS ENUM (
    'FORWARD', 
    'RECEIVE',
    'SEND'
);

CREATE TABLE IF NOT EXISTS routing_events (
    id                 BIGSERIAL PRIMARY KEY,
    node_id            BIGINT NOT NULL,
    event_type         routing_event_type NOT NULL,
    currency           TEXT NOT NULL,
    currency_rate      BIGINT NOT NULL,
    currency_rate_msat BIGINT NOT NULL,
    incoming_chan_id   BIGINT NOT NULL,
    incoming_htlc_id   BIGINT NOT NULL,
    incoming_fiat      FLOAT NOT NULL,
    incoming_msat      BIGINT NOT NULL,
    outgoing_chan_id   BIGINT NOT NULL,
    outgoing_htlc_id   BIGINT NOT NULL,
    outgoing_fiat      FLOAT NOT NULL,
    outgoing_msat      BIGINT NOT NULL,
    fee_fiat           FLOAT NOT NULL,
    fee_msat           BIGINT NOT NULL,
    last_updated       TIMESTAMP NOT NULL
);

ALTER TABLE routing_events 
    ADD CONSTRAINT fk_routing_events_node_id 
    FOREIGN KEY (node_id) 
    REFERENCES nodes(id) 
    ON DELETE CASCADE;
