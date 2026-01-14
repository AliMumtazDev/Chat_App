package postgresdb

import (
    "database/sql"
    "fmt"

    "github.com/pressly/goose/v3"
)

// Migrate applies all pending goose migrations in the ./migrations folder.
// It is safe to call multiple times; migrations are idempotent.
func Migrate(db *sql.DB) error {
    if err := goose.SetDialect("postgres"); err != nil {
        return fmt.Errorf("set dialect: %w", err)
    }
    dir := "migrations"
    if err := goose.Up(db, dir); err != nil {
        return fmt.Errorf("goose up: %w", err)
    }
    return nil
}

