package service

import (
	"money-treez/model"
	"money-treez/repository"
)

type IncomeService interface {
	// Create a new income
	CreateIncome(income model.Income) (model.Income, error)
	// Get a income by id
	GetIncome(id int) (model.Income, error)
	// Get all incomes
	GetIncomes() ([]model.Income, error)
}

type incomeService struct {
	incomeRepo repository.IncomeRepository
}

func NewIncomeService(incomeRepo repository.IncomeRepository) *incomeService {
	return &incomeService{incomeRepo}
}

func (s *incomeService) CreateIncome(income model.Income) (model.Income, error) {
	income, err := s.incomeRepo.CreateIncome(income)
	if err != nil {
		return income, err
	}
	return income, nil

}

func (s *incomeService) GetIncome(id int) (model.Income, error) {
	income, err := s.incomeRepo.GetIncome(id)
	if err != nil {
		return income, err
	}
	return income, nil
}

func (s *incomeService) GetIncomes() ([]model.Income, error) {
	incomes, err := s.incomeRepo.GetIncomes()
	if err != nil {
		return incomes, err
	}
	return incomes, nil

}
