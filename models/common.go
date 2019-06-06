package models

import (
    "time"
)

type MyGormModel struct {
    ID        string `gorm:"primary_key"`
    CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt *time.Time
}

type TimestampModel struct {
    CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt *time.Time
}

type EmailTokenModel struct {
    MyGormModel
    Reference   string `sql:"type:varchar(40); unique; not null"`
    EmailSent   bool   `sql:"index; not null"`
    EmailSentAt *time.Time
    ExpiresAt   time.Time
}