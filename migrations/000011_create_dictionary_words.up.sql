CREATE TABLE IF NOT EXISTS dictionary_word (
    dictionary_id UUID REFERENCES dictionaries (id) ON DELETE CASCADE,
    word_id UUID REFERENCES words (id) ON DELETE CASCADE,
    added_at TIMESTAMPTZ DEFAULT NOW()
)