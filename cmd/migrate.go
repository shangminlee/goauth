package cmd

import (
    "github.com/shangminlee/goauth/log"
    "github.com/shangminlee/goauth/models"
    "github.com/shangminlee/goauth/util/migrations"
)

// Migrate runs database migrations
func Migrate(configFile string) error  {
    log.INFO.Print("迁移数据库")
    _, db, err := initConfigDB(configFile)
    if err != nil {
        return err
    }
    defer func() {
        err := db.Close()
        if err != nil {
            log.FATAL.Print(err)
        }
    }()

    // Bootstrap migrations
    if err := migrations.Bootstrap(db); err != nil {
        return err
    }

    // Run migrations for the oauth service
    if err := models.MigrateAll(db); err != nil {
       return err
    }

    return nil
}