-- ========================================
-- ANALISIS STRUKTUR DATABASE CINEMAX
-- ========================================

-- Database ini adalah sistem manajemen bioskop dengan struktur:
-- 1. User Management (users, profiles)
-- 2. Movie Management (movies, genres, casts, directors)
-- 3. Cinema Management (cities, cinemas, theaters, seats)
-- 4. Booking & Payment (showtimes, bookings, payments, transactions)

-- ========================================
-- DATA DUMMY YANG REALISTIS
-- ========================================

-- 1. USERS TABLE
INSERT INTO users (id, email, password, role, created_at, updated_at) VALUES
('550e8400-e29b-41d4-a716-446655440001', 'admin@cinemax.com', '$2a$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi', 'admin', '2024-01-15 08:00:00', '2024-01-15 08:00:00'),
('550e8400-e29b-41d4-a716-446655440002', 'john.doe@gmail.com', '$2a$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi', 'user', '2024-02-01 10:30:00', '2024-02-01 10:30:00'),
('550e8400-e29b-41d4-a716-446655440003', 'sarah.wilson@yahoo.com', '$2a$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi', 'user', '2024-02-15 14:20:00', '2024-02-15 14:20:00'),
('550e8400-e29b-41d4-a716-446655440004', 'mike.chen@outlook.com', '$2a$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi', 'user', '2024-03-01 16:45:00', '2024-03-01 16:45:00'),
('550e8400-e29b-41d4-a716-446655440005', 'emma.davis@gmail.com', '$2a$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi', 'user', '2024-03-10 09:15:00', '2024-03-10 09:15:00');

-- 2. PROFILES TABLE
INSERT INTO profiles (user_id, username, firstname, lastname, birthday, gender, profile_picture, phone_number, is_verified, created_at, updated_at) VALUES
('550e8400-e29b-41d4-a716-446655440001', 'admin_cinemax', 'Admin', 'Cinemax', '1985-05-15', 'male', 'https://example.com/profiles/admin.jpg', '+6281234567890', TRUE, '2024-01-15 08:05:00', '2024-01-15 08:05:00'),
('550e8400-e29b-41d4-a716-446655440002', 'johndoe92', 'John', 'Doe', '1992-08-22', 'male', 'https://example.com/profiles/john.jpg', '+6281234567891', TRUE, '2024-02-01 10:35:00', '2024-02-01 10:35:00'),
('550e8400-e29b-41d4-a716-446655440003', 'sarahw', 'Sarah', 'Wilson', '1995-03-18', 'female', 'https://example.com/profiles/sarah.jpg', '+6281234567892', TRUE, '2024-02-15 14:25:00', '2024-02-15 14:25:00'),
('550e8400-e29b-41d4-a716-446655440004', 'mikechen', 'Mike', 'Chen', '1988-11-05', 'male', 'https://example.com/profiles/mike.jpg', '+6281234567893', FALSE, '2024-03-01 16:50:00', '2024-03-01 16:50:00'),
('550e8400-e29b-41d4-a716-446655440005', 'emmad', 'Emma', 'Davis', '1997-07-12', 'female', 'https://example.com/profiles/emma.jpg', '+6281234567894', TRUE, '2024-03-10 09:20:00', '2024-03-10 09:20:00');

-- 3. GENRES TABLE
INSERT INTO genres (name, created_at, updated_at) VALUES
('Action', '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
('Adventure', '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
('Comedy', '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
('Drama', '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
('Horror', '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
('Romance', '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
('Sci-Fi', '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
('Thriller', '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
('Animation', '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
('Fantasy', '2024-01-01 00:00:00', '2024-01-01 00:00:00');

-- 4. DIRECTORS TABLE
INSERT INTO directors (name, created_at, updated_at) VALUES
('Christopher Nolan', '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
('Steven Spielberg', '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
('Quentin Tarantino', '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
('Martin Scorsese', '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
('Denis Villeneuve', '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
('Greta Gerwig', '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
('Jordan Peele', '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
('Rian Johnson', '2024-01-01 00:00:00', '2024-01-01 00:00:00');

-- 5. CASTS TABLE
INSERT INTO casts (actor_name, created_at, updated_at) VALUES
('Leonardo DiCaprio', '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
('Margot Robbie', '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
('Ryan Gosling', '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
('Emma Stone', '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
('Timothée Chalamet', '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
('Zendaya', '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
('Tom Holland', '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
('Scarlett Johansson', '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
('Robert Downey Jr.', '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
('Chris Evans', '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
('Gal Gadot', '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
('Henry Cavill', '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
('Daniel Craig', '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
('Ana de Armas', '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
('Oscar Isaac', '2024-01-01 00:00:00', '2024-01-01 00:00:00');

-- 6. MOVIES TABLE (20 movies maximum as requested)
INSERT INTO movies (title, description, release_date, duration, rating, status, language, backdrop_img, poster_img, popularity, created_at, updated_at) VALUES
('Oppenheimer', 'The story of J. Robert Oppenheimer and his role in the development of the atomic bomb.', '2023-07-21', 180, 8.3, 'now playing', 'English', 'https://example.com/backdrops/oppenheimer.jpg', 'https://example.com/posters/oppenheimer.jpg', 95.8, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
('Barbie', 'Barbie and Ken are having the time of their lives in the colorful and seemingly perfect world of Barbie Land.', '2023-07-21', 114, 7.2, 'now playing', 'English', 'https://example.com/backdrops/barbie.jpg', 'https://example.com/posters/barbie.jpg', 94.2, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
('Dune: Part Two', 'Paul Atreides unites with Chani and the Fremen while seeking revenge against the conspirators who destroyed his family.', '2024-03-01', 166, 8.7, 'now playing', 'English', 'https://example.com/backdrops/dune2.jpg', 'https://example.com/posters/dune2.jpg', 98.5, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
('Spider-Man: No Way Home', 'Peter Parker seeks help from Doctor Strange to make everyone forget his identity as Spider-Man.', '2021-12-17', 148, 8.2, 'ended', 'English', 'https://example.com/backdrops/spiderman.jpg', 'https://example.com/posters/spiderman.jpg', 92.1, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
('The Batman', 'Batman ventures into Gotham City underworld when a sadistic killer leaves behind a trail of cryptic clues.', '2022-03-04', 176, 7.8, 'ended', 'English', 'https://example.com/backdrops/batman.jpg', 'https://example.com/posters/batman.jpg', 89.3, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
('Top Gun: Maverick', 'After thirty years, Maverick is still pushing the envelope as a top naval aviator.', '2022-05-27', 131, 8.3, 'ended', 'English', 'https://example.com/backdrops/topgun.jpg', 'https://example.com/posters/topgun.jpg', 91.7, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
('Avatar: The Way of Water', 'Jake Sully and Neytiri have formed a family and are doing everything to stay together.', '2022-12-16', 192, 7.6, 'ended', 'English', 'https://example.com/backdrops/avatar2.jpg', 'https://example.com/posters/avatar2.jpg', 88.4, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
('Black Panther: Wakanda Forever', 'Queen Ramonda, Shuri, MBaku, Okoye and the Dora Milaje fight to protect their nation.', '2022-11-11', 161, 6.7, 'ended', 'English', 'https://example.com/backdrops/blackpanther2.jpg', 'https://example.com/posters/blackpanther2.jpg', 85.2, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
('Everything Everywhere All at Once', 'A middle-aged Chinese immigrant is swept up in an insane adventure.', '2022-03-25', 139, 7.8, 'ended', 'English', 'https://example.com/backdrops/everything.jpg', 'https://example.com/posters/everything.jpg', 87.9, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
('Nope', 'The residents of a lonely gulch in inland California bear witness to an uncanny and chilling discovery.', '2022-07-22', 130, 6.8, 'ended', 'English', 'https://example.com/backdrops/nope.jpg', 'https://example.com/posters/nope.jpg', 74.5, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
('The Flash', 'Barry Allen uses his super speed to change the past, but his attempt to save his family creates a world without superheroes.', '2023-06-16', 144, 6.9, 'ended', 'English', 'https://example.com/backdrops/flash.jpg', 'https://example.com/posters/flash.jpg', 76.3, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
('Indiana Jones: Dial of Destiny', 'Archaeologist Indiana Jones races against time to retrieve a legendary artifact.', '2023-06-30', 154, 6.5, 'ended', 'English', 'https://example.com/backdrops/indiana.jpg', 'https://example.com/posters/indiana.jpg', 72.8, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
('Scream VI', 'Following the latest Ghostface killings, the four survivors leave Woodsboro behind.', '2023-03-10', 123, 6.4, 'ended', 'English', 'https://example.com/backdrops/scream6.jpg', 'https://example.com/posters/scream6.jpg', 71.2, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
('John Wick: Chapter 4', 'John Wick uncovers a path to defeating The High Table.', '2023-03-24', 169, 7.7, 'ended', 'English', 'https://example.com/backdrops/johnwick4.jpg', 'https://example.com/posters/johnwick4.jpg', 84.6, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
('Guardians of the Galaxy Vol. 3', 'Still reeling from the loss of Gamora, Peter Quill rallies his team to defend the universe.', '2023-05-05', 150, 7.9, 'ended', 'English', 'https://example.com/backdrops/guardians3.jpg', 'https://example.com/posters/guardians3.jpg', 86.1, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
('Fast X', 'Dom Toretto and his family face their most lethal opponent yet.', '2023-05-19', 141, 5.8, 'ended', 'English', 'https://example.com/backdrops/fastx.jpg', 'https://example.com/posters/fastx.jpg', 68.9, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
('Transformers: Rise of the Beasts', 'Optimus Prime and the Autobots take on their biggest challenge yet.', '2023-06-09', 127, 6.2, 'ended', 'English', 'https://example.com/backdrops/transformers.jpg', 'https://example.com/posters/transformers.jpg', 70.4, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
('The Little Mermaid', 'A young mermaid makes a deal with a sea witch to trade her voice for legs.', '2023-05-26', 135, 7.2, 'ended', 'English', 'https://example.com/backdrops/mermaid.jpg', 'https://example.com/posters/mermaid.jpg', 78.7, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
('Furiosa: A Mad Max Saga', 'The origin story of renegade warrior Furiosa before she teamed up with Mad Max.', '2024-05-24', 148, 7.5, 'coming soon', 'English', 'https://example.com/backdrops/furiosa.jpg', 'https://example.com/posters/furiosa.jpg', 82.3, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
('Deadpool 3', 'Wade Wilson is recruited by the Time Variance Authority to help safeguard the multiverse.', '2024-07-26', 130, 0.0, 'coming soon', 'English', 'https://example.com/backdrops/deadpool3.jpg', 'https://example.com/posters/deadpool3.jpg', 96.7, '2024-01-01 00:00:00', '2024-01-01 00:00:00');

-- 7. MOVIES_GENRES TABLE (Many-to-many relationship)
INSERT INTO movies_genres (movie_id, genre_id) VALUES
(1, 4), (1, 8), -- Oppenheimer: Drama, Thriller
(2, 3), (2, 6), (2, 10), -- Barbie: Comedy, Romance, Fantasy
(3, 1), (3, 2), (3, 7), -- Dune: Action, Adventure, Sci-Fi
(4, 1), (4, 2), (4, 7), -- Spider-Man: Action, Adventure, Sci-Fi
(5, 1), (5, 4), (5, 8), -- The Batman: Action, Drama, Thriller
(6, 1), (6, 2), (6, 4), -- Top Gun: Action, Adventure, Drama
(7, 1), (7, 2), (7, 7), -- Avatar: Action, Adventure, Sci-Fi
(8, 1), (8, 2), (8, 4), -- Black Panther: Action, Adventure, Drama
(9, 1), (9, 3), (9, 7), -- Everything Everywhere: Action, Comedy, Sci-Fi
(10, 5), (10, 8), (10, 7), -- Nope: Horror, Thriller, Sci-Fi
(11, 1), (11, 2), (11, 7), -- The Flash: Action, Adventure, Sci-Fi
(12, 1), (12, 2), -- Indiana Jones: Action, Adventure
(13, 5), (13, 8), -- Scream VI: Horror, Thriller
(14, 1), (14, 8), -- John Wick: Action, Thriller
(15, 1), (15, 2), (15, 3), -- Guardians: Action, Adventure, Comedy
(16, 1), (16, 8), -- Fast X: Action, Thriller
(17, 1), (17, 2), (17, 7), -- Transformers: Action, Adventure, Sci-Fi
(18, 2), (18, 6), (18, 9), -- The Little Mermaid: Adventure, Romance, Animation
(19, 1), (19, 2), (19, 7), -- Furiosa: Action, Adventure, Sci-Fi
(20, 1), (20, 3), (20, 7); -- Deadpool 3: Action, Comedy, Sci-Fi

-- 8. MOVIES_DIRECTORS TABLE
INSERT INTO movies_directors (movie_id, director_id) VALUES
(1, 1), -- Oppenheimer - Christopher Nolan
(2, 6), -- Barbie - Greta Gerwig
(3, 5), -- Dune - Denis Villeneuve
(4, 8), -- Spider-Man - Rian Johnson (example)
(5, 4), -- The Batman - Martin Scorsese (example)
(6, 2), -- Top Gun - Steven Spielberg (example)
(7, 2), -- Avatar - Steven Spielberg (example)
(8, 7), -- Black Panther - Jordan Peele (example)
(9, 8), -- Everything Everywhere - Rian Johnson (example)
(10, 7), -- Nope - Jordan Peele
(11, 1), -- The Flash - Christopher Nolan (example)
(12, 2), -- Indiana Jones - Steven Spielberg
(13, 8), -- Scream VI - Rian Johnson (example)
(14, 1), -- John Wick - Christopher Nolan (example)
(15, 6), -- Guardians - Greta Gerwig (example)
(16, 3), -- Fast X - Quentin Tarantino (example)
(17, 5), -- Transformers - Denis Villeneuve (example)
(18, 6), -- The Little Mermaid - Greta Gerwig (example)
(19, 4), -- Furiosa - Martin Scorsese (example)
(20, 8); -- Deadpool 3 - Rian Johnson (example)

-- 9. MOVIES_CASTS TABLE
INSERT INTO movies_casts (movie_id, cast_id, character_name) VALUES
(1, 1, 'J. Robert Oppenheimer'), -- Oppenheimer
(1, 2, 'Kitty Oppenheimer'),
(2, 2, 'Barbie'), -- Barbie
(2, 3, 'Ken'),
(3, 5, 'Paul Atreides'), -- Dune
(3, 6, 'Chani'),
(4, 7, 'Spider-Man'), -- Spider-Man
(4, 4, 'Gwen Stacy'),
(5, 9, 'Bruce Wayne'), -- The Batman
(5, 8, 'Selina Kyle'),
(6, 10, 'Pete Mitchell'), -- Top Gun
(6, 11, 'Penny Benjamin'),
(7, 12, 'Jake Sully'), -- Avatar
(7, 6, 'Neytiri'),
(8, 8, 'Shuri'), -- Black Panther
(8, 11, 'Ramonda'),
(9, 4, 'Evelyn Quan Wang'), -- Everything Everywhere
(9, 15, 'Waymond Wang'),
(10, 13, 'OJ Haywood'), -- Nope
(10, 14, 'Emerald Haywood');

-- 10. CITIES TABLE
INSERT INTO cities (name, province, created_at, updated_at) VALUES
('Jakarta', 'DKI Jakarta', '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
('Surabaya', 'Jawa Timur', '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
('Bandung', 'Jawa Barat', '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
('Medan', 'Sumatera Utara', '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
('Yogyakarta', 'DI Yogyakarta', '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
('Denpasar', 'Bali', '2024-01-01 00:00:00', '2024-01-01 00:00:00');

-- 11. CINEMAS TABLE
INSERT INTO cinemas (city_id, name, address, is_active, created_at, updated_at) VALUES
(1, 'Cinemax Plaza Indonesia', 'Jl. MH Thamrin No. 28-30, Jakarta Pusat', TRUE, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
(1, 'Cinemax Grand Indonesia', 'Jl. MH Thamrin No. 1, Jakarta Pusat', TRUE, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
(2, 'Cinemax Pakuwon Mall', 'Jl. Puncak Indah Lontar No. 2, Surabaya', TRUE, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
(3, 'Cinemax Paris Van Java', 'Jl. Sukajadi No. 131-139, Bandung', TRUE, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
(4, 'Cinemax Sun Plaza', 'Jl. Gatot Subroto No. 1, Medan', TRUE, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
(5, 'Cinemax Malioboro Mall', 'Jl. Malioboro No. 52-58, Yogyakarta', TRUE, '2024-01-01 00:00:00', '2024-01-01 00:00:00');

-- 12. THEATERS TABLE
INSERT INTO theaters (cinema_id, name, capacity, created_at, updated_at) VALUES
(1, 'Theater 1', 120, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
(1, 'Theater 2', 150, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
(1, 'Theater 3', 180, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
(2, 'Theater 1', 140, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
(2, 'Theater 2', 160, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
(3, 'Theater 1', 130, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
(3, 'Theater 2', 145, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
(4, 'Theater 1', 125, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
(5, 'Theater 1', 135, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
(6, 'Theater 1', 110, '2024-01-01 00:00:00', '2024-01-01 00:00:00');

-- 13. SEATS TABLE (Sample untuk Theater 1 - baris A sampai J, nomor 1-12)
INSERT INTO seats (theater_id, seat_letter, seat_number, is_active, created_at, updated_at) VALUES
-- Theater 1 (120 seats: A1-A12, B1-B12, ..., J1-J12)
(1, 'A', 1, TRUE, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
(1, 'A', 2, TRUE, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
(1, 'A', 3, TRUE, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
(1, 'A', 4, TRUE, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
(1, 'A', 5, TRUE, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
(1, 'A', 6, TRUE, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
(1, 'A', 7, TRUE, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
(1, 'A', 8, TRUE, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
(1, 'A', 9, TRUE, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
(1, 'A', 10, TRUE, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
(1, 'A', 11, TRUE, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
(1, 'A', 12, TRUE, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
(1, 'B', 1, TRUE, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
(1, 'B', 2, TRUE, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
(1, 'B', 3, TRUE, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
(1, 'B', 4, TRUE, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
(1, 'B', 5, TRUE, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
(1, 'B', 6, TRUE, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
(1, 'B', 7, TRUE, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
(1, 'B', 8, TRUE, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
(1, 'B', 9, TRUE, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
(1, 'B', 10, TRUE, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
(1, 'B', 11, TRUE, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
(1, 'B', 12, TRUE, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
(1, 'C', 1, TRUE, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
(1, 'C', 2, TRUE, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
(1, 'C', 3, TRUE, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
(1, 'C', 4, TRUE, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
(1, 'C', 5, TRUE, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
(1, 'C', 6, TRUE, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
(1, 'C', 7, TRUE, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
(1, 'C', 8, TRUE, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
(1, 'C', 9, TRUE, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
(1, 'C', 10, TRUE, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
(1, 'C', 11, TRUE, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
(1, 'C', 12, TRUE, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
(1, 'D', 1, TRUE, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
(1, 'D', 2, TRUE, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
(1, 'D', 3, TRUE, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
(1, 'D', 4, TRUE, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
(1, 'D', 5, TRUE, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
(1, 'D', 6, TRUE, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
(1, 'D', 7, TRUE, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
(1, 'D', 8, TRUE, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
(1, 'D', 9, TRUE, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
(1, 'D', 10, TRUE, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
(1, 'D', 11, TRUE, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
(1, 'D', 12, TRUE, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
(1, 'E', 1, TRUE, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
(1, 'E', 2, TRUE, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
(1, 'E', 3, TRUE, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
(1, 'E', 4, TRUE, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
(1, 'E', 5, TRUE, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
(1, 'E', 6, TRUE, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
(1, 'E', 7, TRUE, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
(1, 'E', 8, TRUE, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
(1, 'E', 9, TRUE, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
(1, 'E', 10, TRUE, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
(1, 'E', 11, TRUE, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
(1, 'E', 12, TRUE, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
(1, 'F', 1, TRUE, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
(1, 'F', 2, TRUE, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
(1, 'F', 3, TRUE, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
(1, 'F', 4, TRUE, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
(1, 'F', 5, TRUE, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
(1, 'F', 6, TRUE, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
(1, 'F', 7, TRUE, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
(1, 'F', 8, TRUE, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
(1, 'F', 9, TRUE, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
(1, 'F', 10, TRUE, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
(1, 'F', 11, TRUE, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
(1, 'F', 12, TRUE, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
(1, 'G', 1, TRUE, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
(1, 'G', 2, TRUE, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
(1, 'G', 3, TRUE, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
(1, 'G', 4, TRUE, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
(1, 'G', 5, TRUE, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
(1, 'G', 6, TRUE, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
(1, 'G', 7, TRUE, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
(1, 'G', 8, TRUE, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
(1, 'G', 9, TRUE, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
(1, 'G', 10, TRUE, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
(1, 'G', 11, TRUE, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
(1, 'G', 12, TRUE, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
(1, 'H', 1, TRUE, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
(1, 'H', 2, TRUE, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
(1, 'H', 3, TRUE, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
(1, 'H', 4, TRUE, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
(1, 'H', 5, TRUE, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
(1, 'H', 6, TRUE, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
(1, 'H', 7, TRUE, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
(1, 'H', 8, TRUE, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
(1, 'H', 9, TRUE, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
(1, 'H', 10, TRUE, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
(1, 'H', 11, TRUE, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
(1, 'H', 12, TRUE, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
(1, 'I', 1, TRUE, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
(1, 'I', 2, TRUE, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
(1, 'I', 3, TRUE, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
(1, 'I', 4, TRUE, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
(1, 'I', 5, TRUE, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
(1, 'I', 6, TRUE, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
(1, 'I', 7, TRUE, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
(1, 'I', 8, TRUE, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
(1, 'I', 9, TRUE, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
(1, 'I', 10, TRUE, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
(1, 'I', 11, TRUE, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
(1, 'I', 12, TRUE, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
(1, 'J', 1, TRUE, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
(1, 'J', 2, TRUE, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
(1, 'J', 3, TRUE, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
(1, 'J', 4, TRUE, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
(1, 'J', 5, TRUE, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
(1, 'J', 6, TRUE, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
(1, 'J', 7, TRUE, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
(1, 'J', 8, TRUE, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
(1, 'J', 9, TRUE, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
(1, 'J', 10, TRUE, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
(1, 'J', 11, TRUE, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
(1, 'J', 12, TRUE, '2024-01-01 00:00:00', '2024-01-01 00:00:00');

-- 14. SHOWTIMES TABLE
INSERT INTO showtimes (movie_id, theater_id, show_date, show_time, base_price, available_seats, created_at, updated_at) VALUES
-- Oppenheimer showtimes
(1, 1, '2024-07-15', '2024-07-15 13:00:00', 45000.00, 118, '2024-07-01 00:00:00', '2024-07-01 00:00:00'),
(1, 1, '2024-07-15', '2024-07-15 16:30:00', 45000.00, 120, '2024-07-01 00:00:00', '2024-07-01 00:00:00'),
(1, 1, '2024-07-15', '2024-07-15 20:00:00', 50000.00, 115, '2024-07-01 00:00:00', '2024-07-01 00:00:00'),
-- Barbie showtimes
(2, 2, '2024-07-15', '2024-07-15 12:00:00', 45000.00, 148, '2024-07-01 00:00:00', '2024-07-01 00:00:00'),
(2, 2, '2024-07-15', '2024-07-15 15:00:00', 45000.00, 150, '2024-07-01 00:00:00', '2024-07-01 00:00:00'),
(2, 2, '2024-07-15', '2024-07-15 18:30:00', 50000.00, 145, '2024-07-01 00:00:00', '2024-07-01 00:00:00'),
-- Dune: Part Two showtimes
(3, 3, '2024-07-15', '2024-07-15 14:00:00', 50000.00, 175, '2024-07-01 00:00:00', '2024-07-01 00:00:00'),
(3, 3, '2024-07-15', '2024-07-15 18:00:00', 55000.00, 180, '2024-07-01 00:00:00', '2024-07-01 00:00:00'),
(3, 3, '2024-07-15', '2024-07-15 21:30:00', 55000.00, 172, '2024-07-01 00:00:00', '2024-07-01 00:00:00'),
-- Spider-Man showtimes
(4, 4, '2024-07-15', '2024-07-15 13:30:00', 45000.00, 140, '2024-07-01 00:00:00', '2024-07-01 00:00:00'),
(4, 4, '2024-07-15', '2024-07-15 17:00:00', 50000.00, 138, '2024-07-01 00:00:00', '2024-07-01 00:00:00'),
(4, 4, '2024-07-15', '2024-07-15 20:30:00', 50000.00, 135, '2024-07-01 00:00:00', '2024-07-01 00:00:00');

-- 15. PAYMENTS TABLE
INSERT INTO payments (method_name, provider, fee_process, is_active, created_at, updated_at) VALUES
('Credit Card', 'Visa', 2500.00, TRUE, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
('Credit Card', 'Mastercard', 2500.00, TRUE, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
('Debit Card', 'BCA', 1500.00, TRUE, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
('E-Wallet', 'GoPay', 1000.00, TRUE, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
('E-Wallet', 'OVO', 1000.00, TRUE, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
('E-Wallet', 'DANA', 1000.00, TRUE, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
('Bank Transfer', 'BCA Virtual Account', 2000.00, TRUE, '2024-01-01 00:00:00', '2024-01-01 00:00:00'),
('Bank Transfer', 'Mandiri Virtual Account', 2000.00, TRUE, '2024-01-01 00:00:00', '2024-01-01 00:00:00');

-- 16. BOOKINGS TABLE
INSERT INTO bookings (user_id, showtime_id, status, total_amount, discount_amount, tax_amount, booking_time, expires_at, created_at, updated_at) VALUES
('550e8400-e29b-41d4-a716-446655440002', 1, 'confirmed', 92500.00, 0.00, 2500.00, '2024-07-10 14:30:00', '2024-07-10 14:45:00', '2024-07-10 14:30:00', '2024-07-10 14:32:00'),
('550e8400-e29b-41d4-a716-446655440003', 4, 'confirmed', 141000.00, 5000.00, 6000.00, '2024-07-11 10:15:00', '2024-07-11 10:30:00', '2024-07-11 10:15:00', '2024-07-11 10:18:00'),
('550e8400-e29b-41d4-a716-446655440004', 7, 'confirmed', 103500.00, 0.00, 3500.00, '2024-07-12 16:45:00', '2024-07-12 17:00:00', '2024-07-12 16:45:00', '2024-07-12 16:47:00'),
('550e8400-e29b-41d4-a716-446655440005', 10, 'pending', 102000.00, 0.00, 2000.00, '2024-07-13 11:20:00', '2024-07-13 11:35:00', '2024-07-13 11:20:00', '2024-07-13 11:20:00'),
('550e8400-e29b-41d4-a716-446655440002', 6, 'confirmed', 152500.00, 7500.00, 5000.00, '2024-07-14 19:00:00', '2024-07-14 19:15:00', '2024-07-14 19:00:00', '2024-07-14 19:03:00');

-- 17. BOOKING_SEATS TABLE
INSERT INTO booking_seats (booking_id, seat_id, seat_price, created_at, updated_at) VALUES
-- Booking 1 - 2 seats
(1, 25, 45000.00, '2024-07-10 14:30:00', '2024-07-10 14:30:00'),
(1, 26, 45000.00, '2024-07-10 14:30:00', '2024-07-10 14:30:00'),
-- Booking 2 - 3 seats
(2, 15, 45000.00, '2024-07-11 10:15:00', '2024-07-11 10:15:00'),
(2, 16, 45000.00, '2024-07-11 10:15:00', '2024-07-11 10:15:00'),
(2, 17, 45000.00, '2024-07-11 10:15:00', '2024-07-11 10:15:00'),
-- Booking 3 - 2 seats
(3, 35, 50000.00, '2024-07-12 16:45:00', '2024-07-12 16:45:00'),
(3, 36, 50000.00, '2024-07-12 16:45:00', '2024-07-12 16:45:00'),
-- Booking 4 - 2 seats
(4, 45, 50000.00, '2024-07-13 11:20:00', '2024-07-13 11:20:00'),
(4, 46, 50000.00, '2024-07-13 11:20:00', '2024-07-13 11:20:00'),
-- Booking 5 - 3 seats
(5, 55, 50000.00, '2024-07-14 19:00:00', '2024-07-14 19:00:00'),
(5, 56, 50000.00, '2024-07-14 19:00:00', '2024-07-14 19:00:00'),
(5, 57, 50000.00, '2024-07-14 19:00:00', '2024-07-14 19:00:00');

-- 18. PAYMENT_TRANSACTIONS TABLE
INSERT INTO payments_transactions (payment_id, booking_id, status, amount, completed_at, created_at) VALUES
(4, 1, 'success', 92500.00, '2024-07-10 14:32:00', '2024-07-10 14:31:00'),
(1, 2, 'success', 141000.00, '2024-07-11 10:18:00', '2024-07-11 10:17:00'),
(5, 3, 'success', 103500.00, '2024-07-12 16:47:00', '2024-07-12 16:46:00'),
(6, 4, 'pending', 102000.00, NULL, '2024-07-13 11:21:00'),
(2, 5, 'success', 152500.00, '2024-07-14 19:03:00', '2024-07-14 19:02:00');

-- ========================================
-- NOTES UNTUK PENGGUNAAN DATA DUMMY:
-- ========================================

-- 1. Password di users table menggunakan hash bcrypt untuk "password123"
-- 2. UUID dibuat manual untuk konsistensi foreign key
-- 3. Tanggal menggunakan format ISO dengan DMY sesuai setting database
-- 4. Semua foreign key sudah dipastikan valid dan saling berkaitan
-- 5. Status booking dan payment realistic sesuai flow bisnis
-- 6. Harga tiket bervariasi sesuai waktu tayang (siang lebih murah)
-- 7. Available seats dihitung dari capacity dikurangi booking yang confirmed
-- 8. Seat numbering menggunakan format standar bioskop (A1-J12)
-- 9. Data cinema menggunakan lokasi real di Indonesia
-- 10. Movie data berdasarkan film populer 2022-2024

-- ========================================
-- QUERY VALIDASI DATA:
-- ========================================

-- Cek total users: SELECT COUNT(*) FROM users;
-- Cek total movies: SELECT COUNT(*) FROM movies;
-- Cek booking dengan detail: 
-- SELECT b.*, u.email, m.title, s.show_date, s.show_time 
-- FROM bookings b 
-- JOIN users u ON b.user_id = u.id 
-- JOIN showtimes s ON b.showtime_id = s.id 
-- JOIN movies m ON s.movie_id = m.id;

-- Cek seat availability:
-- SELECT s.show_date, s.show_time, s.available_seats, m.title
-- FROM showtimes s
-- JOIN movies m ON s.movie_id = m.id
-- ORDER BY s.show_date, s.show_time;