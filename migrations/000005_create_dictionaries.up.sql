CREATE TABLE IF NOT EXISTS dictionaries (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
    user_id PREFERENCES users (id) ON DELETE CASCADE,
    name TEXT NOT NULL,
    description TEXT,
    created_at TIMESTAMPTZ DEFAULT NOW()
)