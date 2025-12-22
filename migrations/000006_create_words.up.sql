CREATE TABLE IF NOT EXISTS words (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
    word TEXT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW()
)