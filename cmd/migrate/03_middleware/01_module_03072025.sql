CREATE TABLE IF NOT EXISTS module (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL UNIQUE
);
INSERT INTO module (name) VALUES
('admin'),
('enterprise'),
('user')
ON CONFLICT DO NOTHING;
