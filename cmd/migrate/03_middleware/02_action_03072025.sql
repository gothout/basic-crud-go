CREATE TABLE IF NOT EXISTS action (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL UNIQUE
);
INSERT INTO action (name) VALUES
('create'),
('edit'),
('view'),
('delete')
ON CONFLICT DO NOTHING;
