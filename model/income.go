package model

// Income is a representation of an income
type Income struct {
	ID          uint   `json:"id" gorm:"primary_key"`
	Description string `json:"description" gorm:"not null"`
	Amount      uint   `json:"amount" gorm:"not null"`
	OwnerID     uint   `json:"owner_id" gorm:"not null"`
	Date        string `json:"date" gorm:"not null"`
	CreatedAt   string `json:"created_at" gorm:"not null"`
	UpdatedAt   string `json:"updated_at" gorm:"not null"`
}
