-- +goose Up
-- +goose StatementBegin

CREATE TABLE IF NOT EXISTS "users" (
    id BIGSERIAL PRIMARY KEY,
    nama TEXT NOT NULL,
    role VARCHAR(20) DEFAULT 'student'
);

CREATE TABLE IF NOT EXISTS "google_account" (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL REFERENCES "users"(id) ON DELETE CASCADE,
    email TEXT UNIQUE NOT NULL,
    access_token TEXT NOT NULL,
    refresh_token TEXT NOT NULL,
    token_expiry TIMESTAMP WITH TIME ZONE NOT NULL
);

CREATE TABLE IF NOT EXISTS "subject" (
    id BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    code VARCHAR(10) UNIQUE NOT NULL,
    sks INTEGER NOT NULL DEFAULT 3
);

CREATE TABLE IF NOT EXISTS "paralel" (
    id BIGSERIAL PRIMARY KEY,
    nama VARCHAR(5) NOT NULL, 
    subject_id BIGINT NOT NULL REFERENCES "subject"(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS "enrollment" (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL REFERENCES "users"(id) ON DELETE CASCADE,
    paralel_id BIGINT NOT NULL REFERENCES "paralel"(id) ON DELETE CASCADE,
    enrolled_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(user_id, paralel_id) 
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "enrollment";
DROP TABLE IF EXISTS "paralel";
DROP TABLE IF EXISTS "subject";
DROP TABLE IF EXISTS "google_account";
DROP TABLE IF EXISTS "users";
-- +goose StatementEnd
