-- +goose Up
CREATE TABLE IF NOT EXISTS employeedata (
    id SERIAL PRIMARY KEY,
    username VARCHAR(100) NOT NULL,
    email VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    CONSTRAINT unique_username UNIQUE (username),
    CONSTRAINT unique_email UNIQUE (email)
);

-- Indexes to speed up lookups
CREATE INDEX IF NOT EXISTS idx_employeedata_email ON employeedata (email);
CREATE INDEX IF NOT EXISTS idx_employeedata_username ON employeedata (username);

-- +goose Down
DROP TABLE IF EXISTS employeedata;


