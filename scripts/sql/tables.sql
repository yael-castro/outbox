DROP TABLE IF EXISTS purchases;
CREATE TABLE purchases (
    id SERIAL PRIMARY KEY,
    order_id INTEGER UNIQUE,
    -- Common fields
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now(),
    deleted_at TIMESTAMP DEFAULT NULL
);

DROP TABLE IF EXISTS outbox_messages;
CREATE TABLE outbox_messages (
    id SERIAL PRIMARY KEY,
    topic VARCHAR NOT NULL,
    partition_key BYTEA,
    headers BYTEA,
    value BYTEA NOT NULL,
    delivered_at TIMESTAMP DEFAULT NULL,
    -- Common fields
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now(),
    deleted_at TIMESTAMP DEFAULT NULL
);