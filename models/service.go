package models

import "gorm.io/gorm"

type Service struct {
	gorm.Model
	InfluencerID uint    `json:"influencer_id"`
	ServiceName  string  `json:"service_name"`
	PricePerPost float64 `json:"price_per_post"`
	Description  string  `json:"description"`
	Duration     int     `json:"duration"`
}
