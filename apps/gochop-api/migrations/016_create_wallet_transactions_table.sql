CREATE TABLE wallet_transactions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    wallet_id UUID NOT NULL REFERENCES wallets(id) ON DELETE CASCADE,

    type VARCHAR(20) NOT NULL,
    amount DECIMAL(12,2) NOT NULL,

    description TEXT,

    reference VARCHAR(255),

    created_at TIMESTAMP DEFAULT NOW()
);

CREATE INDEX idx_wallet_transactions_wallet
ON wallet_transactions(wallet_id);

CREATE INDEX idx_wallet_transactions_reference
ON wallet_transactions(reference);