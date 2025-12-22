CREATE TABLE IF NOT EXISTS word_translations (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
    meaning_id UUID REFERENCES word_meanings (id) ON DELETE CASCADE,
    language TEXT REFERENCES languages (code) NOT NULL,
    translation TEXT NOT NULL
)