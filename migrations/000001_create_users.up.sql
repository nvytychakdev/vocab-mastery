CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
    email TEXT UNIQUE NOT NULL,
    name TEXT NOT NULL,
    is_email_confirmed BOOLEAN NOT NULL DEFAULT FALSE,
    password_hash TEXT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW()
)