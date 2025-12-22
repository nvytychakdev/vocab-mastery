CREATE TABLE IF NOT EXISTS dictionaries (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
    owner_id UUID REFERENCES users (id) ON DELETE CASCADE,
    title TEXT NOT NULL,
    level TEXT,
    is_default boolean,
    created_at TIMESTAMPTZ DEFAULT NOW()
)