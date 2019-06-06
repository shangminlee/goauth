package database

import (
    "fmt"
    "github.com/jinzhu/gorm"
    "github.com/shangminlee/goauth/config"
    "github.com/shangminlee/goauth/log"
    "time"

    // Driver
    _ "github.com/jinzhu/gorm/dialects/mysql"
)

func init()  {
    gorm.NowFunc = func() time.Time {
        return time.Now().UTC()
    }
}

// NewDatabase returns a gorm.DB struct, gorm.DB.DB() returns a database handle
func NewDatabase(cfg *config.Config) (*gorm.DB, error) {

    // Mysql
    if cfg.Database.Type == "mysql" {

        args := fmt.Sprintf(
            "%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
            cfg.Database.User,
            cfg.Database.Password,
            cfg.Database.Host,
            cfg.Database.Port,
            cfg.Database.DatabaseName,
        )

        log.INFO.Printf("args : %s \n", args)
        db, err := gorm.Open(cfg.Database.Type, args)
        if err != nil {
            return  db, err
        }

        // set max idle connections
        db.DB().SetMaxIdleConns(cfg.Database.MaxIdleConns)
        // set max connections
        db.DB().SetMaxOpenConns(cfg.Database.MaxOpenConns)
        // Database logging  是否开启调试模式
        db.LogMode(cfg.IsDevelopment)

        return db, nil
    }

    // Database type not supported
    return nil, fmt.Errorf("Database type %s not supported", cfg.Database.Type)
}