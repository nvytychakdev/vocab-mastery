CREATE TABLE IF NOT EXISTS word_meanings (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
    word_id UUID REFERENCES words (id) ON DELETE CASCADE,
    definition TEXT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW()
)