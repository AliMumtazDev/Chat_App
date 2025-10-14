package postgresdb

import (
    "context"
    "database/sql"
    "errors"
    "fmt"
    "os"
    "strconv"
    "strings"

    "github.com/lib/pq"
)

// PostgresConn connects to the target database, creating it if it doesn't exist.
func PostgresConn() (*sql.DB, error) {
    host := os.Getenv("DB_HOST")
    portStr := os.Getenv("DB_PORT")
    user := os.Getenv("DB_USER")
    password := os.Getenv("DB_PASSWORD")
    dbname := os.Getenv("DB_NAME")

    port, err := strconv.Atoi(portStr)
    if err != nil {
        return nil, fmt.Errorf("invalid DB_PORT: %w", err)
    }

    // Try target DB first.
    targetDSN := fmt.Sprintf(
        "host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
        host, port, user, password, dbname,
    )
    db, err := sql.Open("postgres", targetDSN)
    if err != nil {
        return nil, fmt.Errorf("open target db: %w", err)
    }
    if pingErr := db.Ping(); pingErr == nil {
        fmt.Println("Connected to PostgreSQL with database/sql!")
        return db, nil
    } else {
        // If database is missing, create it and retry.
        var pqe *pq.Error
        if errors.As(pingErr, &pqe) && pqe.Code == "3D000" { // invalid_catalog_name
            if err := ensureDatabase(host, port, user, password, dbname); err != nil {
                db.Close()
                return nil, fmt.Errorf("ensure database: %w", err)
            }
            if err := db.Close(); err != nil {
                // ignore close error
            }
            // Reconnect to target DB
            db, err = sql.Open("postgres", targetDSN)
            if err != nil {
                return nil, fmt.Errorf("reopen target db: %w", err)
            }
            if err := db.Ping(); err != nil {
                db.Close()
                return nil, fmt.Errorf("ping after create: %w", err)
            }
            fmt.Println("Connected to PostgreSQL with database/sql (created DB)!")
            return db, nil
        }
        db.Close()
        return nil, fmt.Errorf("database ping failed: %v", pingErr)
    }
}

func ensureDatabase(host string, port int, user, password, dbname string) error {
    adminDSN := fmt.Sprintf(
        "host=%s port=%d user=%s password=%s dbname=postgres sslmode=disable",
        host, port, user, password,
    )
    adminDB, err := sql.Open("postgres", adminDSN)
    if err != nil {
        return fmt.Errorf("open admin db: %w", err)
    }
    defer adminDB.Close()

    // Check if DB exists
    var exists bool
    if err := adminDB.QueryRow("SELECT EXISTS (SELECT 1 FROM pg_database WHERE datname = $1)", dbname).Scan(&exists); err != nil {
        return fmt.Errorf("check database exists: %w", err)
    }
    if exists {
        return nil
    }
    // Create database
    if _, err := adminDB.ExecContext(context.Background(), "CREATE DATABASE "+quoteIdent(dbname)); err != nil {
        return fmt.Errorf("create database: %w", err)
    }
    return nil
}

func quoteIdent(s string) string {
    // minimal identifier quoting to avoid injection in DB name
    return `"` + strings.ReplaceAll(s, `"`, `""`) + `"`
}
