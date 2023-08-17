package main

import (
	"fmt"
	"money-treez/config"
	"money-treez/db"
	"money-treez/handler"
	"money-treez/model"
	"money-treez/repository"
	"money-treez/service"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ApiHandler struct {
	userHandler     handler.UserHandler
	expenseHandler  handler.ExpenseHandler
	categoryHandler handler.CategoryHandler
	incomeHandler   handler.IncomeHandler
}

func main() {

	gin.SetMode(gin.ReleaseMode) //release

	wg := sync.WaitGroup{}

	config.Init()

	wg.Add(1)

	go func() {
		defer wg.Done()

		router := gin.New()
		db := db.NewDB()
		router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
			return fmt.Sprintf("[%s] \"%s %s %s\"\n",
				param.TimeStamp.Format(time.RFC822),
				param.Method,
				param.Path,
				param.ErrorMessage,
			)
		}))
		router.Use(gin.Recovery())

		dbCredential := model.DBCredentials{
			Host:     config.Config.DBHost,
			User:     config.Config.DBUsername,
			Password: config.Config.DBPassword,
			DBName:   config.Config.DBName,
			Port:     config.Config.DBPort,
		}

		conn, err := db.Connect(&dbCredential)
		if err != nil {
			panic(err)
		}

		conn.AutoMigrate(
			&model.User{},
			&model.Expense{},
			&model.ExpenseCategory{},
			&model.Income{},
			&model.Session{},
		)

		router = RunServer(conn, router)
		// router = RunClient(conn, router, Resources)

		fmt.Println("Server is running on port 8080")
		err = router.Run(":8080")
		if err != nil {
			panic(err)
		}
	}()

	wg.Wait()

}

func RunServer(db *gorm.DB, routers *gin.Engine) *gin.Engine {

	userRepo := repository.NewUserRepository(db)
	expenseRepo := repository.NewExpenseRepository(db)
	categoryRepo := repository.NewCategoryRepository(db)
	incomeRepo := repository.NewIncomeRepository(db)
	sessionRepo := repository.NewSessionsRepository(db)

	userService := service.NewUserService(userRepo, sessionRepo)
	expenseService := service.NewExpenseService(expenseRepo)
	categoryService := service.NewCategoryService(categoryRepo)
	incomeService := service.NewIncomeService(incomeRepo)

	userHandler := handler.NewUserHandler(userService)
	expenseHandler := handler.NewExpenseHandler(expenseService)
	categoryHandler := handler.NewCategoryHandler(categoryService)
	incomeHandler := handler.NewIncomeHandler(incomeService)

	apiHandler := ApiHandler{
		userHandler:     userHandler,
		expenseHandler:  expenseHandler,
		categoryHandler: categoryHandler,
		incomeHandler:   incomeHandler,
	}

	user := routers.Group("/user")
	{
		user.POST("/register", apiHandler.userHandler.CreateUser)
		user.GET("/:id", apiHandler.userHandler.GetUser)
	}

	expense := routers.Group("/expense")
	{
		expense.POST("/create", apiHandler.expenseHandler.CreateExpense)
		expense.GET("/list", apiHandler.expenseHandler.GetExpenses)
		expense.GET("/detail/:id", apiHandler.expenseHandler.GetExpense)
		expense.POST("/category/create", apiHandler.categoryHandler.CreateExpenseCategory)
		expense.GET("/category/list", apiHandler.categoryHandler.GetExpenseCategories)
	}

	income := routers.Group("/income")
	{
		income.POST("/create", apiHandler.incomeHandler.CreateIncome)
		income.GET("/list", apiHandler.incomeHandler.GetIncomes)
		income.GET("/detail/:id", apiHandler.incomeHandler.GetIncome)

	}

	return routers
}
