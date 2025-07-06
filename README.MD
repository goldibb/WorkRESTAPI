# Work REST API 🚀

**Modern REST API for employee and sales management with automatic PDF report generation**

![Go](https://img.shields.io/badge/Go-1.24-00ADD8?style=for-the-badge&logo=go)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-Latest-336791?style=for-the-badge&logo=postgresql)
![Docker](https://img.shields.io/badge/Docker-Compose-2496ED?style=for-the-badge&logo=docker)
![Echo](https://img.shields.io/badge/Echo-v4-00ADD8?style=for-the-badge)

## 📋 Table of Contents

- [Quick Start](#-quick-start)
- [Features](#-features)
- [Technology Stack](#-technology-stack)
- [Architecture](#-architecture)
- [Installation](#-installation)
- [Usage](#-usage)
- [API Endpoints](#-api-endpoints)
- [Examples](#-examples)
- [Testing](#-testing)
- [Test Data](#-test-data)
- [Troubleshooting](#-troubleshooting)

## ⚡ Quick Start

**Want to try it immediately?**

```bash
# 1. Clone and enter directory
git clone https://github.com/goldibb/WorkRESTAPI
cd WorkRESTAPI

# 2. Create environment file
cp .env.example .env

# 3. Enable test data loading (uncomment in docker-compose.yml)
# - ./test_data.sql:/docker-entrypoint-initdb.d/02-test-data.sql

# 4. Start everything
docker-compose up -d

# 5. Test it works (wait 30 seconds for startup)
curl http://localhost:1323/employees
```

**That's it! 🎉 Your API is running with sample data.**

## ✨ Features

### 👥 Employee Management (CRUD)
- ✅ Add, edit, delete employees
- ✅ Data validation (email, field length)
- ✅ Email uniqueness check
- ✅ Automatic timestamps (created_at, updated_at)

### 💰 Sales Management (CRUD)
- ✅ Add, edit, delete sales
- ✅ Multiple currency support (PLN, EUR, USD)
- ✅ Price validation (must be > 0)
- ✅ **Flexible date formats** - supports multiple formats:
  - ISO 8601: `2025-01-15T10:30:00Z`
  - RFC 3339: `2025-01-15T10:30:00+01:00`
  - Polish format: `15.01.2025 10:30`
  - European: `15-01-2025 10:30`
  - American: `01/15/2025 10:30`

### 📊 PDF Reports
- ✅ **Monthly reports** for employees
- ✅ **Quarterly reports** for employees
- ✅ Automatic PDF generation with gofpdf
- ✅ Statistics: sales count, total revenue
- ✅ Detailed tables of all transactions

### 🔒 Security and Validation
- ✅ Email validation (regex)
- ✅ Field length validation
- ✅ Email uniqueness check
- ✅ Date and time range validation
- ✅ Price and currency validation

## 🛠 Technology Stack

| Component | Technology | Version |
|-----------|------------|---------|
| **Backend** | Go | 1.24 |
| **Framework** | Echo | v4 |
| **Database** | PostgreSQL | Latest |
| **ORM/Query Builder** | sqlc | Latest |
| **DB Connection** | pgx/v5 | v5.7.5 |
| **PDF Generator** | gofpdf | v1.4.3 |
| **Containerization** | Docker + Docker Compose | Latest |

## 🏗 Architecture

```
WorkRESTAPI/
├── cmd/api/                    # Application entry point
│   └── main.go
├── internal/                   # Business logic
│   ├── db.go                   # Database connection
│   ├── models.go               # Data models
│   ├── query.sql.go            # Generated queries (sqlc)
│   └── server/                 # HTTP server
│       ├── routes.go           # API endpoints + logic
│       └── routes_test.go      # Unit tests
├── migrations/                 # Database migrations
├── schema.sql                  # Database schema
├── query.sql                   # SQL query definitions
├── test_data.sql              # Test data
├── docker-compose.yml         # Container orchestration
├── Dockerfile                 # Application image
└── sqlc.yaml                  # sqlc configuration
```

## 🚀 Installation

### Requirements
- Docker & Docker Compose
- Go 1.24+ (optional, for development)

> **💡 New to Docker?** Download it from [docker.com](https://www.docker.com/get-started)

### Step 1: Clone the repository
```bash
git clone <repository-url>
cd WorkRESTAPI
```

### Step 2: Create environment file
```bash
# Copy the example configuration file
cp .env.example .env
```

Edit `.env` (or use defaults):
```env
PORT=1323
APP_ENV=production
DB_HOST=db
DB_USERNAME=postgres
DB_PASSWORD=secretpassword
DB_DATABASE=workrestapi
DB_PORT=5432
DB_SCHEMA=public
```

> **💡 Tip:** You can use the default values above - they work out of the box!

### Step 3: Enable test data (recommended for first run)
In `docker-compose.yml`, uncomment this line:
```yaml
# UNCOMMENT THIS LINE:
- ./test_data.sql:/docker-entrypoint-initdb.d/02-test-data.sql
```

### Step 4: Start the application
```bash
# Start all services
docker-compose up -d

# Wait 30 seconds for startup, then check logs
docker-compose logs -f app
```

### Step 5: Verify it's working
```bash
# Check employees (should return 8 employees)
curl http://localhost:1323/employees

# Check sales (should return 60+ sales)
curl http://localhost:1323/sales
```

**🎉 Success!** Your API is now running at `http://localhost:1323`

## 💡 Usage

### Running the application
```bash
# Start all services
docker-compose up -d

# Check logs
docker-compose logs -f app

# Stop services
docker-compose down
```

### Development
```bash
# Run only the database
docker-compose up -d db

# Run the application locally
go run cmd/api/main.go
```

## 📡 API Endpoints

### 👥 Employees

| Method | Endpoint | Description |
|--------|----------|-------------|
| `GET` | `/employees` | Get all employees |
| `GET` | `/employee?id=1` | Get employee by ID |
| `POST` | `/employee` | Add new employee |
| `PUT` | `/employee/:id` | Update employee |
| `DELETE` | `/employee/:id` | Delete employee |

### 💰 Sales

| Method | Endpoint | Description |
|--------|----------|-------------|
| `GET` | `/sales` | Get all sales |
| `GET` | `/sale?id=1` | Get sale by ID |
| `POST` | `/sale` | Add new sale |
| `PUT` | `/sale/:id` | Update sale |
| `DELETE` | `/sale/:id` | Delete sale |

### 📊 PDF Reports

| Method | Endpoint | Description |
|--------|----------|-------------|
| `GET` | `/employee/:id/report/month?year=2025&month=1` | Monthly PDF report |
| `GET` | `/employee/:id/report/quarter?year=2025&quarter=1` | Quarterly PDF report |

## 🔧 Examples

> **💡 All examples assume the API is running at `http://localhost:1323`**

### Add new employee
```bash
curl -X POST http://localhost:1323/employee \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Anna",
    "surname": "Kowalska",
    "email": "anna.kowalska@company.com"
  }'
```

**Expected response:**
```json
{
  "id": 9,
  "name": "Anna",
  "surname": "Kowalska",
  "email": "anna.kowalska@company.com",
  "created_at": "2025-07-05T12:30:00Z",
  "updated_at": "2025-07-05T12:30:00Z"
}
```

### Add sale (with flexible date format)
```bash
curl -X POST http://localhost:1323/sale \
  -H "Content-Type: application/json" \
  -d '{
    "product_name": "Dell Laptop",
    "category": "Electronics",
    "currency": "PLN",
    "price": 4500.00,
    "sale_date": "15.01.2025 10:30",
    "employee_id": 1
  }'
```

**Expected response:**
```json
{
  "id": 67,
  "product_name": "Dell Laptop",
  "category": "Electronics",
  "currency": "PLN",
  "price": "4500.00",
  "sale_date": "2025-01-15T10:30:00+01:00",
  "employee_id": 1,
  "created_at": "2025-07-05T12:31:00Z",
  "updated_at": "2025-07-05T12:31:00Z"
}
```

### Generate PDF report
```bash
# Monthly report for employee ID=1 for January 2025
curl "http://localhost:1323/employee/1/report/month?year=2025&month=1" \
  --output january_2025_report.pdf

# Quarterly report for employee ID=1 for Q1 2025
curl "http://localhost:1323/employee/1/report/quarter?year=2025&quarter=1" \
  --output q1_2025_report.pdf
```

**What you'll get:**
- PDF files with employee sales statistics
- Tables showing all transactions in the period
- Total sales count and revenue summary

### Get all employees
```bash
curl http://localhost:1323/employees
```

**Sample response:**
```json
[
  {
    "id": 1,
    "name": "Jan",
    "surname": "Kowalski",
    "email": "jan.kowalski@company.com",
    "created_at": "2025-07-05T10:00:00Z",
    "updated_at": "2025-07-05T10:00:00Z"
  },
  ...
]
```

## 🧪 Testing

### Run unit tests
```bash
# All tests
go test ./...

# Tests with coverage
go test -cover ./...

# Server tests
go test ./internal/server/
```

## 📊 Test Data

The project includes ready-to-use test data (`test_data.sql`):

### 👥 8 employees
- Jan Kowalski, Anna Nowak, Piotr Wiśniewski, etc.
- Various Polish company emails

### 💰 60+ sales records
- **Different periods**: January-July 2025 + 2024 data
- **Different categories**: Electronics, Furniture, Software
- **Different currencies**: PLN, EUR, USD
- **Different prices**: from 79.99 to 8999.00

### 📈 Sample data
```sql
-- Employee
('Jan', 'Kowalski', 'jan.kowalski@company.com')

-- Sales
('Dell XPS 13 Laptop', 'Electronics', 'PLN', 4500.00, '2025-01-15 10:30:00+01:00', 1)
('Office 365', 'Software', 'EUR', 299.99, '2025-01-30 12:00:00+01:00', 8)
```

### Loading test data

**Method 1: Automatic (recommended)**
```yaml
# Add to docker-compose.yml:
- ./test_data.sql:/docker-entrypoint-initdb.d/02-test-data.sql
```

**Method 2: Manual**
```bash
# Load data to running database
docker exec -i workrestapi-db-1 psql -U postgres -d workrestapi < test_data.sql
```

## 🐛 Troubleshooting

### Problem: Database connection error
```bash
# Check if database is running
docker-compose ps

# Check database logs
docker-compose logs db

# Restart services
docker-compose restart
```

### Problem: Environment variables
```bash
# Check if .env file exists
ls -la .env

# Check if variables are loaded
docker-compose config
```

### Problem: "Connection refused" when testing API
```bash
# Wait for services to start (can take 30-60 seconds)
docker-compose logs -f app

# Check if containers are running
docker-compose ps

# Try accessing the API
curl http://localhost:1323/employees
```

### Problem: Empty response from API
This usually means test data wasn't loaded:
```bash
# Check if test data line is uncommented in docker-compose.yml
# - ./test_data.sql:/docker-entrypoint-initdb.d/02-test-data.sql

# Or manually load test data
docker exec -i workrestapi-db-1 psql -U postgres -d workrestapi < test_data.sql
```

### Problem: Port already in use
```bash
# Change port in .env file
PORT=8080

# Or stop other services using port 1323
docker-compose down
```

### Still having issues?
1. **Check logs:** `docker-compose logs -f`
2. **Restart everything:** `docker-compose down && docker-compose up -d`
3. **Clean start:** `docker-compose down -v && docker-compose up -d`

## 📝 License

MIT License - you can freely use, modify and distribute.

---

**Author:** Your Name  
**Last Updated:** July 2025  
**Version:** 1.0.0

🌟 **If you find this project helpful, please give it a star!** ⭐
