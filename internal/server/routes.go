package server

import (
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {

	//routes for employee
	e.GET("/employee", GetEmployees)
	e.POST("/employee", CreateEmployee)
	e.PUT("/employee/:id", UpdateEmployee)
	e.DELETE("/employee/:id", DeleteEmployee)

	//routes for employee sales
	e.GET("/sales", GetSales)
	e.POST("/sales", CreateSale)
	e.PUT("/sales/:id", UpdateSale)
	e.DELETE("/sales/:id", DeleteSale)
}

func GetEmployees(c echo.Context) error {
	// Logic to get employees
	return c.JSON(200, "Get Employees")
}
func CreateEmployee(c echo.Context) error {
	// Logic to create an employee
	return c.JSON(201, "Create Employee")
}
func UpdateEmployee(c echo.Context) error {
	// Logic to update an employee
	return c.JSON(200, "Update Employee")
}
func DeleteEmployee(c echo.Context) error {
	// Logic to delete an employee
	return c.JSON(200, "Delete Employee")
}
func GetSales(c echo.Context) error {
	// Logic to get sales
	return c.JSON(200, "Get Sales")
}
func CreateSale(c echo.Context) error {
	// Logic to create a sale
	return c.JSON(201, "Create Sale")
}
func UpdateSale(c echo.Context) error {
	// Logic to update a sale
	return c.JSON(200, "Update Sale")
}
func DeleteSale(c echo.Context) error {
	// Logic to delete a sale
	return c.JSON(200, "Delete Sale")
}
