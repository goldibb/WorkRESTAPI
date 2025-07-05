package server

import (
	internals "WorkRESTAPI/internal"
	"bytes"
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/phpdave11/gofpdf"
)

var queries *internals.Queries

// Email validation regex
var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

// Helper function to validate email format
func isValidEmail(email string) bool {
	return emailRegex.MatchString(email)
}

// Helper function to check if email already exists (excluding current employee ID for updates)
func isEmailExists(ctx echo.Context, email string, excludeID ...int32) (bool, error) {
	_, err := queries.GetEmployeeByEmail(ctx.Request().Context(), email)
	if err != nil {
		// If error is "no rows found", email doesn't exist
		return false, nil
	}

	// If we have an excludeID (for updates), check if it's the same employee
	if len(excludeID) > 0 {
		employee, err := queries.GetEmployeeByEmail(ctx.Request().Context(), email)
		if err != nil {
			return false, err
		}
		// If the email belongs to the same employee being updated, it's ok
		return employee.ID != excludeID[0], nil
	}

	// Email exists and it's not an update scenario
	return true, nil
}

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
	e.GET("/sales", GetAllSales)
	e.POST("/sale", CreateSale)
	e.PUT("/sale/:id", UpdateSale)
	e.DELETE("/sale/:id", DeleteSale)

	e.GET("/employee/:id/report/month", GenerateEmployeeMonthlyReport)
	e.GET("/employee/:id/report/quarter", GenerateEmployeeQuarterlyReport)
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
			// Validate email format
			if !isValidEmail(employeeParams.Email) {
				return c.JSON(400, map[string]string{"error": "Invalid email format"})
			}

			// Check if email already exists
			emailExists, err := isEmailExists(c, employeeParams.Email)
			if err != nil {
				return c.JSON(500, map[string]string{"error": "Failed to check email existence"})
			}
			if emailExists {
				return c.JSON(409, map[string]string{"error": "Email already exists"})
			}

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
		// Validate email format
		if !isValidEmail(email) {
			return c.JSON(400, map[string]string{"error": "Invalid email format"})
		}

		// Check if email already exists
		emailExists, err := isEmailExists(c, email)
		if err != nil {
			return c.JSON(500, map[string]string{"error": "Failed to check email existence"})
		}
		if emailExists {
			return c.JSON(409, map[string]string{"error": "Email already exists"})
		}

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
			// Validate email format
			if !isValidEmail(jsonParams.Email) {
				return c.JSON(400, map[string]string{"error": "Invalid email format"})
			}

			// Check if email already exists (excluding current employee)
			emailExists, err := isEmailExists(c, jsonParams.Email, int32(id))
			if err != nil {
				return c.JSON(500, map[string]string{"error": "Failed to check email existence"})
			}
			if emailExists {
				return c.JSON(409, map[string]string{"error": "Email already exists"})
			}

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
		// Validate email format
		if !isValidEmail(email) {
			return c.JSON(400, map[string]string{"error": "Invalid email format"})
		}

		// Check if email already exists (excluding current employee)
		emailExists, err := isEmailExists(c, email, int32(id))
		if err != nil {
			return c.JSON(500, map[string]string{"error": "Failed to check email existence"})
		}
		if emailExists {
			return c.JSON(409, map[string]string{"error": "Email already exists"})
		}

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
			_, err := queries.GetEmployee(ctx, req.EmployeeID)
			if err != nil {
				return c.JSON(400, map[string]string{"error": "Employee not found"})
			}
			if req.ProductName == "" {
				return c.JSON(400, map[string]string{"error": "Product name is required"})
			}
			if req.SaleDate.After(time.Now()) {
				return c.JSON(400, map[string]string{"error": "Sale date cannot be in the future"})
			}
			if req.Category == "" {
				return c.JSON(400, map[string]string{"error": "Category is required"})
			}
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

		if price <= 0 {
			return c.JSON(400, map[string]string{"error": "Price must be greater than 0"})
		}

		// Check if employee exists
		_, err = queries.GetEmployee(ctx, int32(employeeID))
		if err != nil {
			return c.JSON(400, map[string]string{"error": "Employee not found"})
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
			// Validate price format and value
			price, err := strconv.ParseFloat(jsonParams.Price, 64)
			if err != nil {
				return c.JSON(400, map[string]string{"error": "Invalid price format"})
			}
			if price <= 0 {
				return c.JSON(400, map[string]string{"error": "Price must be greater than 0"})
			}
			updateParams.Price = jsonParams.Price
		}
		if !jsonParams.SaleDate.IsZero() {
			// Validate sale date is not in the future
			if jsonParams.SaleDate.After(time.Now()) {
				return c.JSON(400, map[string]string{"error": "Sale date cannot be in the future"})
			}
			updateParams.SaleDate = jsonParams.SaleDate
		}
		if jsonParams.EmployeeID != 0 {
			// Check if employee exists
			_, err := queries.GetEmployee(ctx, jsonParams.EmployeeID)
			if err != nil {
				return c.JSON(400, map[string]string{"error": "Employee not found"})
			}
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
		if price <= 0 {
			return c.JSON(400, map[string]string{"error": "Price must be greater than 0"})
		}
		updateParams.Price = fmt.Sprintf("%.2f", price)
	}

	if saleDateStr := c.QueryParam("sale_date"); saleDateStr != "" {
		saleDate, err := time.Parse(time.RFC3339, saleDateStr)
		if err != nil {
			return c.JSON(400, map[string]string{"error": "Invalid sale_date format"})
		}
		// Validate sale date is not in the future
		if saleDate.After(time.Now()) {
			return c.JSON(400, map[string]string{"error": "Sale date cannot be in the future"})
		}
		updateParams.SaleDate = saleDate
	}

	if employeeIDStr := c.QueryParam("employee_id"); employeeIDStr != "" {
		employeeID, err := strconv.ParseInt(employeeIDStr, 10, 32)
		if err != nil {
			return c.JSON(400, map[string]string{"error": "Invalid employee_id format"})
		}
		// Check if employee exists
		_, err = queries.GetEmployee(ctx, int32(employeeID))
		if err != nil {
			return c.JSON(400, map[string]string{"error": "Employee not found"})
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

func GenerateEmployeeMonthlyReport(c echo.Context) error {
	ctx := c.Request().Context()

	idStr, yearStr, monthStr := c.Param("id"), c.QueryParam("year"), c.QueryParam("month")

	if idStr == "" || yearStr == "" || monthStr == "" {
		return c.JSON(400, map[string]string{
			"error": "Employee ID, year and month are required",
		})
	}
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		return c.JSON(400, map[string]string{"error": "Invalid employee ID"})
	}

	year, err := strconv.Atoi(yearStr)
	if err != nil {
		return c.JSON(400, map[string]string{"error": "Invalid year"})
	}

	month, err := strconv.Atoi(monthStr)
	if err != nil || month < 1 || month > 12 {
		return c.JSON(400, map[string]string{"error": "Invalid month (1-12)"})
	}
	employee, err := queries.GetEmployee(ctx, int32(id))
	if err != nil {
		return c.JSON(404, map[string]string{"error": "Employee not found"})
	}
	startDate := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
	endDate := startDate.AddDate(0, 1, -1).Add(23*time.Hour + 59*time.Minute + 59*time.Second)

	sales, err := queries.GetSalesByDateRange(ctx, internals.GetSalesByDateRangeParams{
		SaleDate:   startDate,
		SaleDate_2: endDate,
	})
	if err != nil {
		return c.JSON(500, map[string]string{"error": "Failed to get sales data"})
	}
	var employeeSales []internals.Sale
	for _, sale := range sales {
		if sale.EmployeeID == int32(id) {
			employeeSales = append(employeeSales, sale)
		}
	}

	// GENERATE PDF PDF
	pdf := generateMonthlyReportPDF(employee, employeeSales, year, month)

	// RETURNS PDF
	c.Response().Header().Set("Content-Type", "application/pdf")
	c.Response().Header().Set("Content-Disposition",
		fmt.Sprintf("attachment; filename=\"raport_%s_%s_%d_%d.pdf\"",
			employee.Name, employee.Surname, year, month))

	return c.Stream(200, "application/pdf", pdf)

}

func GenerateEmployeeQuarterlyReport(c echo.Context) error {
	ctx := c.Request().Context()

	idStr, yearStr, quarterStr := c.Param("id"), c.QueryParam("year"), c.QueryParam("quarter")

	if idStr == "" || yearStr == "" || quarterStr == "" {
		return c.JSON(400, map[string]string{
			"error": "Employee ID, year and month are required",
		})
	}
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		return c.JSON(400, map[string]string{"error": "Invalid employee ID"})
	}

	year, err := strconv.Atoi(yearStr)
	if err != nil {
		return c.JSON(400, map[string]string{"error": "Invalid year"})
	}

	employee, err := queries.GetEmployee(ctx, int32(id))
	if err != nil {
		return c.JSON(404, map[string]string{"error": "Employee not found"})
	}
	if year < employee.CreatedAt.Time.Year() || year > time.Now().Year() {
		return c.JSON(400, map[string]string{"error": "Invalid year"})
	}

	quarter, err := strconv.Atoi(quarterStr)
	if err != nil || quarter < 1 || quarter > 4 {
		return c.JSON(400, map[string]string{"error": "Invalid month (1-4)"})
	}

	startMonth := (quarter-1)*3 + 1
	startDate := time.Date(year, time.Month(startMonth), 1, 0, 0, 0, 0, time.UTC)
	endDate := startDate.AddDate(0, 3, -1).Add(23*time.Hour + 59*time.Minute + 59*time.Second)

	sales, err := queries.GetSalesByDateRange(ctx, internals.GetSalesByDateRangeParams{
		SaleDate:   startDate,
		SaleDate_2: endDate,
	})
	if err != nil {
		return c.JSON(500, map[string]string{"error": "Failed to get sales data"})
	}
	var employeeSales []internals.Sale
	for _, sale := range sales {
		if sale.EmployeeID == int32(id) {
			employeeSales = append(employeeSales, sale)
		}
	}

	// GENERATE PDF PDF
	pdf := generateQuarterlyReportPDF(employee, employeeSales, year, quarter)

	// RETURNS PDF
	c.Response().Header().Set("Content-Type", "application/pdf")
	c.Response().Header().Set("Content-Disposition",
		fmt.Sprintf("attachment; filename=\"raport_%s_%s_%d_%d.pdf\"",
			employee.Name, employee.Surname, year, quarter))

	return c.Stream(200, "application/pdf", pdf)
}

func generateMonthlyReportPDF(employee internals.Employee, sales []internals.Sale, year, month int) *bytes.Buffer {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	// header
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(0, 10, fmt.Sprintf("Monthly report - %s %s", employee.Name, employee.Surname))
	pdf.Ln(10)

	pdf.SetFont("Arial", "", 12)
	pdf.Cell(0, 10, fmt.Sprintf("Period: %s %d", time.Month(month).String(), year))
	pdf.Ln(10)

	// Statistics
	totalSales := len(sales)
	var totalRevenue float64
	for _, sale := range sales {
		price, _ := strconv.ParseFloat(sale.Price, 64)
		totalRevenue += price
	}

	pdf.Cell(0, 10, fmt.Sprintf("Number of sales: %d", totalSales))
	pdf.Ln(5)
	pdf.Cell(0, 10, fmt.Sprintf("Total revenue: %.2f PLN", totalRevenue))
	pdf.Ln(10)

	// Sales table
	if len(sales) > 0 {
		pdf.SetFont("Arial", "B", 10)
		pdf.Cell(30, 10, "Date")
		pdf.Cell(50, 10, "Product")
		pdf.Cell(30, 10, "Category")
		pdf.Cell(30, 10, "Prize")
		pdf.Ln(10)

		pdf.SetFont("Arial", "", 9)
		for _, sale := range sales {
			pdf.Cell(30, 8, sale.SaleDate.Format("2006-01-02"))
			pdf.Cell(50, 8, sale.ProductName)
			pdf.Cell(30, 8, sale.Category)
			pdf.Cell(30, 8, sale.Price+" "+sale.Currency)
			pdf.Ln(8)
		}
	}

	var buf bytes.Buffer
	pdf.Output(&buf)
	return &buf
}

func generateQuarterlyReportPDF(employee internals.Employee, sales []internals.Sale, year, quarter int) *bytes.Buffer {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	// header
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(0, 10, fmt.Sprintf("Quarterly report- %s %s", employee.Name, employee.Surname))
	pdf.Ln(10)

	pdf.SetFont("Arial", "", 12)
	pdf.Cell(0, 10, fmt.Sprintf("Period: Q%d %d", quarter, year))
	pdf.Ln(10)

	// Statistics
	totalSales := len(sales)
	var totalRevenue float64
	for _, sale := range sales {
		price, _ := strconv.ParseFloat(sale.Price, 64)
		totalRevenue += price
	}

	pdf.Cell(0, 10, fmt.Sprintf("Number of sales: %d", totalSales))
	pdf.Ln(5)
	pdf.Cell(0, 10, fmt.Sprintf("Total revenue: %.2f PLN", totalRevenue))
	pdf.Ln(10)

	if len(sales) > 0 {
		pdf.SetFont("Arial", "B", 10)
		pdf.Cell(30, 10, "Date")
		pdf.Cell(50, 10, "Product")
		pdf.Cell(30, 10, "Category")
		pdf.Cell(30, 10, "Price")
		pdf.Ln(10)

		pdf.SetFont("Arial", "", 9)
		for _, sale := range sales {
			pdf.Cell(30, 8, sale.SaleDate.Format("2006-01-02"))
			pdf.Cell(50, 8, sale.ProductName)
			pdf.Cell(30, 8, sale.Category)
			pdf.Cell(30, 8, sale.Price+" "+sale.Currency)
			pdf.Ln(8)
		}
	}

	var buf bytes.Buffer
	pdf.Output(&buf)
	return &buf
}
