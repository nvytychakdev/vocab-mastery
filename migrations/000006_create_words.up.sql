CREATE TABLE IF NOT EXISTS words (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
    dictionary_id UUID REFERENCES dictionaries (id) ON DELETE CASCADE,
    word TEXT NOT NULL,
    language TEXT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW()
)