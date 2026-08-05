CREATE TABLE commission_settings (

    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    commission_rate NUMERIC(5,2) NOT NULL,

    minimum_withdrawal NUMERIC(12,2) DEFAULT 2000,

    withdrawal_fee NUMERIC(12,2) DEFAULT 0,

    active BOOLEAN DEFAULT TRUE,

    created_at TIMESTAMP DEFAULT NOW()
);