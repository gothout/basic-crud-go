-- Habilitar a extensão (executar uma vez no banco)
CREATE EXTENSION IF NOT EXISTS "pgcrypto";
CREATE TABLE IF NOT EXISTS "user" (
    id VARCHAR(36) PRIMARY KEY,
    enterprise_id INTEGER REFERENCES enterprise(id) ON DELETE CASCADE,
    number VARCHAR(30),
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP
);
INSERT INTO "user" (id, enterprise_id, number, first_name, last_name, email, password, created_at, updated_at)
SELECT 
    gen_random_uuid(),
    e.id,
    '+00000000000',
    'System',
    'Admin',
    'system@admin.local',
    'admin123',
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
FROM enterprise e
WHERE e.name = 'System Enterprise'
ON CONFLICT DO NOTHING;
