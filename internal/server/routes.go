package server

import (
	internals "WorkRESTAPI/internal"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

var queries *internals.Queries

func RegisterRoutes(e *echo.Echo, q *internals.Queries) {
	queries = q

	//routes for employee
	e.GET("/employee", GetEmployee)
	e.GET("/employees", GetAllEmployees)
	e.POST("/employee", CreateEmployee)
	e.PUT("/employee/:id", UpdateEmployee)
	e.DELETE("/employee/:id", DeleteEmployee)

	//routes for employee sales
	e.GET("/sale", GetSales)
	e.GET("sales", GetAllSales)
	e.POST("/sale", CreateSale)
	e.PUT("/sale/:id", UpdateSale)
	e.DELETE("/sale/:id", DeleteSale)

}

func GetEmployee(c echo.Context) error {
	// Logic to get employee by ID
	ctx := c.Request().Context()

	idStr := c.QueryParam("id")
	if idStr == "" {
		return c.JSON(400, map[string]string{"error": "Employee ID is required"})
	}
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		return c.JSON(400, map[string]string{"error": "Invalid ID format"})
	}
	employee, err := queries.GetEmployee(ctx, int32(id))
	if err != nil {
		return c.JSON(404, map[string]string{"error": "Employee not found"})
	}

	return c.JSON(http.StatusOK, employee)
}
func CreateEmployee(c echo.Context) error {
	// Logic to create an employee

	ctx := c.Request().Context()
	var employeeParams internals.CreateEmployeeParams

	if err := c.Bind(&employeeParams); err == nil {
		if employeeParams.Name != "" && employeeParams.Surname != "" && employeeParams.Email != "" {
			employee, err := queries.CreateEmployee(ctx, employeeParams)
			if err != nil {
				return c.JSON(500, map[string]string{"error": "Failed to create employee"})
			}
			return c.JSON(http.StatusCreated, employee)
		}
	}
	name := c.QueryParam("name")
	surname := c.QueryParam("surname")
	email := c.QueryParam("email")

	if name != "" && surname != "" && email != "" {
		employeeParams = internals.CreateEmployeeParams{
			Name:    name,
			Surname: surname,
			Email:   email,
		}

		employee, err := queries.CreateEmployee(ctx, employeeParams)
		if err != nil {
			return c.JSON(500, map[string]string{"error": "Failed to create employee"})
		}
		return c.JSON(http.StatusCreated, employee)
	}

	return c.JSON(400, map[string]string{
		"error": "Provide employee data either as JSON body or query parameters (name, surname, email)",
	})
}
func UpdateEmployee(c echo.Context) error {
	ctx := c.Request().Context()

	idStr := c.Param("id")
	if idStr == "" {
		return c.JSON(400, map[string]string{"error": "Employee ID is required"})
	}

	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		return c.JSON(400, map[string]string{"error": "Invalid ID format"})
	}

	currentEmployee, err := queries.GetEmployee(ctx, int32(id))
	if err != nil {
		return c.JSON(404, map[string]string{"error": "Employee not found"})
	}

	updateParams := internals.UpdateEmployeeParams{
		ID:      currentEmployee.ID,
		Name:    currentEmployee.Name,
		Surname: currentEmployee.Surname,
		Email:   currentEmployee.Email,
	}

	var jsonParams internals.UpdateEmployeeParams
	if err := c.Bind(&jsonParams); err == nil {
		if jsonParams.Name != "" {
			updateParams.Name = jsonParams.Name
		}
		if jsonParams.Surname != "" {
			updateParams.Surname = jsonParams.Surname
		}
		if jsonParams.Email != "" {
			updateParams.Email = jsonParams.Email
		}
	}

	if name := c.QueryParam("name"); name != "" {
		updateParams.Name = name
	}
	if surname := c.QueryParam("surname"); surname != "" {
		updateParams.Surname = surname
	}
	if email := c.QueryParam("email"); email != "" {
		updateParams.Email = email
	}
	if updateParams.Name == "" && updateParams.Surname == "" && updateParams.Email == "" {
		return c.JSON(400, map[string]string{"error": "No fields to update"})
	}
	updatedEmployee, err := queries.UpdateEmployee(ctx, updateParams)
	if err != nil {
		return c.JSON(500, map[string]string{"error": "Failed to update employee"})
	}
	return c.JSON(http.StatusOK, updatedEmployee)

}
func DeleteEmployee(c echo.Context) error {
	// Logic to delete an employee
	ctx := c.Request().Context()
	idStr := c.Param("id")
	if idStr == "" {
		return c.JSON(400, map[string]string{"error": "Employee ID is required"})
	}
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		return c.JSON(400, map[string]string{"error": "Invalid ID format"})
	}
	err = queries.DeleteEmployee(ctx, int32(id))
	if err != nil {
		return c.JSON(404, map[string]string{"error": "Employee not found"})
	}

	return c.JSON(200, "Delete Employee")
}

func GetAllEmployees(c echo.Context) error {
	ctx := c.Request().Context()
	employees, err := queries.GetEmployees(ctx)
	if err != nil {
		return c.JSON(500, map[string]string{"error": "Failed to get employees"})
	}
	return c.JSON(200, employees)

}

func GetSales(c echo.Context) error {
	ctx := c.Request().Context()
	idStr := c.QueryParam("id")
	if idStr == "" {
		return c.JSON(400, map[string]string{"error": "Sale ID is required"})
	}
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		return c.JSON(400, map[string]string{"error": "Invalid ID format"})
	}
	sale, err := queries.GetSale(ctx, int32(id))
	if err != nil {
		return c.JSON(404, map[string]string{"error": "Sale not found"})
	}
	return c.JSON(200, sale)
}
func CreateSale(c echo.Context) error {
	ctx := c.Request().Context()

	type CreateSaleRequest struct {
		ProductName string    `json:"product_name"`
		Category    string    `json:"category"`
		Currency    string    `json:"currency"`
		Price       float64   `json:"price"`
		SaleDate    time.Time `json:"sale_date"`
		EmployeeID  int32     `json:"employee_id"`
	}

	var req CreateSaleRequest
	if err := c.Bind(&req); err == nil {
		if req.EmployeeID != 0 && req.Price > 0 {
			saleParams := internals.CreateSaleParams{
				ProductName: req.ProductName,
				Category:    req.Category,
				Currency:    req.Currency,
				Price:       fmt.Sprintf("%.2f", req.Price), // Converting float64 -> string
				SaleDate:    req.SaleDate,
				EmployeeID:  req.EmployeeID,
			}

			sale, err := queries.CreateSale(ctx, saleParams)
			if err != nil {
				return c.JSON(500, map[string]string{"error": "Failed to create sale"})
			}
			return c.JSON(201, sale)
		}
	}

	productName := c.QueryParam("product_name")
	category := c.QueryParam("category")
	currency := c.QueryParam("currency")
	priceStr := c.QueryParam("price")
	employeeIDStr := c.QueryParam("employee_id")

	if productName != "" && category != "" && priceStr != "" && employeeIDStr != "" {
		price, err := strconv.ParseFloat(priceStr, 64)
		if err != nil {
			return c.JSON(400, map[string]string{"error": "Invalid price format"})
		}

		employeeID, err := strconv.ParseInt(employeeIDStr, 10, 32)
		if err != nil {
			return c.JSON(400, map[string]string{"error": "Invalid employee_id format"})
		}

		if currency == "" {
			currency = "PLN"
		}

		saleParams := internals.CreateSaleParams{
			ProductName: productName,
			Category:    category,
			Currency:    currency,
			Price:       fmt.Sprintf("%.2f", price), // Converting float64 -> string
			SaleDate:    time.Now(),
			EmployeeID:  int32(employeeID),
		}

		sale, err := queries.CreateSale(ctx, saleParams)
		if err != nil {
			return c.JSON(500, map[string]string{"error": "Failed to create sale"})
		}
		return c.JSON(201, sale)
	}

	return c.JSON(400, map[string]string{
		"error": "Provide sale data either as JSON body or query parameters (product_name, category, price, employee_id)",
	})
}
func UpdateSale(c echo.Context) error {
	ctx := c.Request().Context()

	idStr := c.Param("id")
	if idStr == "" {
		return c.JSON(400, map[string]string{"error": "Sale ID is required"})
	}

	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		return c.JSON(400, map[string]string{"error": "Invalid ID format"})
	}

	currentSale, err := queries.GetSale(ctx, int32(id))
	if err != nil {
		return c.JSON(404, map[string]string{"error": "Sale not found"})
	}

	updateParams := internals.UpdateSaleParams{
		ID:          currentSale.ID,
		ProductName: currentSale.ProductName,
		Category:    currentSale.Category,
		Currency:    currentSale.Currency,
		Price:       currentSale.Price,
		SaleDate:    currentSale.SaleDate,
		EmployeeID:  currentSale.EmployeeID,
	}

	var jsonParams internals.UpdateSaleParams
	if err := c.Bind(&jsonParams); err == nil {
		if jsonParams.ProductName != "" {
			updateParams.ProductName = jsonParams.ProductName
		}
		if jsonParams.Category != "" {
			updateParams.Category = jsonParams.Category
		}
		if jsonParams.Currency != "" {
			updateParams.Currency = jsonParams.Currency
		}
		if jsonParams.Price != "" {
			updateParams.Price = jsonParams.Price
		}
		if !jsonParams.SaleDate.IsZero() {
			updateParams.SaleDate = jsonParams.SaleDate
		}
		if jsonParams.EmployeeID != 0 {
			updateParams.EmployeeID = jsonParams.EmployeeID
		}
	}

	if productName := c.QueryParam("product_name"); productName != "" {
		updateParams.ProductName = productName
	}

	if category := c.QueryParam("category"); category != "" {
		updateParams.Category = category
	}

	if currency := c.QueryParam("currency"); currency != "" {
		updateParams.Currency = currency
	}

	if priceStr := c.QueryParam("price"); priceStr != "" {
		price, err := strconv.ParseFloat(priceStr, 64)
		if err != nil {
			return c.JSON(400, map[string]string{"error": "Invalid price format"})
		}
		updateParams.Price = fmt.Sprintf("%.2f", price)
	}

	if saleDateStr := c.QueryParam("sale_date"); saleDateStr != "" {
		saleDate, err := time.Parse(time.RFC3339, saleDateStr)
		if err != nil {
			return c.JSON(400, map[string]string{"error": "Invalid sale_date format"})
		}
		updateParams.SaleDate = saleDate
	}

	if employeeIDStr := c.QueryParam("employee_id"); employeeIDStr != "" {
		employeeID, err := strconv.ParseInt(employeeIDStr, 10, 32)
		if err != nil {
			return c.JSON(400, map[string]string{"error": "Invalid employee_id format"})
		}
		updateParams.EmployeeID = int32(employeeID)
	}

	if updateParams.ProductName == "" && updateParams.Category == "" && updateParams.Currency == "" && updateParams.Price == "" && updateParams.SaleDate.IsZero() && updateParams.EmployeeID == 0 {
		return c.JSON(400, map[string]string{"error": "No fields to update"})
	}

	updatedSale, err := queries.UpdateSale(ctx, updateParams)

	return c.JSON(http.StatusOK, updatedSale)
}
func DeleteSale(c echo.Context) error {
	// Logic to delete a sale
	ctx := c.Request().Context()
	idStr := c.Param("id")
	if idStr == "" {
		return c.JSON(400, map[string]string{"error": "Sale ID is required"})
	}
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		return c.JSON(400, map[string]string{"error": "Invalid ID format"})
	}
	err = queries.DeleteSale(ctx, int32(id))
	if err != nil {
		return c.JSON(404, map[string]string{"error": "Sale not found"})
	}

	return c.JSON(200, map[string]string{"message": "Sale deleted successfully"})
}
func GetAllSales(c echo.Context) error {
	ctx := c.Request().Context()
	sales, err := queries.GetSales(ctx)
	if err != nil {
		return c.JSON(500, map[string]string{"error": "Failed to get sales"})
	}
	return c.JSON(200, sales)
}
