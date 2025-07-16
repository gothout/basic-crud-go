CREATE TABLE IF NOT EXISTS user_permission (
    id SERIAL PRIMARY KEY,
    user_id VARCHAR(36) NOT NULL REFERENCES "user"(id) ON DELETE CASCADE,
    permission_id INTEGER NOT NULL REFERENCES admin_permission(id) ON DELETE CASCADE,
    UNIQUE (user_id, permission_id)
);
--INSERT INTO user_permission (user_id, permission_id)
--SELECT u.id, p.id
--FROM "user" u
--JOIN permission p ON TRUE
--WHERE u.email = 'system@admin.local'
--ON CONFLICT DO NOTHING;
