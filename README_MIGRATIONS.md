## Database migrations (goose)

This project uses `goose` with a separate `migrations/` folder.

### 1) Install goose (CLI) [optional]
```bash
go install github.com/pressly/goose/v3/cmd/goose@latest
```

You can also run migrations via the Go program at `cmd/migrate`.

### 2) Environment variables
The migrator uses the same env vars as the app:

- `DB_HOST`
- `DB_PORT`
- `DB_USER`
- `DB_PASSWORD`
- `DB_NAME`

When using `docker-compose`, these are provided to the `app` service; for local runs create a `.env` file in repo root.

### 3) Run migrations locally (via Go program)
```bash
go run ./cmd/migrate            # applies all pending migrations (Up)
go run ./cmd/migrate status     # prints status
go run ./cmd/migrate down       # rolls back one migration
go run ./cmd/migrate reset      # rolls back all
go run ./cmd/migrate up-to 1    # migrate up to a specific version
```

You can pass a custom directory with `-dir` if needed (defaults to `./migrations`).

### 4) Run migrations with goose CLI (alternative)
```bash
goose -dir ./migrations postgres "host=$DB_HOST port=$DB_PORT user=$DB_USER password=$DB_PASSWORD dbname=$DB_NAME sslmode=disable" up
```

### 5) Creating new migrations
Use SQL migrations in `migrations/`:

- Up/Down markers are required:
```sql
-- +goose Up
-- migration statements

-- +goose Down
-- rollback statements
```

Name files like `00002_add_index_to_username.sql`.

### Current migrations
- `00001_create_employeedata.sql`: Creates `employeedata` with unique constraints and useful indexes.


