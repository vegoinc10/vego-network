CREATE TABLE stores (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    owner_id UUID NOT NULL REFERENCES users(id),

    name VARCHAR(255) NOT NULL,

    slug VARCHAR(255) UNIQUE,

    description TEXT,

    logo_url TEXT,

    banner_url TEXT,

    email VARCHAR(255),

    phone VARCHAR(30),

    state VARCHAR(100),

    lga VARCHAR(100),

    address TEXT,

    latitude NUMERIC(10,8),

    longitude NUMERIC(11,8),

    verified BOOLEAN DEFAULT FALSE,

    status VARCHAR(30) DEFAULT 'active',

    created_at TIMESTAMP DEFAULT NOW(),

    updated_at TIMESTAMP DEFAULT NOW()
);