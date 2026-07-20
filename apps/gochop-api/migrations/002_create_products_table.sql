CREATE TABLE products (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    seller_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,

    name VARCHAR(255) NOT NULL,

    description TEXT,

    category VARCHAR(100),

    price NUMERIC(12,2) NOT NULL,

    quantity INTEGER NOT NULL DEFAULT 0,

    image_url TEXT,

    market_name VARCHAR(255),

    location VARCHAR(255),

    status VARCHAR(30) DEFAULT 'available',

    created_at TIMESTAMP DEFAULT NOW(),

    updated_at TIMESTAMP DEFAULT NOW()
);