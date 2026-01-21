CREATE TABLE IF NOT EXISTS flashcard_attempts (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
    session_id UUID REFERENCES flashcard_sessions (id) ON DELETE CASCADE,
    meaning_id UUID REFERENCES word_meanings (id) ON DELETE CASCADE,
    direction TEXT NOT NULL,
    prompt_language TEXT NOT NULL,
    answer_language TEXT NOT NULL,
    is_correct BOOLEAN,
    response_time_ms INT,
    created_at TIMESTAMPTZ DEFAULT NOW()
)