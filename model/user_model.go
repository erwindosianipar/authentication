package model

import (
	"database/sql"
	"time"
)

type GormCustom struct {
	ID        uint         `json:"id" gorm:"primary_key"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
	DeletedAt sql.NullTime `json:"deleted_at" sql:"index"`
}

type User struct {
	GormCustom
	Username string `json:"username" gorm:"unique"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
