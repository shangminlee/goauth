package cmd

import (
    "github.com/RichardKnop/go-fixtures"
    "github.com/shangminlee/goauth/log"
)

// LoadData loads fixtures
func LoadData(paths []string, configFile string) error {
    log.INFO.Print("Load Data ")
    cfg, db, err := initConfigDB(configFile)
    if err != nil {
        return err
    }
    defer func() {
        err := db.Close()
        log.INFO.Fatal(err)
    }()

    log.INFO.Printf("configuration : %v \n",cfg)
    for _, dataPath := range paths {
        log.INFO.Printf("DataPath : %s \n", dataPath)
    }

    return fixtures.LoadFiles(paths, db.DB(), cfg.Database.Type)
}