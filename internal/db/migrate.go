package db

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
)

func (d *Database) MigrateDB() error {
	fmt.Println("Migrating our Database")
	driver, err := postgres.WithInstance(d.Client.DB, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("Could Not create Postgres Driver: %w", err)
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file:///migrations",
		"postgres", driver)
	if err != nil {
		fmt.Println(err)
		return err
	}
	if err := m.Up(); err != nil {
		return fmt.Errorf("Could not run up Migrations: %w", err)
	}
	fmt.Println("Successfully Migrated the Database")
	return nil
}
