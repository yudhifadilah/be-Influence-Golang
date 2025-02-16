package models

import "time"

type Influencer struct {
	ID                 uint      `gorm:"primaryKey" json:"id"`
	Email              string    `gorm:"unique;not null" json:"email"`
	Password           string    `json:"password"`
	FullName           string    `json:"full_name"`
	BirthDate          string    `json:"birth_date"`
	Gender             string    `json:"gender"`
	InfluencerCategory string    `json:"influencer_category"`
	PhoneNumber        string    `json:"phone_number"`
	KTPNumber          string    `json:"ktp_number"`
	NPWPNumber         string    `json:"npwp_number"`
	InstagramLink      string    `json:"instagram_link"`
	FollowersCount     int       `json:"followers_count"`
	ProfilePicture     string    `json:"profile_picture"`
	BankAccount        string    `json:"bank_account"`
	AccountNumber      string    `json:"account_number"`
	Province           string    `json:"province"`
	City               string    `json:"city"`
	RegistrationDate   time.Time `gorm:"autoCreateTime" json:"registration_date"`
	Role               string    `json:"role" gorm:"default:'influencer'"`

	// Relasi
	Services []Service `gorm:"foreignKey:InfluencerID;constraint:OnDelete:CASCADE;"` // Relasi dengan Service

}
