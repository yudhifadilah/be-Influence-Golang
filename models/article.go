package models

import "time"

// Article struct merepresentasikan tabel articles di database
type Article struct {
	ID      uint      `gorm:"primaryKey" json:"id"`
	Title   string    `gorm:"type:varchar(255);not null" json:"title"`
	Excerpt string    `gorm:"type:text;not null" json:"excerpt"`
	Content string    `gorm:"type:text;not null" json:"content"`
	Image   string    `gorm:"type:varchar(255);not null" json:"image"`
	RegDate time.Time `gorm:"autoCreateTime;autoUpdateTime" json:"reg_date"`
}
