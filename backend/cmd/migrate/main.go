package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/pressly/goose"
	"gshop/pkg/config"

	_ "github.com/lib/pq"
)

const dialect = "postgres"

var (
	flags = flag.NewFlagSet("migrate", flag.ExitOnError)
	dir   = flags.String("dir", "./migrations", "directory with migration files")
)

func main() {
	args := os.Args[1:]
	flags.Usage = usage

	if len(args) == 0 || args[0] == "-h" || args[0] == "--help" {
		flags.Usage()

		return
	}

	command := args[0]
	switch command {
	case "create":
		if err := goose.Run("create", nil, *dir, args[1:]...); err != nil {
			log.Fatalf("migrate run: %v", err)
		}
		return
	case "fix":
		if err := goose.Run("fix", nil, *dir); err != nil {
			log.Fatalf("migrate run: %v", err)
		}
		return
	}

	cfg := config.Config

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s", cfg.POSTGRES.Host, cfg.POSTGRES.User, cfg.POSTGRES.Pass, cfg.POSTGRES.Database, cfg.POSTGRES.Port, cfg.POSTGRES.Sslmode)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(fmt.Errorf("sql.Open %w", err))
	}

	if err := db.Ping(); err != nil {
		panic(fmt.Errorf("db.Ping %w", err))
	}

	if err := goose.SetDialect(dialect); err != nil {
		log.Fatal(err)
	}

	if err := goose.Run(command, db, *dir, args[1:]...); err != nil {
		log.Fatalf("migrate run: %v", err)
	}
}

func usage() {
	fmt.Println(usagePrefix)
	flags.PrintDefaults()
	fmt.Println(usageCommands)
}

var (
	usagePrefix = `Usage: migrate [OPTIONS] COMMAND
Examples:
    migrate status
Options:
`

	usageCommands = `
Commands:
    up                   Migrate the DB to the most recent version available
    up-by-one            Migrate the DB up by 1
    up-to VERSION        Migrate the DB to a specific VERSION
    down                 Roll back the version by 1
    down-to VERSION      Roll back to a specific VERSION
    redo                 Re-run the latest migration
    reset                Roll back all migrations
    status               Dump the migration status for the current DB
    version              Print the current version of the database
    create NAME [sql|go] Creates new migration file with the current timestamp
    fix                   Apply sequential ordering to migrations
`
)
