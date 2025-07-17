CREATE TABLE IF NOT EXISTS admin_permission (
    id SERIAL PRIMARY KEY,
    code VARCHAR(100) NOT NULL UNIQUE, -- Ex: 'create-enterprise'
    description TEXT NOT NULL          -- Ex: 'Allows the user to create an enterprise'
);
CREATE TABLE IF NOT EXISTS user_permission (
    id SERIAL PRIMARY KEY,
    user_id VARCHAR(36) NOT NULL REFERENCES "user"(id) ON DELETE CASCADE,
    permission_id INTEGER NOT NULL REFERENCES admin_permission(id) ON DELETE CASCADE,
    UNIQUE (user_id, permission_id)
);
