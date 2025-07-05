-- Test data for WorkRESTAPI
-- Insert employees
INSERT INTO employees (name, surname, email) VALUES 
('Jan', 'Kowalski', 'jan.kowalski@firma.pl'),
('Anna', 'Nowak', 'anna.nowak@firma.pl'),
('Piotr', 'Wiśniewski', 'piotr.wisniewski@firma.pl'),
('Maria', 'Wójcik', 'maria.wojcik@firma.pl'),
('Tomasz', 'Kowalczyk', 'tomasz.kowalczyk@firma.pl'),
('Katarzyna', 'Kamińska', 'katarzyna.kaminska@firma.pl'),
('Michał', 'Lewandowski', 'michal.lewandowski@firma.pl'),
('Magdalena', 'Zielińska', 'magdalena.zielinska@firma.pl');

-- Insert sales with various dates from different months and quarters
INSERT INTO sales (product_name, category, currency, price, sale_date, employee_id) VALUES 
-- January 2025 sales
('Laptop Dell XPS 13', 'Electronics', 'PLN', 4500.00, '2025-01-15 10:30:00+01:00', 1),
('iPhone 15', 'Electronics', 'PLN', 3999.99, '2025-01-20 14:15:00+01:00', 2),
('Samsung Galaxy S24', 'Electronics', 'PLN', 3200.50, '2025-01-25 16:45:00+01:00', 3),
('MacBook Pro', 'Electronics', 'PLN', 8999.00, '2025-01-28 11:20:00+01:00', 1),

-- February 2025 sales
('Monitor 4K LG', 'Electronics', 'PLN', 1299.99, '2025-02-05 09:15:00+01:00', 2),
('Klawiatura mechaniczna', 'Electronics', 'PLN', 450.00, '2025-02-10 13:30:00+01:00', 4),
('Mysz gamingowa', 'Electronics', 'PLN', 199.99, '2025-02-14 15:45:00+01:00', 3),
('Słuchawki Sony', 'Electronics', 'PLN', 799.00, '2025-02-20 12:10:00+01:00', 5),
('Tablet iPad', 'Electronics', 'PLN', 2299.00, '2025-02-25 17:30:00+01:00', 1),

-- March 2025 sales
('Smartwatch Apple', 'Electronics', 'PLN', 1899.00, '2025-03-02 10:45:00+01:00', 6),
('Kamera Canon', 'Electronics', 'PLN', 3500.00, '2025-03-08 14:20:00+01:00', 2),
('Drukarka HP', 'Electronics', 'PLN', 899.99, '2025-03-15 11:55:00+01:00', 7),
('Router WiFi 6', 'Electronics', 'PLN', 299.00, '2025-03-22 16:10:00+01:00', 4),
('Dysk SSD 1TB', 'Electronics', 'PLN', 399.99, '2025-03-28 13:25:00+01:00', 8),

-- April 2025 sales (Q2)
('Laptop Gaming ASUS', 'Electronics', 'PLN', 5999.00, '2025-04-03 09:30:00+02:00', 1),
('Monitor Gaming 144Hz', 'Electronics', 'PLN', 1599.99, '2025-04-10 15:20:00+02:00', 3),
('Karta graficzna RTX', 'Electronics', 'PLN', 2999.00, '2025-04-18 12:40:00+02:00', 5),
('Procesor Intel i7', 'Electronics', 'PLN', 1499.99, '2025-04-25 14:55:00+02:00', 2),

-- May 2025 sales
('Smartphone Xiaomi', 'Electronics', 'PLN', 1299.00, '2025-05-05 11:15:00+02:00', 6),
('Głośnik Bluetooth', 'Electronics', 'PLN', 249.99, '2025-05-12 16:30:00+02:00', 4),
('Powerbank 20000mAh', 'Electronics', 'PLN', 149.00, '2025-05-20 10:20:00+02:00', 7),
('Ładowarka bezprzewodowa', 'Electronics', 'PLN', 99.99, '2025-05-28 13:45:00+02:00', 8),

-- June 2025 sales
('Laptop Lenovo ThinkPad', 'Electronics', 'PLN', 4299.00, '2025-06-02 14:10:00+02:00', 1),
('Mikrofon USB', 'Electronics', 'PLN', 399.00, '2025-06-15 12:25:00+02:00', 3),
('Webcam 4K', 'Electronics', 'PLN', 299.99, '2025-06-22 15:40:00+02:00', 5),
('Podkładka pod mysz RGB', 'Electronics', 'PLN', 79.99, '2025-06-28 11:50:00+02:00', 2),

-- July 2025 sales (current month - few days back)
('Monitor ultrawide', 'Electronics', 'PLN', 1999.00, '2025-07-01 10:00:00+02:00', 6),
('Klawiatura bezprzewodowa', 'Electronics', 'PLN', 199.99, '2025-07-03 14:30:00+02:00', 4),
('Mysz ergonomiczna', 'Electronics', 'PLN', 159.00, '2025-07-04 16:15:00+02:00', 7),

-- Some older sales for testing (2024)
('Laptop HP Pavilion', 'Electronics', 'PLN', 2999.00, '2024-12-15 13:20:00+01:00', 1),
('Tablet Samsung', 'Electronics', 'PLN', 1599.00, '2024-11-20 15:30:00+01:00', 2),
('Smartfon OnePlus', 'Electronics', 'PLN', 2499.00, '2024-10-10 12:45:00+02:00', 3),

-- Different categories for variety
('Krzesło biurowe', 'Furniture', 'PLN', 899.00, '2025-02-12 11:30:00+01:00', 4),
('Biurko regulowane', 'Furniture', 'PLN', 1299.00, '2025-03-18 14:20:00+01:00', 5),
('Lampa LED', 'Furniture', 'PLN', 199.99, '2025-04-22 16:10:00+02:00', 6),
('Organizer na biurko', 'Furniture', 'PLN', 79.99, '2025-05-15 10:45:00+02:00', 7),

-- Different currencies
('Software licencja', 'Software', 'EUR', 299.99, '2025-01-30 12:00:00+01:00', 8),
('Antywirus premium', 'Software', 'USD', 89.99, '2025-03-12 15:30:00+01:00', 1),
('Office 365', 'Software', 'PLN', 449.00, '2025-06-08 11:15:00+02:00', 3);
