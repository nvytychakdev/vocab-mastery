CREATE TABLE IF NOT EXISTS word_examples (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
    meaning_id UUID REFERENCES word_meanings (id) ON DELETE CASCADE,
    text TEXT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW()
)