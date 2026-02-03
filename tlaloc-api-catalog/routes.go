package main

import (
	"github.com/labstack/echo/v4"
	"tlaloc-catalog/handler"
)

func RegisterRoutes(e *echo.Echo, h *handler.Handler) {

	e.POST("/expenses/create", h.CreateExpenses)
	e.POST("/expenses/update", h.UpdateExpenses)
	e.GET("/expenses/findAll", h.FindAllExpenses)

	e.POST("/expenses-categories/create", h.CreateExpensesCategories)
	e.POST("/expenses-categories/update", h.UpdateExpensesCategories)
	e.GET("/expenses-categories/findAll", h.FindAllExpensesCategories)

	e.POST("/commerce-categories/create", h.CreateCommercesCategories)
	e.POST("/commerce-categories/update", h.UpdateCommercesCategories)
	e.GET("/commerce-categories/findAll", h.FindAllCommercesCategories)

	e.POST("/commerce-subcategories/create", h.CreateCommercesSubcategories)
	e.POST("/commerce-subcategories/update", h.UpdateCommercesSubcategories)
	e.GET("/commerce-subcategories/findAll", h.FindAllCommercesSubcategories)

	e.POST("/commerces/create", h.CreateCommerce)
	e.POST("/commerces/update", h.UpdateCommerces)
	e.GET("/commerces/findAll", h.FindAllCommerces)

	e.POST("/bank/create", h.Create)
	e.GET("/bank/findAll", h.FindAll)
	e.POST("/bank/update", h.Update)
	e.POST("/bank-products/create", h.CreateBanksProduct)
	e.GET("/bank-products/findAll", h.FindAllBanksProducts)
	e.POST("/bank-products/update", h.UpdateBanksProducts)

	e.POST("/income-types/create", h.CreateIncomeTypes)
	e.POST("/income-types/update", h.UpdateIncomeTypes)
	e.GET("/income-types/findAll", h.FindAllIncomeType)

	e.POST("/product-types/create", h.CreateProductType)
	e.GET("/product-types/GetAll", h.GetAllProductType)
	e.POST("/product-types/update", h.UpdateProductType)

	e.POST("/interest-rate/create", h.CreateInterestRate)
	e.GET("/interest-rate/findAll", h.FindAllInterestRates)
	e.POST("/interest-rate/update", h.UpdateInterestRate)

}
