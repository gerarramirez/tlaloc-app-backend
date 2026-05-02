package main

import (
	"github.com/labstack/echo/v4"
	"tlaloc-catalog-service/handler"
)

func RegisterRoutes(e *echo.Echo, h *handler.Handler) {

	e.POST("/expenses", h.CreateExpense)
	e.PUT("/expenses/:id", h.UpdateExpense)
	e.GET("/expenses", h.GetExpenses)

	e.POST("/expenses-categories", h.CreateExpenseCategory)
	e.PUT("/expenses-categories/:id", h.UpdateExpenseCategory)
	e.GET("/expenses-categories", h.GetExpenseCategories)

	e.POST("/commerce-categories", h.CreateCommerceCategory)
	e.PUT("/commerce-categories/:id", h.UpdateCommerceCategory)
	e.GET("/commerce-categories", h.GetCommerceCategories)

	e.POST("/commerce-subcategories", h.CreateCommerceSubcategory)
	e.PUT("/commerce-subcategories/:id", h.UpdateCommerceSubcategory)
	e.GET("/commerce-subcategories", h.GetCommerceSubcategories)

	e.POST("/commerces", h.CreateCommerce)
	e.PUT("/commerces/:id", h.UpdateCommerce)
	e.GET("/commerces", h.GetCommerces)

	e.POST("/banks", h.CreateBank)
	e.PUT("/banks/:id", h.UpdateBank)
	e.GET("/banks", h.GetBanks)

	e.POST("/bank-products", h.CreateBankProduct)
	e.PUT("/bank-products/:id", h.UpdateBankProduct)
	e.GET("/bank-products", h.GetBankProducts)

	e.POST("/income-types", h.CreateIncomeType)
	e.PUT("/income-types/:id", h.UpdateIncomeType)
	e.GET("/income-types", h.GetIncomeTypes)

	e.POST("/product-types", h.CreateProductType)
	e.PUT("/product-types/:id", h.UpdateProductType)
	e.GET("/product-types", h.GetProductTypes)

}
