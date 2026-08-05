CREATE TABLE IF NOT EXISTS orders (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    buyer_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,

    status VARCHAR(30) NOT NULL DEFAULT 'pending',

    total_amount NUMERIC(12,2) NOT NULL DEFAULT 0,

    payment_status VARCHAR(30) NOT NULL DEFAULT 'pending',

    delivery_status VARCHAR(30) NOT NULL DEFAULT 'pending',

    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS order_items (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    order_id UUID NOT NULL REFERENCES orders(id) ON DELETE CASCADE,

    product_id UUID NOT NULL REFERENCES products(id),

    quantity INT NOT NULL,

    price NUMERIC(12,2) NOT NULL,

    subtotal NUMERIC(12,2) NOT NULL
);