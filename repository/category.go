package repository

import (
	"money-treez/model"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	// Create a new category
	CreateExpenseCategory(expenseCategory model.ExpenseCategory) (model.ExpenseCategory, error)
	GetExpenseCategories() ([]model.ExpenseCategory, error)
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *categoryRepository {
	return &categoryRepository{db}

}


func (r *categoryRepository) CreateExpenseCategory(expenseCategory model.ExpenseCategory) (model.ExpenseCategory, error) {
	err := r.db.Create(&expenseCategory).Error
	if err != nil {
		return expenseCategory, err
	}
	return expenseCategory, nil
}

func (r *categoryRepository) GetExpenseCategories() ([]model.ExpenseCategory, error) {
	var expenseCategories []model.ExpenseCategory
	if err := r.db.Find(&expenseCategories).Error; err != nil {
		return expenseCategories, err
	}
	return expenseCategories, nil

}