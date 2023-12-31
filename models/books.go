package models

import "gorm.io/gorm"

type Books struct {
	ID        string  `json:"id" gorm:"primaryKey"`
	Author    *string `json:"author"`
	Title     *string `json:"title"`
	Publisher *string `json:"publisher"`
}

func MigrateBooks(db *gorm.DB) error {
	err := db.AutoMigrate(&Books{})
	return err
}
