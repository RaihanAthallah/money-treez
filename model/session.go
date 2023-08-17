package model

import "time"

type Session struct {
	ID     int       `gorm:"primaryKey" json:"id"`
	Token  string    `json:"token" gorm:"not null"`
	Email  string    `json:"email" gorm:"not null"`
	Expiry time.Time `json:"expiry" gorm:"not null"`
}
