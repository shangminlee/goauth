package config

/**
    配置文件ORM
 */

// DatabaseConfig stores database connection options
type DatabaseConfig struct {
    Type         string `default:"mysql"`     // 数据类型
    Host         string `default:"localhost"` // 数据库地址
    Port         int    `default:"3306"`      // 数据库端口
    User         string `default:"root"`      // 数据库用户名
    Password     string `default:"Abc_123"`   // 数据库密码
    DatabaseName string `default:"goauth"`    // 数据库名字
    MaxIdleConns int    `default:"5"`         // 最大空闲连接
    MaxOpenConns int    `default:"5"`         // 最大连接数
}

// OauthConfig stores oauth service configuration options
type OauthConfig struct {
    AccessTokenLifetime  int  `default:"3600"`    // token 过期时间 default to 1 hour
    RefreshTokenLifetime int `default:"1209600"`  // 刷新Token 过期时间 default to 14 days
    AuthCodeLifeTime     int `default:"3600"`     // 授权码过期时间 default to 1 hour
}

// SessionConfig stores session configuration for the web app
type sessionConfig struct {
    Secret  string `default:"test_secret"`
    Path    string `default:"/"`
    // MaxAge=0 means no 'Max-Age' attribute specified
    // MaxAge<0 means delete cookie now, equivalently 'Max-Age=0'
    // MaxAge>0 means Max-Age attribute present and given in seconds
    MaxAge   int    `default:"604800"`
    // When you tag a cookie with the HttpOnly flag, it tells the browser that
    // this particular cookie should only be accessed by the server
    // Any attempt to access the cookie from client script is strictly forbidden
    HTTPOnly bool   `default:"True"`
}

// Config stores all configuration options
type Config struct {
    Database      DatabaseConfig        // 数据库配置
    Oauth         OauthConfig           // 授权码配置
    Session       sessionConfig         // 会话配置
    ServerPort    int `default:"8080"`  // 服务端口
    IsDevelopment bool `default:"True"` // 是否开发模式
}