DROP TABLE IF EXISTS purchases_outbox;
DROP TABLE IF EXISTS purchases;

CREATE TABLE purchases (
    id SERIAL PRIMARY KEY,
    order_id INTEGER UNIQUE,
    -- Common fields
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now(),
    deleted_at TIMESTAMP DEFAULT now()
);

CREATE TABLE purchases_outbox (
    id SERIAL PRIMARY KEY,
    order_id INTEGER NOT NULL,
    delivered_at TIMESTAMP DEFAULT NULL,
    purchase_id INTEGER NOT NULL REFERENCES purchases(id),
    -- Common fields
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now(),
    deleted_at TIMESTAMP DEFAULT now()
);