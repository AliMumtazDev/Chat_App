package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	postgresdb "github.com/AliMumtazDev/Go_Chat_App/database/postgresdb"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
)

func main() {
	migrationsDir := flag.String("dir", "migrations", "directory with migration files")
	flag.Parse()

	// Load env from .env if present (local runs)
	if err := godotenv.Load(".env"); err != nil {
		log.Println("No .env file found; relying on shell/Docker env vars")
	}

	// Fast-fail with a clearer message if required DB envs are missing
	if os.Getenv("DB_HOST") == "" || os.Getenv("DB_PORT") == "" || os.Getenv("DB_USER") == "" || os.Getenv("DB_PASSWORD") == "" || os.Getenv("DB_NAME") == "" {
		log.Println("Missing one or more DB_* env vars (DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME)")
		log.Println("Set them in your shell or provide a .env file. Example:")
		log.Println("DB_HOST=localhost\nDB_PORT=5433\nDB_USER=postgres\nDB_PASSWORD=1233456\nDB_NAME=gotest")
	}

	db, err := postgresdb.PostgresConn()
	if err != nil {
		log.Fatalf("failed to connect to postgres: %v", err)
	}
	defer db.Close()

	if err := runGoose(db, *migrationsDir, flag.Args()); err != nil {
		log.Fatalf("migration failed: %v", err)
	}
}

func runGoose(db *sql.DB, dir string, args []string) error {
	// When using goose with the default OS filesystem, paths may include './'.
	// If a base FS is set, io/fs requires paths without '.' segments.
	// We don't set a base FS; still normalize dir for safety.
	if strings.HasPrefix(dir, "./") {
		dir = strings.TrimPrefix(dir, "./")
	}
	if err := goose.SetDialect("postgres"); err != nil {
		return fmt.Errorf("set dialect: %w", err)
	}

	if len(args) == 0 {
		// default to up
		return goose.Up(db, dir)
	}

	// support: up, down, status, redo, reset, up-by-one, down-to, up-to <version>
	cmd := args[0]
	cmdArgs := args[1:]
	return goose.Run(cmd, db, dir, cmdArgs...)
}
