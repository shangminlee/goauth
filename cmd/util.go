package cmd

import (
    "github.com/jinzhu/gorm"
    "github.com/shangminlee/goauth/config"
    "github.com/shangminlee/goauth/database"
    "github.com/shangminlee/goauth/log"
)

// initConfigDB loads the configuration and connects to the database
func initConfigDB(configFile string) (*config.Config, *gorm.DB, error) {
    log.INFO.Print("initConfigDB")

    // config
    cfg := config.NewConfig(configFile)

    // database
    db, err := database.NewDatabase(cfg)
    if err != nil {
        return nil, nil, err
    }

    return cfg, db, nil
}