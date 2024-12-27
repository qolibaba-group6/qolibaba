
CREATE TABLE IF NOT EXISTS access_logs (
    id SERIAL PRIMARY KEY,
    user_id VARCHAR(50),
    resource VARCHAR(50),
    action VARCHAR(50),
    status VARCHAR(20), -- 'SUCCESS' or 'FAILURE'
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS user_activities (
    id SERIAL PRIMARY KEY,
    user_id VARCHAR(50),
    activity VARCHAR(255),
    description TEXT,
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS transactions (
    id SERIAL PRIMARY KEY,
    user_id VARCHAR(50),
    transaction_type VARCHAR(50), -- 'CHARGE', 'REFUND', etc.
    amount NUMERIC(10, 2),
    description TEXT,
    created_at TIMESTAMP DEFAULT NOW()
);
