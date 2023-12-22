package domain

import "time"

type User struct {
	ID      uint   `json:"id" gorm:"primaryKey"`
	Contact string `json:"contact" gorm:"unique;not null"`
}

type Card struct {
	ID              uint              `json:"card_id" gorm:"primaryKey"`
	Number          string            `json:"number"`
	UserID          uint              `json:"-"`
	User            User              `json:"user" gorm:"foreignKey:UserID"`
	DeliveryDetails []DeliveryDetails `json:"delivery_details" gorm:"foreignKey:CardID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}

type DeliveryDetails struct {
	ID               uint `gorm:"primaryKey"`
	CardID           uint `gorm:"index"`
	DeliveryAttempts int
	DeliveryDate     time.Time
	Comment          string
}
