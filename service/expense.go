package service

import (
	"money-treez/model"
	"money-treez/repository"
)

type ExpenseService interface {
	// Create a new expense
	CreateExpense(expense model.Expense) (model.Expense, error)

	// Get a expense by id
	GetExpense(id int) (model.Expense, error)
	// Get all expenses
	GetExpenses() ([]model.Expense, error)
}

type expenceService struct {
	expenseRepo repository.ExpenseRepository
}

func NewExpenseService(expenseRepo repository.ExpenseRepository) *expenceService {
	return &expenceService{expenseRepo}
}

func (s *expenceService) CreateExpense(expense model.Expense) (model.Expense, error) {
	expense, err := s.expenseRepo.CreateExpense(expense)
	if err != nil {
		return expense, err
	}
	return expense, nil

}

func (s *expenceService) GetExpense(id int) (model.Expense, error) {
	expense, err := s.expenseRepo.GetExpense(id)
	if err != nil {
		return expense, err
	}
	return expense, nil
}

func (s *expenceService) GetExpenses() ([]model.Expense, error) {
	expenses, err := s.expenseRepo.GetExpenses()
	if err != nil {
		return expenses, err
	}
	return expenses, nil

}
