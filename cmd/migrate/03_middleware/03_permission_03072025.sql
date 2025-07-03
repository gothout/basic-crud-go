CREATE TABLE IF NOT EXISTS permission (
    id SERIAL PRIMARY KEY,
    module_id INTEGER NOT NULL REFERENCES module(id) ON DELETE CASCADE,
    action_id INTEGER NOT NULL REFERENCES action(id) ON DELETE CASCADE,
    UNIQUE (module_id, action_id)
);
INSERT INTO permission (module_id, action_id)
SELECT m.id, a.id
FROM module m, action a
ON CONFLICT DO NOTHING;
