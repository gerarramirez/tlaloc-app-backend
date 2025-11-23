package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"tlaloc-api-budget/model"
)

func (handler *Handler) CreateBudget(c echo.Context) error {
	u := new(model.Budget)

	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	error := handler.budgetDao.CreateBudget(u)

	if error != nil {
		c.JSON(http.StatusInternalServerError, "Error saving budget")
	}

	return c.JSON(http.StatusOK, "Registered of buget is done")

}

func (handler *Handler) CreateBudgetExpenseCate(c echo.Context) error {
	b := new(model.BudgetExpenseCategories)
	if err := c.Bind(b); err != nil {
		return c.JSON(http.StatusInternalServerError, "Error Parsing JSON")
	}

	if result := handler.budgetDao.CreateBudgetExpenseCate(b); result != nil {
		return c.JSON(http.StatusInternalServerError, "Internal Server Error")
	}

	return c.JSON(http.StatusOK, "SUCCCESS")

}

func (handler *Handler) GetWholeBudget(c echo.Context) error {
	var result *model.BudgetWithWholeExpensesCategory
	var id string

	if err := echo.PathParamsBinder(c).String("id", &id).BindErrors(); err != nil {
		return c.JSON(http.StatusInternalServerError, "Internal Server Error")
	}

	result, err := handler.budgetDao.GetWholeBudget(&id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Error in the server")
	}

	return c.JSON(http.StatusOK, result)
}
