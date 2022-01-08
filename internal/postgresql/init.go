package postgresql

import (
	"database/sql"
	"fmt"

	"github.com/spf13/viper"
)

func InitDb() (*sql.DB, error) {
	db, err := sql.Open("postgres", viper.GetString("DATABASE_URL"))

	if err != nil {
		return nil, fmt.Errorf("sql.Open %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("db.Ping %w", err)
	}

	return db, nil
}
