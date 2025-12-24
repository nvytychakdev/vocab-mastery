CREATE TABLE IF NOT EXISTS parts_of_speech (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
    code TEXT NOT NULL UNIQUE
)