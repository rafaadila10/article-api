package models

import "time"

type Post struct {
	ID          uint      `gorm:"primaryKey;autoIncrement" json:"Id"`
	Title       string    `gorm:"type:varchar(200)" json:"Title"`
	Content     string    `gorm:"type:text" json:"Content"`
	Category    string    `gorm:"type:varchar(100)" json:"Category"`
	CreatedDate time.Time `gorm:"autoCreateTime" json:"Created_date"`
	UpdatedDate time.Time `gorm:"autoUpdateTime" json:"Updated_date"`
	Status      string    `gorm:"type:varchar(100)" json:"Status"` // Publish | Draft | Thrash
}
