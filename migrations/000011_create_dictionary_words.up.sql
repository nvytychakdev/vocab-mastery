CREATE TABLE IF NOT EXISTS dictionary_words (
    dictionary_id UUID REFERENCES dictionaries (id) ON DELETE CASCADE,
    word_id UUID REFERENCES words (id) ON DELETE CASCADE,
    added_at TIMESTAMPTZ DEFAULT NOW()
)