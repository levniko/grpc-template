package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id           int            `json:"id" gorm:"primaryKey;autoIncrement;notNull"`
	FirstName    string         `json:"first_name" gorm:"type:varchar(50)"`
	LastName     string         `json:"last_name" gorm:"type:varchar(50)"`
	PhoneNumber  string         `json:"phone_number" gorm:"type:varchar(12)"`
	Email        string         `json:"email" gorm:"type:varchar(75)"`
	Password     string         `json:"password" gorm:"type:varchar(100)"`
	PasswordHash string         `json:"password_hash" gorm:"type:varchar(100)"`
	IsAdmin      bool           `json:"is_admin" gorm:"type:bool"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at"`
}

func (user *User) TableName() string {
	return "users"
}

func (user *User) BeforeCreate(tx *gorm.DB) error {
	user.CreatedAt = time.Now()
	return nil
}

func (user *User) BeforeUpdate(tx *gorm.DB) error {
	user.UpdatedAt = time.Now()
	return nil
}
