CREATE TABLE IF NOT EXISTS user_words_progress (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
    user_id UUID REFERENCES users (id) ON DELETE CASCADE,
    meaning_id UUID REFERENCES word_meanings (id) ON DELETE CASCADE,
    status TEXT,
    difficulty INT,
    times_seen_recall INT,
    times_correct_recall INT,
    times_incorrect_recall INT,
    next_review_at_recall TIMESTAMPTZ,
    times_seen_recognition INT,
    times_correct_recognition INT,
    times_incorrect_recognition INT,
    next_review_at_recognition TIMESTAMPTZ,
    last_seen_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ DEFAULT NOW()
)