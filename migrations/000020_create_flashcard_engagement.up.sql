CREATE TABLE IF NOT EXISTS flashcard_engagement (
    user_id UUID PRIMARY KEY REFERENCES users (id) ON DELETE CASCADE,
    last_active_at TIMESTAMPTZ NOT NULL,
    last_session_date DATE,
    reminder_stage TEXT NOT NULL,
    missed_days_count INT NOT NULL DEFAULT 0,
    next_reminder_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
)