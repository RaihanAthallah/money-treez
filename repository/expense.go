package repository

import (
	"money-treez/model"

	"gorm.io/gorm"
)

type ExpenseRepository interface {
	// Create a new expense
	CreateExpense(expense model.Expense) (model.Expense, error)
	
	// Get a expense by id
	GetExpense(id int) (model.Expense, error)
	// Get all expenses
	GetExpenses() ([]model.Expense, error)
	
}

type expenseRepository struct {
	db *gorm.DB
}

func NewExpenseRepository(db *gorm.DB) *expenseRepository {
	return &expenseRepository{db}
}

func (r *expenseRepository) CreateExpense(expense model.Expense) (model.Expense, error) {
	err := r.db.Create(&expense).Error
	if err != nil {
		return expense, err
	}
	return expense, nil
}

func (r *expenseRepository) CreateExpenseCategory(expenseCategory model.ExpenseCategory) (model.ExpenseCategory, error) {
	err := r.db.Create(&expenseCategory).Error
	if err != nil {
		return expenseCategory, err
	}
	return expenseCategory, nil
}

func (r *expenseRepository) GetExpense(id int) (model.Expense, error) {
	var expense model.Expense
	if err := r.db.First(&expense, id).Error; err != nil {
		return expense, err
	}
	return expense, nil
}

func (r *expenseRepository) GetExpenses() ([]model.Expense, error) {
	var expenses []model.Expense
	if err := r.db.Find(&expenses).Error; err != nil {
		return expenses, err
	}
	return expenses, nil

}

func (r *expenseRepository) GetExpenseCategories() ([]model.ExpenseCategory, error) {
	var expenseCategories []model.ExpenseCategory
	if err := r.db.Find(&expenseCategories).Error; err != nil {
		return expenseCategories, err
	}
	return expenseCategories, nil

}
