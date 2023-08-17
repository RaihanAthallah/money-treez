package handler

import (
	"money-treez/model"
	"money-treez/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CategoryHandler interface {
	// Create a new category
	CreateExpenseCategory(c *gin.Context)
	// Get all categories
	GetExpenseCategories(c *gin.Context)
}

type categoryHandler struct {
	categoryService service.CategoryService
}

func NewCategoryHandler(categoryService service.CategoryService) *categoryHandler {
	return &categoryHandler{categoryService}
}

func (h *categoryHandler) CreateExpenseCategory(c *gin.Context) {
	var expenseCategory model.ExpenseCategory
	if err := c.ShouldBindJSON(&expenseCategory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	expenseCategory, err := h.categoryService.CreateExpenseCategory(expenseCategory)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, expenseCategory)

}

func (h *categoryHandler) GetExpenseCategories(c *gin.Context) {
	expenseCategories, err := h.categoryService.GetExpenseCategories()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, expenseCategories)
}
