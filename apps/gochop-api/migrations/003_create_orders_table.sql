CREATE TABLE orders (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    buyer_id UUID NOT NULL REFERENCES users(id),

    total_amount NUMERIC(12,2) NOT NULL,

    status VARCHAR(30) DEFAULT 'pending',

    payment_status VARCHAR(30) DEFAULT 'unpaid',

    delivery_status VARCHAR(30) DEFAULT 'pending',

    created_at TIMESTAMP DEFAULT NOW(),

    updated_at TIMESTAMP DEFAULT NOW()
);