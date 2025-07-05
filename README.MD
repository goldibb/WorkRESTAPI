# Work REST API 🚀

**Modern REST API for employee and sales management with automatic PDF report generation**

![Go](https://img.shields.io/badge/Go-1.24-00ADD8?style=for-the-badge&logo=go)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-Latest-336791?style=for-the-badge&logo=postgresql)
![Docker](https://img.shields.io/badge/Docker-Compose-2496ED?style=for-the-badge&logo=docker)
![Echo](https://img.shields.io/badge/Echo-v4-00ADD8?style=for-the-badge)

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

Edit `.env`:
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

### Step 3: Run with test data
```bash
# Uncomment the line in docker-compose.yml for automatic data loading
# - ./test_data.sql:/docker-entrypoint-initdb.d/02-test-data.sql

# Start the application
docker-compose up -d
```

### Step 4: Verify it's working
```bash
# Check employees
curl http://localhost:1323/employees

# Check sales
curl http://localhost:1323/sales
```

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

### Generate PDF report
```bash
# Monthly report for employee ID=1 for January 2025
curl "http://localhost:1323/employee/1/report/month?year=2025&month=1" \
  --output january_2025_report.pdf

# Quarterly report for employee ID=1 for Q1 2025
curl "http://localhost:1323/employee/1/report/quarter?year=2025&quarter=1" \
  --output q1_2025_report.pdf
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

## 📝 License

MIT License - you can freely use, modify and distribute.

---
