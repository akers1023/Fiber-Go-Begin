package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID           string    `json:"id" gorm:"primaryKey"`
	Full_name    string    `json:"full_name"`
	Email        *string   `json:"email" validate:"email,required"`
	Password     *string   `json:"password" validate:"required,min=6"`
	Role         *string   `json:"role" validate:"required"`
	Token        *string   `json:"token"`
	RefreshToken *string   `json:"refresh_token"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func MigrateUser(db *gorm.DB) error {
	err := db.AutoMigrate(&User{})
	return err
}
