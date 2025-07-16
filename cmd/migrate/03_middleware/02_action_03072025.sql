CREATE TABLE IF NOT EXISTS admin_action (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL UNIQUE
);
INSERT INTO admin_action (name) VALUES
('create'),
('edit'),
('view'),
('delete')
ON CONFLICT DO NOTHING;
