package model

// User is a representation of a user
type User struct {
	ID        uint   `json:"id" gorm:"primary_key"`
	FirstName string `json:"first_name" gorm:"not null"`
	LastName  string `json:"last_name" gorm:"not null"`
	Email     string `json:"email" gorm:"not null"`
	Password  string `json:"password" gorm:"not null"`
}
