-- name: GetEmployee :one
SELECT id, name, surname, email, created_at, updated_at 
FROM employees 
WHERE id = $1;

-- name: GetEmployees :many
SELECT id, name, surname, email, created_at, updated_at 
FROM employees 
ORDER BY id;

-- name: CreateEmployee :one
INSERT INTO employees (name, surname, email) 
VALUES ($1, $2, $3) 
RETURNING id, name, surname, email, created_at, updated_at;

-- name: UpdateEmployee :one
UPDATE employees 
SET name = $2, surname = $3, email = $4, updated_at = CURRENT_TIMESTAMP 
WHERE id = $1 
RETURNING id, name, surname, email, created_at, updated_at;

-- name: DeleteEmployee :exec
DELETE FROM employees 
WHERE id = $1;

-- name: GetEmployeeByEmail :one
SELECT id, name, surname, email, created_at, updated_at 
FROM employees 
WHERE email = $1;

-- name: GetSale :one
SELECT id, product_name, category, currency, price, sale_date, employee_id, created_at, updated_at 
FROM sales 
WHERE id = $1;

-- name: GetSales :many
SELECT id, product_name, category, currency, price, sale_date, employee_id, created_at, updated_at 
FROM sales 
ORDER BY sale_date DESC;

-- name: GetSalesByEmployee :many
SELECT id, product_name, category, currency, price, sale_date, employee_id, created_at, updated_at 
FROM sales 
WHERE employee_id = $1 
ORDER BY sale_date DESC;

-- name: CreateSale :one
INSERT INTO sales (product_name, category, currency, price, sale_date, employee_id) 
VALUES ($1, $2, $3, $4, $5, $6) 
RETURNING id, product_name, category, currency, price, sale_date, employee_id, created_at, updated_at;

-- name: UpdateSale :one
UPDATE sales 
SET product_name = $2, category = $3, currency = $4, price = $5, sale_date = $6, employee_id = $7, updated_at = CURRENT_TIMESTAMP 
WHERE id = $1 
RETURNING id, product_name, category, currency, price, sale_date, employee_id, created_at, updated_at;

-- name: DeleteSale :exec
DELETE FROM sales 
WHERE id = $1;

-- name: GetSalesByDateRange :many
SELECT id, product_name, category, currency, price, sale_date, employee_id, created_at, updated_at 
FROM sales 
WHERE sale_date BETWEEN $1 AND $2 
ORDER BY sale_date DESC;

-- name: GetSalesByCategory :many
SELECT id, product_name, category, currency, price, sale_date, employee_id, created_at, updated_at 
FROM sales 
WHERE category = $1 
ORDER BY sale_date DESC;

-- name: GetSalesStatsByEmployee :many
SELECT 
    e.id,
    e.name,
    e.surname,
    e.email,
    COUNT(s.id) as total_sales,
    COALESCE(SUM(s.price), 0) as total_revenue,
    COALESCE(AVG(s.price), 0) as avg_sale_value
FROM employees e
LEFT JOIN sales s ON e.id = s.employee_id
GROUP BY e.id, e.name, e.surname, e.email
ORDER BY total_revenue DESC;

-- name: GetEmployeeWithSales :one
SELECT 
    e.id,
    e.name,
    e.surname,
    e.email,
    e.created_at,
    e.updated_at,
    COUNT(s.id) as total_sales,
    COALESCE(SUM(s.price), 0) as total_revenue
FROM employees e
LEFT JOIN sales s ON e.id = s.employee_id
WHERE e.id = $1
GROUP BY e.id, e.name, e.surname, e.email, e.created_at, e.updated_at;
