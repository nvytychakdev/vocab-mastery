CREATE TABLE IF NOT EXISTS sessions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
    user_id UUID REFERENCES users (id) ON DELETE CASCADE,
    jti UUID UNIQUE NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    expires_at TIMESTAMPTZ NOT NULL,
    refreshed_at TIMESTAMPTZ
)