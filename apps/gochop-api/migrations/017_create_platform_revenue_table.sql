CREATE TABLE platform_revenue (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    order_id UUID NOT NULL
        REFERENCES orders(id)
        ON DELETE CASCADE,

    vendor_id UUID NOT NULL
        REFERENCES users(id)
        ON DELETE CASCADE,

    gross_amount NUMERIC(12,2) NOT NULL,

    commission_rate NUMERIC(5,2) NOT NULL,

    commission_amount NUMERIC(12,2) NOT NULL,

    vendor_amount NUMERIC(12,2) NOT NULL,

    created_at TIMESTAMP DEFAULT NOW()
);

CREATE INDEX idx_platform_revenue_order
ON platform_revenue(order_id);

CREATE INDEX idx_platform_revenue_vendor
ON platform_revenue(vendor_id);