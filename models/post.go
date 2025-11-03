package models

import "time"

type Post struct {
	ID          uint      `gorm:"primaryKey;autoIncrement" json:"Id"`
	Title       string    `gorm:"type:varchar(200)" json:"Title" validate:"required,min=20"`
	Content     string    `gorm:"type:text" json:"Content" validate:"required,min=200"`
	Category    string    `gorm:"type:varchar(100)" json:"Category" validate:"required,min=3"`
	CreatedDate time.Time `gorm:"autoCreateTime" json:"Created_date"`
	UpdatedDate time.Time `gorm:"autoUpdateTime" json:"Updated_date"`
	Status      string    `gorm:"type:varchar(100)" json:"Status" validate:"required,oneof=publish draft thrash"` // Publish | Draft | Thrash
}
