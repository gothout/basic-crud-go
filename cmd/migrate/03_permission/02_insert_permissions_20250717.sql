INSERT INTO admin_permission (code, description) VALUES
    ('system', 'Systematic user'),
    ('create-enterprise', 'Allows the user to create an enterprise'),
    ('read-enterprise', 'Allows the user to read enterprises'),
    ('delete-enterprise', 'Allows the user to delete an enterprise'),
    ('update-enterprise', 'Allows the user to update enterprise data'),
    ('read-user', 'Allows the user to read user data'),
    ('read-user-enterprise', 'Allows the user to read users within the enterprise'),
    ('update-enterprise-user', 'Allows the user to update users within the enterprise'),
    ('delete-enterprise-user', 'Allows the user to delete users within the enterprise'),
    ('create-user-enterprise', 'Allows the user to create a user within the enterprise'),
    ('create-user-admin', 'Allows the user to create an admin user'),
    ('delete-enterprise-admin', 'Allows the user to delete an admin from an enterprise'),
    ('permission-apply-admin', 'Allows the user to assign admin permissions'),
    ('permission-apply-enterprise', 'Allows the user to assign enterprise permissions'),
    ('read-permission', 'Allows the user to read permissions')
ON CONFLICT (code) DO NOTHING;

-- Aplica a permissão "system" ao usuário "system@admin.local"
INSERT INTO user_permission (user_id, permission_id)
SELECT
    u.id,
    p.id
FROM "user" u
         JOIN admin_permission p ON p.code = 'system'
WHERE u.email = 'system@admin.local'
ON CONFLICT (user_id, permission_id) DO NOTHING;
