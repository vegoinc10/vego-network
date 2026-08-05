-- 015_create_wallets_table.sql

CREATE TABLE wallets (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    user_id UUID NOT NULL UNIQUE
        REFERENCES users(id) ON DELETE CASCADE,

    available_balance NUMERIC(12,2) DEFAULT 0,
    pending_balance   NUMERIC(12,2) DEFAULT 0,
    total_earned      NUMERIC(12,2) DEFAULT 0,
    total_withdrawn   NUMERIC(12,2) DEFAULT 0,

    currency VARCHAR(3) DEFAULT 'NGN',

    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE INDEX idx_wallets_user_id ON wallets(user_id);