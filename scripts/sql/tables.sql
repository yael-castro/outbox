DROP TABLE IF EXISTS outbox_messages;
DROP TABLE IF EXISTS purchases;

CREATE TABLE purchases (
    id SERIAL PRIMARY KEY,
    order_id INTEGER UNIQUE,
    -- Common fields
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now(),
    deleted_at TIMESTAMP DEFAULT now()
);

CREATE TABLE outbox_messages (
    id SERIAL PRIMARY KEY,
    topic VARCHAR NOT NULL,
    partition_key VARCHAR,
    header  BYTEA,
    content BYTEA NOT NULL,
    delivered_at TIMESTAMP DEFAULT NULL,
    -- Common fields
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now(),
    deleted_at TIMESTAMP DEFAULT now()
);