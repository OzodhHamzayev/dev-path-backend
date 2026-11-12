CREATE TABLE orders (
    id SERIAL PRIMARY KEY,

    user_id INTEGER NOT NULL,
    course_id INTEGER NOT NULL,

    amount BIGINT NOT NULL CHECK (amount >= 0),
    currency VARCHAR(3) NOT NULL DEFAULT 'UZS',

    status VARCHAR(20) NOT NULL DEFAULT 'pending'
        CHECK (status IN ('pending', 'expired', 'paid', 'cancelled')),

    expires_at TIMESTAMPTZ NOT NULL
        DEFAULT (NOW() + INTERVAL '30 minutes'),

    paid_at TIMESTAMPTZ,

    idempotency_key VARCHAR(100) NOT NULL,

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    UNIQUE (user_id, idempotency_key)
);

ALTER TABLE orders 
ADD CONSTRAINT fk_orders_user
FOREIGN KEY (user_id) REFERENCES users(id);

ALTER TABLE orders 
ADD CONSTRAINT fk_orders_course 
FOREIGN KEY (course_id) REFERENCES courses(id);

