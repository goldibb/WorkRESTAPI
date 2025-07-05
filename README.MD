# Work REST API 🚀

**Nowoczesne REST API do zarządzania pracownikami i sprzedażą z automatycznym generowaniem raportów PDF**

![Go](https://img.shields.io/badge/Go-1.24-00ADD8?style=for-the-badge&logo=go)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-Latest-336791?style=for-the-badge&logo=postgresql)
![Docker](https://img.shields.io/badge/Docker-Compose-2496ED?style=for-the-badge&logo=docker)
![Echo](https://img.shields.io/badge/Echo-v4-00ADD8?style=for-the-badge)

## 📋 Spis treści

- [Funkcjonalności](#-funkcjonalności)
- [Technologie](#-technologie)
- [Architektura](#-architektura)
- [Instalacja](#-instalacja)
- [Użycie](#-użycie)
- [API Endpoints](#-api-endpoints)
- [Przykłady](#-przykłady)
- [Testowanie](#-testowanie)
- [Dane testowe](#-dane-testowe)

## ✨ Funkcjonalności

### 👥 Zarządzanie pracownikami (CRUD)
- ✅ Dodawanie, edycja, usuwanie pracowników
- ✅ Walidacja danych (email, długość pól)
- ✅ Sprawdzanie unikalności email
- ✅ Automatyczne timestamps (created_at, updated_at)

### 💰 Zarządzanie sprzedażą (CRUD)
- ✅ Dodawanie, edycja, usuwanie sprzedaży
- ✅ Obsługa różnych walut (PLN, EUR, USD)
- ✅ Walidacja cen (muszą być > 0)
- ✅ **Elastyczne formaty dat** - obsługa wielu formatów:
  - ISO 8601: `2025-01-15T10:30:00Z`
  - RFC 3339: `2025-01-15T10:30:00+01:00`
  - Polski format: `15.01.2025 10:30`
  - Europejski: `15-01-2025 10:30`
  - Amerykański: `01/15/2025 10:30`

### 📊 Raporty PDF
- ✅ **Miesięczne raporty** pracowników
- ✅ **Kwartalne raporty** pracowników
- ✅ Automatyczne generowanie PDF z gofpdf
- ✅ Statystyki: liczba sprzedaży, suma przychodów
- ✅ Szczegółowe tabele wszystkich transakcji

### 🔒 Bezpieczeństwo i walidacja
- ✅ Walidacja email (regex)
- ✅ Walidacja długości pól
- ✅ Sprawdzanie unikalności email
- ✅ Walidacja dat i przedziałów czasowych
- ✅ Walidacja cen i walut

## 🛠 Technologie

| Komponent | Technologia | Wersja |
|-----------|-------------|---------|
| **Backend** | Go | 1.24 |
| **Framework** | Echo | v4 |
| **Baza danych** | PostgreSQL | Latest |
| **ORM/Query Builder** | sqlc | Latest |
| **Połączenie DB** | pgx/v5 | v5.7.5 |
| **PDF Generator** | gofpdf | v1.4.3 |
| **Konteneryzacja** | Docker + Docker Compose | Latest |

## 🏗 Architektura

```
WorkRESTAPI/
├── cmd/api/                    # Punkt wejścia aplikacji
│   └── main.go
├── internal/                   # Logika biznesowa
│   ├── db.go                   # Połączenie z bazą
│   ├── models.go               # Modele danych
│   ├── query.sql.go            # Wygenerowane zapytania (sqlc)
│   └── server/                 # Serwer HTTP
│       ├── routes.go           # Endpointy API + logika
│       └── routes_test.go      # Testy jednostkowe
├── migrations/                 # Migracje bazy danych
├── schema.sql                  # Schemat bazy danych
├── query.sql                   # Definicje zapytań SQL
├── test_data.sql              # Dane testowe
├── docker-compose.yml         # Orkiestracja kontenerów
├── Dockerfile                 # Obraz aplikacji
└── sqlc.yaml                  # Konfiguracja sqlc
```

## 🚀 Instalacja

### Wymagania
- Docker & Docker Compose
- Go 1.24+ (opcjonalnie, do developmentu)

### Krok 1: Sklonuj repozytorium
```bash
git clone <repository-url>
cd WorkRESTAPI
```

### Krok 2: Utwórz plik środowiskowy
```bash
# Skopiuj przykładowy plik konfiguracyjny
cp .env.example .env
```

Edytuj `.env`:
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

### Krok 3: Uruchom z danymi testowymi
```bash
# Odkomentuj linię w docker-compose.yml dla automatycznego ładowania danych
# - ./test_data.sql:/docker-entrypoint-initdb.d/02-test-data.sql

# Uruchom aplikację
docker-compose up -d
```

### Krok 4: Sprawdź czy działa
```bash
# Sprawdź pracowników
curl http://localhost:1323/employees

# Sprawdź sprzedaże
curl http://localhost:1323/sales
```

## 💡 Użycie

### Uruchomienie
```bash
# Uruchom wszystkie serwisy
docker-compose up -d

# Sprawdź logi
docker-compose logs -f app

# Zatrzymaj serwisy
docker-compose down
```

### Development
```bash
# Uruchom tylko bazę danych
docker-compose up -d db

# Uruchom aplikację lokalnie
go run cmd/api/main.go
```

## 📡 API Endpoints

### 👥 Pracownicy

| Metoda | Endpoint | Opis |
|--------|----------|------|
| `GET` | `/employees` | Pobierz wszystkich pracowników |
| `GET` | `/employee?id=1` | Pobierz pracownika po ID |
| `POST` | `/employee` | Dodaj nowego pracownika |
| `PUT` | `/employee/:id` | Zaktualizuj pracownika |
| `DELETE` | `/employee/:id` | Usuń pracownika |

### 💰 Sprzedaże

| Metoda | Endpoint | Opis |
|--------|----------|------|
| `GET` | `/sales` | Pobierz wszystkie sprzedaże |
| `GET` | `/sale?id=1` | Pobierz sprzedaż po ID |
| `POST` | `/sale` | Dodaj nową sprzedaż |
| `PUT` | `/sale/:id` | Zaktualizuj sprzedaż |
| `DELETE` | `/sale/:id` | Usuń sprzedaż |

### 📊 Raporty PDF

| Metoda | Endpoint | Opis |
|--------|----------|------|
| `GET` | `/employee/:id/report/month?year=2025&month=1` | Raport miesięczny PDF |
| `GET` | `/employee/:id/report/quarter?year=2025&quarter=1` | Raport kwartalny PDF |

## 🔧 Przykłady

### Dodaj nowego pracownika
```bash
curl -X POST http://localhost:1323/employee \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Anna",
    "surname": "Kowalska",
    "email": "anna.kowalska@firma.pl"
  }'
```

### Dodaj sprzedaż (z elastycznym formatem daty)
```bash
curl -X POST http://localhost:1323/sale \
  -H "Content-Type: application/json" \
  -d '{
    "product_name": "Laptop Dell",
    "category": "Electronics",
    "currency": "PLN",
    "price": 4500.00,
    "sale_date": "15.01.2025 10:30",
    "employee_id": 1
  }'
```

### Wygeneruj raport PDF
```bash
# Miesięczny raport pracownika ID=1 za styczeń 2025
curl "http://localhost:1323/employee/1/report/month?year=2025&month=1" \
  --output raport_styczen_2025.pdf

# Kwartalny raport pracownika ID=1 za Q1 2025
curl "http://localhost:1323/employee/1/report/quarter?year=2025&quarter=1" \
  --output raport_q1_2025.pdf
```

## 🧪 Testowanie

### Uruchom testy jednostkowe
```bash
# Wszystkie testy
go test ./...

# Testy z pokryciem
go test -cover ./...

# Testy serwera
go test ./internal/server/
```

### Testy integracyjne
```bash
# Testowanie API (PowerShell)
.\test_api.ps1

# Testowanie API (Bash)
./test_api.sh
```

## 📊 Dane testowe

Projekt zawiera gotowe dane testowe (`test_data.sql`):

### 👥 8 pracowników
- Jan Kowalski, Anna Nowak, Piotr Wiśniewski, etc.
- Różne emaile firm polskich

### 💰 60+ sprzedaży
- **Różne okresy**: styczeń-lipiec 2025 + dane z 2024
- **Różne kategorie**: Electronics, Furniture, Software
- **Różne waluty**: PLN, EUR, USD
- **Różne ceny**: od 79.99 do 8999.00

### 📈 Przykładowe dane
```sql
-- Pracownik
('Jan', 'Kowalski', 'jan.kowalski@firma.pl')

-- Sprzedaże
('Laptop Dell XPS 13', 'Electronics', 'PLN', 4500.00, '2025-01-15 10:30:00+01:00', 1)
('Office 365', 'Software', 'EUR', 299.99, '2025-01-30 12:00:00+01:00', 8)
```

### Ładowanie danych testowych

**Sposób 1: Automatyczne (zalecany)**
```yaml
# W docker-compose.yml dodaj:
- ./test_data.sql:/docker-entrypoint-initdb.d/02-test-data.sql
```

**Sposób 2: Ręczne**
```bash
# Załaduj dane do działającej bazy
docker exec -i workrestapi-db-1 psql -U postgres -d workrestapi < test_data.sql
```

## 🔍 Funkcje zaawansowane

### Elastyczne daty
API automatycznie rozpoznaje i parsuje różne formaty dat:
- `"2025-01-15T10:30:00Z"` (ISO 8601)
- `"15.01.2025 10:30"` (Polski)
- `"15-01-2025 10:30"` (Europejski)
- `"01/15/2025 10:30"` (Amerykański)

### Walidacje
- ✅ Email: regex pattern validation
- ✅ Ceny: muszą być > 0
- ✅ Daty: sprawdzanie poprawności
- ✅ Unikalność: email nie może się powtarzać

### Raporty PDF
- 📊 Statystyki: liczba sprzedaży, suma przychodów
- 📋 Tabele: szczegółowe dane wszystkich transakcji
- 📅 Okresy: miesięczne i kwartalne
- 💾 Auto-download: pliki PDF z nazwami pracowników

## 🐛 Rozwiązywanie problemów

### Problem: Błąd połączenia z bazą
```bash
# Sprawdź czy baza jest uruchomiona
docker-compose ps

# Sprawdź logi bazy
docker-compose logs db

# Restartuj serwisy
docker-compose restart
```

### Problem: Zmienne środowiskowe
```bash
# Sprawdź czy plik .env istnieje
ls -la .env

# Sprawdź czy zmienne są załadowane
docker-compose config
```

## 📝 Licencja

MIT License - możesz swobodnie używać, modyfikować i dystrybuować.

---

**Autor**: Twój Nick  
**Ostatnia aktualizacja**: Lipiec 2025  
**Wersja**: 1.0.0

🌟 **Jeśli projekt Ci się podoba, zostaw gwiazdkę!** ⭐
