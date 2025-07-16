CREATE TABLE IF NOT EXISTS admin_permission (
    id SERIAL PRIMARY KEY,
    module_id INTEGER NOT NULL REFERENCES admin_module(id) ON DELETE CASCADE,
    action_id INTEGER NOT NULL REFERENCES admin_action(id) ON DELETE CASCADE,
    UNIQUE (module_id, action_id)
);
INSERT INTO admin_permission (module_id, action_id)
SELECT m.id, a.id
FROM admin_module m, admin_action a
ON CONFLICT DO NOTHING;
