package config

import (
    "github.com/jinzhu/configor"
    "github.com/shangminlee/goauth/log"
)

// DefaultConfig ... 默认配置
// Let's start with some sensible defaults 为配置项设置一些合理的默认值
var DefaultConfig = &Config{
    Database: DatabaseConfig{
        Type:         "mysql",
        Host:         "127.0.0.1",
        Port:         3306,
        User:         "root",
        Password:     "Abc_123456",
        DatabaseName: "goauth",
        MaxIdleConns: 5,
        MaxOpenConns: 5,
    },
    Oauth: OauthConfig{
        AccessTokenLifetime:  3600,    // 1 hour
        RefreshTokenLifetime: 1209600, // 14 days
        AuthCodeLifeTime:     3600,    // 1 hour
    },
    Session: sessionConfig{
        Secret:   "test_secret",
        Path:     "/",
        MaxAge:   86400 * 7, // 7 days
        HTTPOnly: true,
    },

    ServerPort:    8080,
    IsDevelopment: true,
}

func NewDefaultConfig() *Config {
    return DefaultConfig
}

func NewConfig(configFile string) *Config {
    if configFile != "" {
        config := &Config{}
        err := configor.Load(config, configFile)
        if err != nil {
            log.FATAL.Println(err)
        }
        return config
    }
    return NewDefaultConfig()
}
