package db

import (
	"database/sql"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
)

func DbConnection() (*sql.DB, error) {
	db, err := sql.Open("postgres", "postgres://postgres:postgres@localhost:5432/students?sslmode=disable")
	if err != nil {
		return nil, err
	}
	return db, nil

}

// MIGRATIONS UP
func RunMigrations() error {
	m, err := migrate.New(
		"file://db/migrations", "postgres://postgres:postgres@localhost:5432/students?sslmode=disable")

	if err != nil {
		return err
	}
	fmt.Println("Migrations up successfully")
	m.Up()

	return nil
}
