INSERT INTO admin_permission (code, description) VALUES
    ('create-enterprise', 'Allows the user to create an enterprise'),
    ('delete-enterprise', 'Allows the user to delete an enterprise'),
    ('update-enterprise', 'Allows the user to update enterprise data'),
    ('read-user', 'Allows the user to read user data'),
    ('read-user-enterprise', 'Allows the user to read users within the enterprise'),
    ('update-enterprise-user', 'Allows the user to update users within the enterprise'),
    ('create-user-enterprise', 'Allows the user to create a user within the enterprise'),
    ('create-user-admin', 'Allows the user to create an admin user'),
    ('delete-enterprise-admin', 'Allows the user to delete an admin from an enterprise'),
    ('permission-apply-admin', 'Allows the user to assign admin permissions'),
    ('permission-apply-enterprise', 'Allows the user to assign enterprise permissions'),
    ('read-permission', 'Allows the user to read permissions')
ON CONFLICT (code) DO NOTHING;
