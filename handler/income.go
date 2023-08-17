package handler

import (
	"money-treez/model"
	"money-treez/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type IncomeHandler interface {
	// Create a new income
	CreateIncome(c *gin.Context)
	// Get a income by id
	GetIncome(c *gin.Context)
	// Get all incomes
	GetIncomes(c *gin.Context)
}

type incomeHandler struct {
	incomeService service.IncomeService
}

func NewIncomeHandler(incomeService service.IncomeService) *incomeHandler {
	return &incomeHandler{incomeService}
}

func (h *incomeHandler) CreateIncome(c *gin.Context) {
	var income model.Income
	if err := c.ShouldBindJSON(&income); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	income, err := h.incomeService.CreateIncome(income)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, income)
}

func (h *incomeHandler) GetIncome(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	income, err := h.incomeService.GetIncome(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, income)
}

func (h *incomeHandler) GetIncomes(c *gin.Context) {
	incomes, err := h.incomeService.GetIncomes()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, incomes)
}
