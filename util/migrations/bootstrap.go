package migrations

import (
    "fmt"
    "github.com/jinzhu/gorm"
    "github.com/shangminlee/goauth/log"
)

// Bootstrap creates "migrations" table
// to keep tract of already run database migrations
func Bootstrap(db *gorm.DB) error{
    migrationName := "bootstrap_migrations"

    migration := new(Migration)

    // Using Error instead of RecordNotFound because we want to check
    // if the migrations table exists. This is different from later migrations
    // where we query the already create migrations table.
    // SELECT * FROM migration WHERE name = 'bootstrap_migrations' LIMIT 1;
    exists := nil == db.Where("name = ?", migrationName).First(migration)

    if exists {
        log.INFO.Printf("Skipping %s migration", migrationName)
        return nil
    }

    log.INFO.Printf("Skipping %s migration", migrationName)

    // Create migrations table
    if err := db.CreateTable(new(Migration)).Error; err != nil {
        return fmt.Errorf("error creating migrations table: %v", db.Error)
    }

    // Save a record to migrations table,
    // so we don't return this migration again
    if err := db.Create(migration).Error; err != nil {
        return fmt.Errorf("error saving record to migrations table: %v", db.Error)
    }

    return nil

}