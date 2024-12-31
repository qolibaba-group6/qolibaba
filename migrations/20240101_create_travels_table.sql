
CREATE TABLE IF NOT EXISTS travels (
    id UUID PRIMARY KEY,
    type VARCHAR(50) NOT NULL,
    origin VARCHAR(100) NOT NULL,
    destination VARCHAR(100) NOT NULL,
    price NUMERIC(10, 2) NOT NULL,
    release_date TIMESTAMP NOT NULL
);

INSERT INTO travels (id, type, origin, destination, price, release_date)
VALUES 
    ('123e4567-e89b-12d3-a456-426614174000', 'هوایی', 'تهران', 'مشهد', 500000.00, '2024-01-01 10:00:00'),
    ('123e4567-e89b-12d3-a456-426614174001', 'دریایی', 'بندرعباس', 'کیش', 700000.00, '2024-01-05 14:00:00'),
    ('123e4567-e89b-12d3-a456-426614174002', 'زمینی', 'اصفهان', 'شیراز', 300000.00, '2024-01-10 08:00:00');
