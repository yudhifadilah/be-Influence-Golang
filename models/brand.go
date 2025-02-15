package models

import (
	"gorm.io/gorm"
)

type Brand struct {
	gorm.Model
	Email     string `json:"email" gorm:"unique;not null"`
	Password  string `json:"password" gorm:"not null"`
	BrandName string `json:"brand_name" gorm:"not null"`
	PicName   string `json:"pic_name" gorm:"not null"`
	PicPhone  string `json:"pic_phone" gorm:"not null"`
	Province  string `json:"province" gorm:"not null"`
	City      string `json:"city" gorm:"not null"`
	BrandLogo string `json:"brand_logo"`
	Role      string `json:"role" gorm:"default:'brand'"`
}
