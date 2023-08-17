package handler

import (
	"money-treez/model"
	"money-treez/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ExpenseHandler interface {
	// Create a new expense
	CreateExpense(c *gin.Context)
	// Get a expense by id
	GetExpense(c *gin.Context)
	// Get all expenses
	GetExpenses(c *gin.Context)
}

type expenseHandler struct {
	expenseService service.ExpenseService
}

func NewExpenseHandler(expenseService service.ExpenseService) *expenseHandler {
	return &expenseHandler{expenseService}
}

func (h *expenseHandler) CreateExpense(c *gin.Context) {
	var expense model.Expense
	if err := c.ShouldBindJSON(&expense); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	expense, err := h.expenseService.CreateExpense(expense)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, expense)

}

func (h *expenseHandler) GetExpense(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	expense, err := h.expenseService.GetExpense(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, expense)
}

func (h *expenseHandler) GetExpenses(c *gin.Context) {
	expenses, err := h.expenseService.GetExpenses()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, expenses)
}
