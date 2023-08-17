package repository

import (
	"money-treez/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	// Create a new user
	CreateUser(user model.User) (model.User, error)
	// Get a user by id
	GetUser(id int) (model.User, error)
	// Get a user by email
	GetUserByEmail(email string) (model.User, error)
	// Get all users
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) CreateUser(user model.User) (model.User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *userRepository) GetUser(id int) (model.User, error) {
	var user model.User
	if err := r.db.First(&user, id).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (r *userRepository) GetUserByEmail(email string) (model.User, error) {
	var user model.User
	err := r.db.Where("email = ?", email).First(&user).Scan(&user).Error
	if err == gorm.ErrRecordNotFound {
		return user, nil
	}
	return user, nil
}
