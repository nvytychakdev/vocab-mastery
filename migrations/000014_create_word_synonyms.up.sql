CREATE TABLE IF NOT EXISTS word_synonyms (
    meaning_id UUID REFERENCES word_meanings (id) ON DELETE CASCADE,
    synonym_word_id UUID REFERENCES words (id) ON DELETE CASCADE
)