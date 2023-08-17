package service

import (
	"money-treez/model"
	"money-treez/repository"
)

type CategoryService interface {
	// Create a new category
	CreateExpenseCategory(expenseCategory model.ExpenseCategory) (model.ExpenseCategory, error)
	GetExpenseCategories() ([]model.ExpenseCategory, error)
}

type categoryService struct {
	categoryRepo repository.CategoryRepository
}

func NewCategoryService(categoryRepo repository.CategoryRepository) *categoryService {
	return &categoryService{categoryRepo}
}

func (s *categoryService) CreateExpenseCategory(expenseCategory model.ExpenseCategory) (model.ExpenseCategory, error) {
	expenseCategory, err := s.categoryRepo.CreateExpenseCategory(expenseCategory)
	if err != nil {
		return expenseCategory, err
	}
	return expenseCategory, nil
}

func (s *categoryService) GetExpenseCategories() ([]model.ExpenseCategory, error) {
	expenseCategories, err := s.categoryRepo.GetExpenseCategories()
	if err != nil {
		return expenseCategories, err
	}
	return expenseCategories, nil

}
