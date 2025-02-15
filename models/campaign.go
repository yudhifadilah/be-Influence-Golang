package models

type Campaign struct {
	ID           uint   `gorm:"primaryKey" json:"id"`
	Name         string `gorm:"not null" json:"name" binding:"required"`
	Category     string `gorm:"not null" json:"category" binding:"required"`
	InfluencerID uint   `gorm:"not null" json:"influencer_id" binding:"required"`
	StartDate    string `gorm:"not null" json:"start_date" binding:"required"`
	EndDate      string `gorm:"not null" json:"end_date" binding:"required"`
	Status       string `json:"status"`
	PDFFile      string `json:"pdf_file"`
}
