package models

import "gorm.io/gorm"

type Payment struct {
	gorm.Model
	TransactionID string  `json:"transaction_id" gorm:"uniqueIndex"`
	BuyerID       uint    `json:"buyer_id"`
	ServiceID     uint    `json:"service_id"`
	Amount        float64 `json:"amount"`
	Status        string  `json:"status" gorm:"type:enum('pending','success','failed');default:'pending'"`

	// Relasi
	Buyer   Brand   `gorm:"foreignKey:BuyerID;constraint:OnDelete:CASCADE;"`   // Buyer (Brand atau Customer)
	Service Service `gorm:"foreignKey:ServiceID;constraint:OnDelete:CASCADE;"` // Layanan yang dibeli
}
