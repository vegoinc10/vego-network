CREATE TABLE withdrawals (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    wallet_id UUID NOT NULL
        REFERENCES wallets(id)
        ON DELETE CASCADE,

    amount NUMERIC(12,2) NOT NULL,

    bank_name VARCHAR(255) NOT NULL,

    account_name VARCHAR(255) NOT NULL,

    account_number VARCHAR(20) NOT NULL,

    status VARCHAR(30) NOT NULL DEFAULT 'pending
    approved,
    processing,
    completed,
    rejected,
    failed,
    cancelled',

    reference VARCHAR(255),

    failure_reason TEXT,

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),

    processed_at TIMESTAMP,

    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_withdrawals_wallet
ON withdrawals(wallet_id);

CREATE INDEX idx_withdrawals_status
ON withdrawals(status);

CREATE INDEX idx_withdrawals_created
ON withdrawals(created_at);