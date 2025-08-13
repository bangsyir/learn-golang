package database

import (
	"embed"
	"fmt"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
	"golang-blueprint-v1/config"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

// run all migration file
func RunMigrations(cfg *config.DBConfig) error {
	db, err := goose.OpenDBWithDriver("pgx", cfg.DSN)
	if err != nil {
		return err
	}

	defer db.Close()
	// setup embedded FS
	goose.SetBaseFS(embedMigrations)

	// run migrations
	goose.SetTableName("schema_migrations") // custom migration table
	if err := goose.SetDialect("postgres"); err != nil {
		return err
	}

	if err := goose.Up(db, "migrations"); err != nil {
		return err
	}

	fmt.Println("Migrations applied successfully")
	return nil
}
