CREATE TABLE order_items (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    order_id UUID NOT NULL REFERENCES orders(id) ON DELETE CASCADE,

    product_id UUID NOT NULL REFERENCES products(id) ON DELETE CASCADE,

    quantity INTEGER NOT NULL,

    unit_price DECIMAL(12,2) NOT NULL,

    total_price DECIMAL(12,2) NOT NULL,

    created_at TIMESTAMP DEFAULT NOW()
);