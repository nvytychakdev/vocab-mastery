CREATE TABLE IF NOT EXISTS word_meanings (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
    word_id UUID REFERENCES words (id) ON DELETE CASCADE,
    part_of_speech_id UUID REFERENCES parts_of_speech (id),
    definition TEXT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW()
)