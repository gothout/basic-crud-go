CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE TABLE IF NOT EXISTS "user" (
  id VARCHAR(36) PRIMARY KEY DEFAULT gen_random_uuid()::text,
  enterprise_id INTEGER REFERENCES enterprise(id) ON DELETE CASCADE,
  number VARCHAR(30),
  first_name VARCHAR(255) NOT NULL,
  last_name VARCHAR(255) NOT NULL,
  email VARCHAR(255) NOT NULL UNIQUE,
  password VARCHAR(255) NOT NULL,
  created_at TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Insere um usuário administrador apenas se não existir o mesmo email
INSERT INTO "user" (id, enterprise_id, number, first_name, last_name, email, password, created_at, updated_at)
SELECT
    gen_random_uuid()::text,
    e.id,
    '554791447700',
    'System',
    'Admin',
    'system@admin.local',
    '$argon2id$v=19$m=65536,t=3,p=2$xh0OyTJkyVQzSQkFNPlOSQ$FEIU7Emdq/v83NLnkW2Da/ZuPP+ClZT2JULpU1/6t+8', --admin123
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
FROM enterprise e
WHERE e.name = 'System Enterprise'
ON CONFLICT (email) DO NOTHING;

CREATE TABLE IF NOT EXISTS admin_api_token (
    id SERIAL PRIMARY KEY,
    user_id VARCHAR(36) NOT NULL REFERENCES "user"(id) ON DELETE CASCADE,
    token TEXT NOT NULL UNIQUE,
    created_at TIMESTAMP NOT NULL,
    end_date TIMESTAMP NOT NULL
);


CREATE TABLE IF NOT EXISTS admin_user_token (
    id SERIAL PRIMARY KEY,
    user_id VARCHAR(36) NOT NULL REFERENCES "user"(id) ON DELETE CASCADE,
    token TEXT NOT NULL UNIQUE,
    created_at TIMESTAMP NOT NULL
);