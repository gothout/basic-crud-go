CREATE TABLE IF NOT EXISTS admin_module (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL UNIQUE
);
INSERT INTO admin_module (name) VALUES
('admin'),
('enterprise')
ON CONFLICT DO NOTHING;
