CREATE TABLE IF NOT EXISTS flashcard_sessions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
    user_id UUID REFERENCES users (id) ON DELETE CASCADE,
    started_at TIMESTAMPTZ DEFAULT NOW(),
    ended_at TIMESTAMPTZ,
    current_meaning_id UUID REFERENCES word_meanings (id) ON DELETE CASCADE,
    current_meaning_translation_id UUID REFERENCES word_translations (id) ON DELETE CASCADE,
    current_meaning_choices_ids UUID [],
    cards_total INT NOT NULL,
    cards_completed INT NOT NULL DEFAULT 0
)