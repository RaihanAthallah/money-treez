package model

// Expense is a representation of an expense
type Expense struct {
	ID          uint   `json:"id" gorm:"primary_key"`
	Description string `json:"description" gorm:"not null"`
	Amount      uint   `json:"amount" gorm:"not null"`
	OwnerID     uint   `json:"owner_id" gorm:"not null"`
	Date        string `json:"date" gorm:"not null"`
	CategoryID  uint   `json:"category_id"`
	CreatedAt   string `json:"created_at" `
	UpdatedAt   string `json:"updated_at"`
}

type ExpenseCategory struct {
	ID          uint   `json:"id" gorm:"primary_key"`
	Description string `json:"description" gorm:"not null"`
	OwnerID     uint   `json:"owner_id" gorm:"not null"`
}
