INSERT INTO users (username, email, password, role) VALUES
('johndoe', 'john@email.com', 'johndoe123', 'user'),
('janedoe', 'jane@email.com', 'janedoe123', 'user'),
('adminuser', 'admin@cinemax.com', 'admin123', 'admin'),
('mikebrown', 'mike@email.com', 'mikebrown123', 'user'),
('sarahjones', 'sarah@email.com', '$sarahjones123', 'user');

SELECT * FROM users;
SELECT * FROM users WHERE email = 'denis@gmail.com'

SELECT * FROM movies_genres LIMIT 100;
INSERT INTO genres (name) VALUES
('Action'),
('Comedy'),
('Drama'),
('Horror'),
('Sci-Fi'),
('Romance'),
('Thriller'),
('Adventure'),
('Fantasy'),
('Animation'),
('Documentary'),
('Mystery'),
('Crime'),
('Family'),
('Musical'),
('Western'),
('War'),
('History');

SELECT id FROM genres WHERE name = 'War';

SELECT * FROM cities LIMIT 100;
INSERT INTO cities (name, province) VALUES
('Jakarta', 'DKI Jakarta'),
('Bandung', 'Jawa Barat'),
('Surabaya', 'Jawa Timur'),
('Medan', 'Sumatera Utara'),
('Yogyakarta', 'DI Yogyakarta');


SELECT * FROM cinemas LIMIT 100;
INSERT INTO cinemas (city_id, name, address, is_active) VALUES
(1, 'Cinema XXI Grand Indonesia', 'Jl. MH Thamrin No.1, Jakarta Pusat', true),
(1, 'CGV Plaza Senayan', 'Jl. Asia Afrika No.8, Jakarta Selatan', true),
(2, 'Cinema 21 Bandung Indah Plaza', 'Jl. Merdeka No.56, Bandung', true),
(3, 'Cinepolis Pakuwon Mall', 'Jl. Puncak Indah Lontar No.2, Surabaya', true),
(4, 'XXI Medan Fair', 'Jl. Gagak Hitam No.8, Medan', false);

SELECT * FROM theaters LIMIT 100;
INSERT INTO theaters (cinema_id, name, capacity) VALUES
(1, 'Theater 1', 120),
(1, 'Theater 2', 100),
(2, 'Theater A', 150),
(2, 'Theater B', 80),
(3, 'Studio 1', 110),
(3, 'Studio 2', 90),
(4, 'Hall 1', 130);


SELECT * FROM seats;
INSERT INTO seats (theater_id, seat_letter, seat_number, is_active) VALUES
(1, 'A', 1, true),
(1, 'A', 2, true),
(1, 'A', 3, true),
(1, 'B', 1, true),
(1, 'B', 2, true),
(1, 'B', 3, false), -- kursi rusak
(1, 'C', 1, true),
(1, 'C', 2, true),
(1, 'C', 3, true),
(1, 'D', 1, true);

SELECT * FROM showtimes;
INSERT INTO showtimes (movie_id, theater_id, show_date, show_time, base_price, available_seats) VALUES
(19, 1, '2024-07-08', '2024-07-08 14:00:00', 45000.00, 115),
(11, 1, '2024-07-08', '2024-07-08 19:00:00', 55000.00, 118),
(20, 2, '2024-07-08', '2024-07-08 16:30:00', 50000.00, 98),
(12, 3, '2024-07-09', '2024-07-09 13:00:00', 40000.00, 148),
(16, 4, '2024-07-09', '2024-07-09 20:00:00', 45000.00, 78);

SELECT * FROM payments;
INSERT INTO payments (method_name, provider, fee_process, is_active) VALUES
('Credit Card', 'Visa', 2500.00, true),
('Debit Card', 'BCA', 1500.00, true),
('E-Wallet', 'GoPay', 1000.00, true),
('E-Wallet', 'OVO', 1200.00, true),
('Bank Transfer', 'Mandiri', 5000.00, true);

SELECT * FROM bookings;
INSERT INTO bookings (user_id, showtime_id, status, total_amount, discount_amount, tax_amount, booking_time, expires_at) VALUES
('a025f497-60bf-42b1-be3a-4740a6d37b1d', 141, 'confirmed', 90000.00, 0.00, 9000.00, '2024-07-08 12:30:00', '2024-07-08 13:30:00'),
('a025f497-60bf-42b1-be3a-4740a6d37b1d', 142, 'confirmed', 110000.00, 5000.00, 11000.00, '2024-07-08 17:00:00', '2024-07-08 18:00:00'),
('a025f497-60bf-42b1-be3a-4740a6d37b1d', 143, 'pending', 100000.00, 0.00, 10000.00, '2024-07-08 14:45:00', '2024-07-08 15:45:00'),
('4a3cec21-2daa-4d1d-ba4b-7c398ecc9a16', 141, 'cancelled', 45000.00, 0.00, 4500.00, '2024-07-08 10:00:00', '2024-07-08 11:00:00'),
('4a3cec21-2daa-4d1d-ba4b-7c398ecc9a16', 144, 'confirmed', 80000.00, 10000.00, 8000.00, '2024-07-09 11:30:00', '2024-07-09 12:30:00');

INSERT INTO booking_seats (booking_id, seat_id, seat_price) VALUES
(11, 1, 45000.00),
(11, 2, 45000.00),
(12, 1, 55000.00),
(12, 4, 55000.00),
(13, 1, 50000.00),
(13, 7, 50000.00),
(15, 1, 40000.00),
(15, 8, 40000.00);


SELECT * FROM payments_transactions;
INSERT INTO payments_transactions (payment_id, booking_id, status, amount, completed_at) VALUES
(1, 11, 'success', 90000.00, '2024-07-08 12:35:00'),
(2, 13, 'success', 110000.00, '2024-07-08 17:05:00'),
(3, 12, 'pending', 100000.00, NULL),
(4, 11, 'refunded', 45000.00, '2024-07-08 10:30:00'),
(5, 14, 'success', 80000.00, '2024-07-09 11:35:00');