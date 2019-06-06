package models

import (
    "database/sql"
    "github.com/RichardKnop/uuid"
    "github.com/jinzhu/gorm"
    "github.com/spring2go/gravitee/util"
    "time"
)

// oauth 授权表
type OauthClient struct {
    MyGormModel
    ClientKey    string          `sql:"type:varchar(254);unique;not null"`
    ClientSecret string          `sql:"type:varchar(60);not null"`
    RedirectURI  sql.NullString  `sql:"type:varchar(200)"`
}

func (c *OauthClient) TableName() string {
    return "oauth_clients"
}

// 权限明细
type OauthScope struct {
    MyGormModel
    Scope       string         `sql:"type:varchar(200); unique; not null"`
    Description sql.NullString
    IsDefault   bool           `sql:"default:false"`
}

func (c *OauthScope) TableName() string {
    return "oauth_scopes"
}

// 角色信息
type OauthRole struct {
    TimestampModel
    ID   string `gorm:"primary_key" sql:"type:varchar(20)"`
    Name string `sql:"type:varchar(50); unique; not null"`
}

func (c *OauthRole) TableName() string {
    return "oauth_roles"
}

// 用户-角色信息
type OauthUser struct {
    MyGormModel
    RoleID   sql.NullString `sql:"type:varchar(20);index;not null"`
    Role     *OauthRole
    Username string         `sql:"type:varchar(254);unique;not null"`
    Password sql.NullString `sql:"type:varchar(60)"`
}

func (u *OauthUser) TableName() string {
    return "oauth_users"
}

// 刷新Token表
type OauthRefreshToken struct {
    MyGormModel
    ClientID  sql.NullString `sql:"index;not null"`
    UserID    sql.NullString `sql:"index"`
    Client    *OauthClient
    User      *OauthUser
    Token     string         `sql:"type:varchar(40);unique;not null"`
    ExpiresAt time.Time      `sql:"not null;DEFAULT:current_timestamp"`
    Scope     string         `sql:"type:varchar(200);not null"`
}

func (rt *OauthRefreshToken) TableName() string {
    return "oauth_refresh_tokens"
}

// Token 表
type OauthAccessToken struct {
    MyGormModel
    ClientID  sql.NullString `sql:"index;not null"`
    UserID    sql.NullString `sql:"index"`
    Client    *OauthClient
    User      *OauthUser
    Token     string         `sql:"type:varchar(40);unique;not null"`
    ExpiresAt time.Time      `sql:"not null;DEFAULT:current_timestamp"`
    Scope     string         `sql:"type:varchar(200);not null"`
}

func (at *OauthAccessToken) TableName() string {
    return "oauth_access_tokens"
}

// 授权码表
type OauthAuthorizationCode struct {
    MyGormModel
    ClientID    sql.NullString `sql:"index;not null"`
    UserID      sql.NullString `sql:"index;not null"`
    Client      *OauthClient
    User        *OauthUser
    Code        string         `sql:"type:varchar(40);unique;not null"`
    RedirectURI sql.NullString `sql:"type:varchar(200)"`
    ExpiresAt   time.Time      `sql:"not null;DEFAULT:current_timestamp"`
    Scope       string         `sql:"type:varchar(200);not null"`
}

func (ac *OauthAuthorizationCode) TableName() string {
    return "oauth_authorization_codes"
}

// NewOauthRefreshToken creates new OauthRefreshToken instance
func NewOauthRefreshToken(client *OauthClient, user *OauthUser, expiresIn int, scope string) *OauthRefreshToken {
    refreshToken := &OauthRefreshToken{
        MyGormModel: MyGormModel{
            ID:        uuid.New(),
            CreatedAt: time.Now().UTC(),
        },
        ClientID:  util.StringOrNull(string(client.ID)),
        Token:     uuid.New(),
        ExpiresAt: time.Now().UTC().Add(time.Duration(expiresIn) * time.Second),
        Scope:     scope,
    }
    if user != nil {
        refreshToken.UserID = util.StringOrNull(string(user.ID))
    }
    return refreshToken
}

// NewOauthAccessToken creates new OauthAccessToken instance
func NewOauthAccessToken(client *OauthClient, user *OauthUser, expiresIn int, scope string) *OauthAccessToken {
    accessToken := &OauthAccessToken{
        MyGormModel: MyGormModel{
            ID:        uuid.New(),
            CreatedAt: time.Now().UTC(),
        },
        ClientID:  util.StringOrNull(string(client.ID)),
        Token:     uuid.New(),
        ExpiresAt: time.Now().UTC().Add(time.Duration(expiresIn) * time.Second),
        Scope:     scope,
    }
    if user != nil {
        accessToken.UserID = util.StringOrNull(string(user.ID))
    }
    return accessToken
}

// NewOauthAuthorizationCode creates new OauthAuthorizationCode instance
func NewOauthAuthorizationCode(client *OauthClient, user *OauthUser, expiresIn int, redirectURI, scope string) *OauthAuthorizationCode {
    return &OauthAuthorizationCode{
        MyGormModel: MyGormModel{
            ID:        uuid.New(),
            CreatedAt: time.Now().UTC(),
        },
        ClientID:    util.StringOrNull(string(client.ID)),
        UserID:      util.StringOrNull(string(user.ID)),
        Code:        uuid.New(),
        ExpiresAt:   time.Now().UTC().Add(time.Duration(expiresIn) * time.Second),
        RedirectURI: util.StringOrNull(redirectURI),
        Scope:       scope,
    }
}

// OauthAuthorizationCodePreload sets up Gorm preloads for an auth code object
func OauthAuthorizationCodePreload(db *gorm.DB) *gorm.DB {
    return OauthAuthorizationCodePreloadWithPrefix(db, "")
}

// OauthAuthorizationCodePreloadWithPrefix sets up Gorm preloads for an auth code object,
// and prefixes with prefix for nested objects
func OauthAuthorizationCodePreloadWithPrefix(db *gorm.DB, prefix string) *gorm.DB {
    return db.
        Preload(prefix + "Client").Preload(prefix + "User")
}

// OauthAccessTokenPreload sets up Gorm preloads for an access token object
func OauthAccessTokenPreload(db *gorm.DB) *gorm.DB {
    return OauthAccessTokenPreloadWithPrefix(db, "")
}

// OauthAccessTokenPreloadWithPrefix sets up Gorm preloads for an access token object,
// and prefixes with prefix for nested objects
func OauthAccessTokenPreloadWithPrefix(db *gorm.DB, prefix string) *gorm.DB {
    return db.
        Preload(prefix + "Client").Preload(prefix + "User")
}

// OauthRefreshTokenPreload sets up Gorm preloads for a refresh token object
func OauthRefreshTokenPreload(db *gorm.DB) *gorm.DB {
    return OauthRefreshTokenPreloadWithPrefix(db, "")
}

// OauthRefreshTokenPreloadWithPrefix sets up Gorm preloads for a refresh token object,
// and prefixes with prefix for nested objects
func OauthRefreshTokenPreloadWithPrefix(db *gorm.DB, prefix string) *gorm.DB {
    return db.
        Preload(prefix + "Client").Preload(prefix + "User")
}