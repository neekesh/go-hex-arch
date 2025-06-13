package storage

import (
	"context"
	"embed"
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/Masterminds/squirrel"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/thapakazi/go-hex-arch/internal/adapter/config"
)

var migrationsFS embed.FS

type DB struct {
	*pgxpool.Pool
	QueryBuilder *squirrel.StatementBuilderType
	url          string
}

var Database *DB

func init() {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Environment.DBHost,
		config.Environment.DBPort,
		config.Environment.DBUser,
		config.Environment.DBPassword,
		config.Environment.DBName,
	)

	ctx := context.Background()
	db, err := pgxpool.New(ctx, dsn)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping(ctx)
	if err != nil {
		log.Fatal(err)
	}

	psql := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

	Database = &DB{
		Pool:         db,
		QueryBuilder: &psql,
		url:          dsn,
	}
}

// Migrate runs the database migration
func (db *DB) Migrate() error {
	driver, err := iofs.New(migrationsFS, "migrations")
	if err != nil {
		return err
	}

	migrations, err := migrate.NewWithSourceInstance("iofs", driver, db.url)
	if err != nil {
		return err
	}

	err = migrations.Up()
	if err != nil && err != migrate.ErrNoChange {
		return err
	}

	return nil
}

// ErrorCode returns the error code of the given error
func (Database *DB) ErrorCode(err error) string {
	pgErr := err.(*pgconn.PgError)
	return pgErr.Code
}

// ExecContext executes a SQL statement with the provided context
func (db *DB) ExecContext(ctx context.Context, sql string, args ...interface{}) (pgconn.CommandTag, error) {
	return db.Pool.Exec(ctx, sql, args...)
}

// ExecProcedureTx executes a stored procedure with given name and arguments inside a transaction.
func (db *DB) ExecProcedureTx(ctx context.Context, procName string, args []interface{}) error {
	tx, ok := ctx.Value(txKey{}).(pgx.Tx)
	if !ok {
		return fmt.Errorf("pgx.Tx not found in context")
	}

	if !regexp.MustCompile(`^[a-zA-Z_][a-zA-Z0-9_]*(\.[a-zA-Z_][a-zA-Z0-9_]*)?$`).MatchString(procName) {
		return fmt.Errorf("unsafe procedure name: %q", procName)
	}

	placeholders := make([]string, len(args))
	for i := range args {
		placeholders[i] = fmt.Sprintf("$%d", i+1)
	}

	stmt := fmt.Sprintf("CALL %s(%s)", procName, strings.Join(placeholders, ", "))

	_, err := tx.Exec(ctx, stmt, args...)
	if err != nil {
		return fmt.Errorf("procedure call failed [%s]: %w", stmt, err)
	}

	return nil
}

func (db *DB) Expr(query string, args ...interface{}) squirrel.Sqlizer {
	return squirrel.Expr(query, args...)
}
