package services

import (
	"database/sql"
	"fmt"
	"os"
)

func loadSQL(db *sql.DB, path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("repository error loading schema: %w", err)
	}
	_, err = db.Exec(string(data))
	return err
}

func loadSchemaAndSeed(db *sql.DB) error {
	if err := loadSQL(db, "../../db/schema.sql"); err != nil {
		return err
	}
	if err := loadSQL(db, "../../db/seed.sql"); err != nil {
		return err
	}
	return nil
}
