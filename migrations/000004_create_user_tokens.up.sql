CREATE TABLE IF NOT EXISTS user_tokens (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
    user_id UUID REFERENCES users (id) ON DELETE CASCADE,
    token TEXT NOT NULL,
    type TEXT NOT NULL CHECK (
        type IN (
            'email_confirm',
            'password_reset'
        )
    ),
    created_at TIMESTAMPTZ DEFAULT NOW(),
    expires_at TIMESTAMPTZ NOT NULL,
    used_at TIMESTAMPTZ
)