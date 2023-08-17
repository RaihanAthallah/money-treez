package repository

import (
	"money-treez/model"

	"gorm.io/gorm"
)

type IncomeRepository interface {
	// Create a new income
	CreateIncome(income model.Income) (model.Income, error)
	// Get a income by id
	GetIncome(id int) (model.Income, error)
	// Get all incomes
	GetIncomes() ([]model.Income, error)
}

type incomeRepository struct {
	db *gorm.DB
}

func NewIncomeRepository(db *gorm.DB) *incomeRepository {
	return &incomeRepository{db}
}

func (r *incomeRepository) CreateIncome(income model.Income) (model.Income, error) {
	err := r.db.Create(&income).Error
	if err != nil {
		return income, err
	}
	return income, nil
}

func (r *incomeRepository) GetIncome(id int) (model.Income, error) {
	var income model.Income
	if err := r.db.First(&income, id).Error; err != nil {
		return income, err
	}
	return income, nil
}

func (r *incomeRepository) GetIncomes() ([]model.Income, error) {
	var incomes []model.Income
	if err := r.db.Find(&incomes).Error; err != nil {
		return incomes, err
	}
	return incomes, nil

}
