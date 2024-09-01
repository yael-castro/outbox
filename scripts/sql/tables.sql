DROP TABLE IF EXISTS purchases;
DROP TABLE IF EXISTS purchases_outbox;

CREATE TABLE purchases (
    id SERIAL PRIMARY KEY,
    order_id INTEGER UNIQUE,
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now(),
    deleted_at TIMESTAMP DEFAULT now()
);

CREATE TABLE purchases_outbox (
    id SERIAL PRIMARY KEY,
    purchase_id INTEGER NOT NULL REFERENCES purchases(id),
    order_id INTEGER NOT NULL,
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now(),
    deleted_at TIMESTAMP DEFAULT now()
);